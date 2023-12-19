package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type point struct {
	x, y int
}

var directions = map[string]point{
	"0": {x: 1, y: 0},
	"1": {x: 0, y: -1},
	"2": {x: -1, y: 0},
	"3": {x: 0, y: 1},
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
		color = strings.Trim(color, "(")
		color = strings.Trim(color, ")")
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

		direction := string(plan.color[len(plan.color)-1])
		// fmt.Println("direction: ", direction)
		digits := plan.color[1:6]
		meters, _ := strconv.ParseInt(digits, 16, 64)

		for m := 0; m < int(meters); m++ {
			if direction == "2" || direction == "1" {
				downOrLeftLine++
			}

			current = point{x: current.x + directions[direction].x, y: current.y + directions[direction].y}
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
