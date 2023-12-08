package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

type node struct {
	left  *node
	right *node
	value string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	steps := ""
	match := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	nodes := make(map[string]*node)
	for i, l := range split {
		if i == 0 {
			steps = l

			continue
		}

		if l == "" {
			continue
		}

		var (
			n, left, right string
		)

		matches := match.FindAllStringSubmatch(l, -1)
		// fmt.Println(matches)
		m := matches[0]
		n = m[1]
		left = m[2]
		right = m[3]

		// fmt.Println(node, left, right)

		if v, ok := nodes[n]; ok {
			v.left = &node{value: left}
			v.right = &node{value: right}
		} else {
			nodes[n] = &node{value: n, right: &node{value: right}, left: &node{value: left}}
		}
	}

	// fmt.Println(steps)

	// for k, v := range nodes {
	// 	fmt.Println(k, v.left.value, v.right.value)
	// }
	index := 0
	taken := 0
	current := nodes["AAA"]
	for current.value != "ZZZ" {
		if index >= len(steps) {
			index = 0
		}
		nextStep := steps[index]
		index++
		taken++

		if nextStep == 'L' {
			current = nodes[current.left.value]
		} else {
			current = nodes[current.right.value]
		}
	}

	fmt.Println("steps: ", taken)
}
