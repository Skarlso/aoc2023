package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

var max = map[string]int{
	"red":   12,
	"green": 13,
	"blue":  14,
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
			bag string
		)
		sp := strings.Split(l, ": ")
		// fmt.Println(game)
		bag = sp[1]
		// fmt.Printf("game: %s: bag: %s\n", game, bag)

		maxes := make(map[string]int)
		hands := strings.Split(bag, ";")
		// fmt.Println(hands)
		for _, h := range hands {
			colors := strings.Split(h, ", ")

			for _, c := range colors {
				num := strings.Split(strings.Trim(c, " "), " ")
				// fmt.Println(num)
				n, err := strconv.Atoi(num[0])
				if err != nil {
					log.Fatal(err)
				}
				// fmt.Println("num: ", n)
				color := num[1]

				if n > maxes[color] {
					maxes[color] = n
				}

				// score[color] += n
				// if score[color] > max[color] {
				// 	continue loop
				// }
			}
		}

		// fmt.Println(maxes)
		power := 1
		for _, v := range maxes {
			power *= v
		}
		sum += power
	}

	fmt.Println("valid: ", sum)
}
