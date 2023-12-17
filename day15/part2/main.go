package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type lense struct {
	label       string
	focalLenght int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), ",")
	boxes := map[int][]*lense{}
	for _, r := range split {
		if strings.Contains(r, "-") {
			removeLense(boxes, r)
		} else {
			applyLense(boxes, r)
		}
	}

	sum := 0
	for k, v := range boxes {
		for i, l := range v {
			sum += (k + 1) * (i + 1) * l.focalLenght
		}
	}

	fmt.Println("sum: ", sum)
}

var reg = regexp.MustCompile(`(\w+)=(\d+)`)

// if exists: update focal length
// if not: add as is at the end
func applyLense(boxes map[int][]*lense, l string) {
	m := reg.FindAllStringSubmatch(l, -1)
	label := m[0][1]
	fl, _ := strconv.Atoi(m[0][2])

	lense := &lense{label: label, focalLenght: fl}
	h := hash(label)

	// update if it exists
	for i, b := range boxes[h] {
		if b.label == lense.label {
			boxes[h][i].focalLenght = fl

			return
		}
	}

	// append if it doesn't
	boxes[h] = append(boxes[h], lense)
}

func removeLense(boxes map[int][]*lense, l string) {
	l = string(l[:len(l)-1])
	h := hash(l)
	if _, ok := boxes[h]; ok {
		for i, r := range boxes[h] {
			if r.label == l {
				boxes[h] = append(boxes[h][:i], boxes[h][i+1:]...)

				return
			}
		}
	}
}

func hash(seq string) int {
	s := 0
	for _, r := range seq {
		s += int(r)
		s *= 17
		s %= 256
	}

	return s
}
