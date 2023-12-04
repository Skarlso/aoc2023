package main

import (
	"fmt"
	"os"
	"strings"
)

type card struct {
	numbersYouHave []string
	winningNumbers map[string]struct{}
	copies         int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	// ids := 0
	cards := make([]*card, 0)
	for _, l := range split {
		split := strings.Split(l, ": ")
		c := split[1]
		hands := strings.Split(c, " | ")
		winning := hands[0]
		got := hands[1]

		m := make(map[string]struct{})
		for _, n := range strings.Split(winning, " ") {
			m[n] = struct{}{}
		}

		cards = append(cards, &card{
			numbersYouHave: strings.Split(got, " "),
			winningNumbers: m,
		})
	}

	for j, c := range cards {
		copies := 0

		// TODO: Pretty sure this is some kind of sum... but I don't have the time to figure it out.
		for _, n := range c.numbersYouHave {
			if n != "" {
				if _, ok := c.winningNumbers[n]; ok {
					copies++
				}
			}
		}

		for i := 1; i <= copies; i++ {
			if j+i >= len(cards) {
				break
			}

			cards[j+i].copies++
		}

		for i := 0; i < c.copies; i++ {
			copies := 0

			for _, n := range c.numbersYouHave {
				if n != "" {
					if _, ok := c.winningNumbers[n]; ok {
						copies++
					}
				}
			}

			for i := 1; i <= copies; i++ {
				if j+i >= len(cards) {
					break
				}

				cards[j+i].copies++
			}
		}
	}

	sum := 0
	for _, c := range cards {
		sum += 1 + c.copies
	}

	fmt.Println(sum)
}
