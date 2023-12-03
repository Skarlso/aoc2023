package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	sum := 0
	for _, l := range split {
		var (
			first, last string
		)
		for _, r := range []rune(l) {
			if unicode.IsDigit(r) {
				if first == "" {
					first = string(r)
				} else {
					last = string(r)
				}
			}
		}
		if last == "" {
			last = first
		}

		n, _ := strconv.Atoi(fmt.Sprintf("%s%s", first, last))
		sum += n
	}

	fmt.Println("sum: ", sum)
}
