package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), ",")
	sum := 0
	for _, seq := range split {
		s := 0
		for _, r := range seq {
			s += int(r)
			s *= 17
			s %= 256
		}

		sum += s
	}

	fmt.Println("sum: ", sum)
}
