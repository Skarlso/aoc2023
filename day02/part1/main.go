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
	ids := 0
loop:
	for _, l := range split {
		var (
			game string
			bag  string
		)
		sp := strings.Split(l, ": ")
		game = sp[0]
		// fmt.Println(game)
		bag = sp[1]
		// fmt.Printf("game: %s: bag: %s\n", game, bag)

		// score := make(map[string]int)

		hands := strings.Split(bag, ";")
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

				if n > max[color] {
					continue loop
				}

				// score[color] += n
				// if score[color] > max[color] {
				// 	continue loop
				// }
			}
			// fmt.Println(score)
		}

		var id int
		fmt.Sscanf(game, "Game %d", &id)
		fmt.Println("id: ", id)
		ids += id
	}

	fmt.Println("valid: ", ids)
}
