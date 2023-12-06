package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type boat struct {
	speed int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	var times []int
	var distances []int

	for _, l := range split {
		time := false
		if strings.Contains(l, "Time:") {
			time = true
		}
		split := strings.Split(l, ": ")
		numbers := strings.Split(split[1], " ")
		for _, n := range numbers {
			num, err := strconv.Atoi(n)
			if err != nil {
				continue
			}

			if time {
				times = append(times, num)
			} else {
				distances = append(distances, num)
			}
		}
	}

	// fmt.Println(times, distances)
	sum := 1
	for i, t := range times {
		// As soon as it starts to decline and reaches below the current max, we stop
		// We start counting as soon as it's over the current distance max.

		currentMax := distances[i]
		prev := 0

		// if smaller than prev and equals or smaller than currentMax, stop.
		wins := 0
		for i := 0; i < t; i++ {
			speed := i

			distanceTraveled := (t - i) * speed

			if distanceTraveled < prev && currentMax >= distanceTraveled {
				break
			}

			prev = distanceTraveled

			if distanceTraveled > currentMax {
				wins++
			}
			// fmt.Println(distanceTraveled)
		}

		sum *= wins
	}

	fmt.Println(sum)
}
