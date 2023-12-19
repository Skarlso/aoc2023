package main

import (
	"fmt"
	"math"
	"os"
	"strings"
)

type elf struct {
	stepsInDirection    int
	location, direction point
}

type point struct {
	x, y int
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	maze := [][]int{}
	for _, l := range split {
		var row []int
		for _, r := range l {
			row = append(row, int(r)-'0')
		}

		maze = append(maze, row)
	}

	start := point{x: 0, y: 0}
	goal := point{x: len(maze[0]) - 1, y: len(maze) - 1}

	starterElf := elf{location: start, direction: point{x: 1, y: 0}, stepsInDirection: 0}
	cost := map[elf]int{
		starterElf: 0,
	}

	queue := []elf{starterElf}
	var current elf
	minCost := math.MaxInt

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		// display(maze, current.location)
		// time.Sleep(300 * time.Millisecond)

		if current.location == goal && current.stepsInDirection >= 4 {
			minCost = min(minCost, cost[current])
		}

		for _, d := range []point{current.direction, left(current.direction), right(current.direction)} {
			np := point{x: current.location.x + d.x, y: current.location.y + d.y}
			if np.x < 0 || np.y < 0 || np.y >= len(maze) || np.x >= len(maze[np.y]) {
				continue
			}

			totalLoss := cost[current] + maze[np.y][np.x]
			nextStepsInDirection := 1

			if d == current.direction {
				nextStepsInDirection = current.stepsInDirection + 1
			}

			if (d == current.direction && current.stepsInDirection < 10) || (d != current.direction && current.stepsInDirection >= 4) {
				next := elf{location: np, direction: d, stepsInDirection: nextStepsInDirection}

				if v, ok := cost[next]; !ok || v > totalLoss {
					cost[next] = totalLoss
					queue = append(queue, next)
				}
			}
		}

	}

	fmt.Println("min cost: ", minCost)
}

func display(maze [][]int, loc point) {
	for y, row := range maze {
		for x, col := range row {
			if y == loc.y && x == loc.x {
				fmt.Print(".")
			} else {
				fmt.Print(col)
			}
		}

		fmt.Println()
	}

	fmt.Println()
}

func left(p point) point {
	return point{x: p.y, y: -p.x}
}

func right(p point) point {
	return point{x: -p.y, y: p.x}
}
