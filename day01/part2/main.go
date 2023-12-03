package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

var numbers = map[string]string{
	"one":   "1",
	"two":   "2",
	"three": "3",
	"four":  "4",
	"five":  "5",
	"six":   "6",
	"seven": "7",
	"eight": "8",
	"nine":  "9",
}

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

		for i := 0; i < len(l); i++ {
			if unicode.IsDigit(rune(l[i])) {
				if first == "" {
					first = string(l[i])
				} else {
					last = string(l[i])
				}
			} else {
				for k, v := range numbers {
					if i+len(k) <= len(l) {
						n := l[i : i+len(k)]

						if n == k {
							if first == "" {
								first = v
							} else {
								last = v
							}
						}

					}
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
