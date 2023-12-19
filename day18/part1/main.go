package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

var directions = map[string]point{
	"R": {x: 1, y: 0},
	"D": {x: 0, y: -1},
	"L": {x: -1, y: 0},
	"U": {x: 0, y: 1},
}

type digPlan struct {
	direction string
	meter     int
	color     string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	digPlans := []digPlan{}
	for _, l := range split {
		var (
			d     string
			meter int
			color string
		)

		fmt.Sscanf(l, "%s %d %s", &d, &meter, &color)

		digPlans = append(digPlans, digPlan{direction: d, meter: meter, color: color})
	}

	// fmt.Println(digPlans)
	start := point{x: 0, y: 0}
	var (
		x []int
		y []int
	)

	maxX := 0
	current := start
	downOrLeftLine := 0 // because the lines have width
	for _, plan := range digPlans {
		for m := 0; m < plan.meter; m++ {
			if plan.direction == "L" || plan.direction == "D" {
				downOrLeftLine++
			}

			current = point{x: current.x + directions[plan.direction].x, y: current.y + directions[plan.direction].y}
			x = append(x, current.x)
			y = append(y, current.y)
			if current.x > maxX {
				maxX = current.x
			}
		}
	}
	// fmt.Println(len(x))
	fmt.Println(area(x, y) + downOrLeftLine + 1) // +1 for starting position
}

func area(x []int, y []int) int {
	area := 0
	j := len(x) - 1

	for i := 0; i < len(x); i++ {
		area += (x[j] + x[i]) * (y[j] - y[i])
		j = i
	}

	return area / 2
}
