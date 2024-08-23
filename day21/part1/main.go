package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type step struct {
	p     point
	count int // carry the step count with us
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	garden := [][]rune{}
	start := point{}
	for y, line := range split {
		garden = append(garden, []rune(line))
		for x, c := range line {
			if c == 'S' {
				start.x = x
				start.y = y
			}
		}
	}

	// fmt.Println(garden)
	// fmt.Println(start)

	queue := []step{{p: start}}
	seen := map[point]struct{}{}
	answer := []point{}
	var current step

	for len(queue) > 0 {
		current, queue = queue[0], queue[1:]

		if current.count <= 64 && current.count%2 == 0 {
			answer = append(answer, current.p)
		}

		for _, next := range steps(garden, current.p) {
			if _, ok := seen[next]; !ok {
				queue = append(queue, step{p: next, count: current.count + 1})
				seen[next] = struct{}{}
			}
		}
	}

	fmt.Println(len(answer) - 1)
}

var directions = []point{
	{x: 0, y: 1},
	{x: 1, y: 0},
	{x: 0, y: -1},
	{x: -1, y: 0},
}

func steps(garden [][]rune, loc point) []point {
	var result []point
	for _, d := range directions {
		next := point{x: loc.x + d.x, y: loc.y + d.y}
		if next.x < 0 || next.y < 0 || next.y >= len(garden) || next.x >= len(garden[next.y]) || garden[next.y][next.x] == '#' {
			continue
		}

		result = append(result, next)
	}

	return result
}
