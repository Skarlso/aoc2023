package main

import (
	"fmt"
	"os"
	"strings"
)

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

	maze := [][]rune{}
	for _, l := range split {
		maze = append(maze, []rune(l))
	}

	// Tilt
	// column (x), spot
	freeSpots := map[int]point{}
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == '.' {
				if _, ok := freeSpots[x]; !ok {
					freeSpots[x] = point{x: x, y: y}
				}
			}

			if maze[y][x] == '#' {
				freeSpots[x] = point{x: x, y: y + 1}
			}

			if maze[y][x] == 'O' {
				if p, ok := freeSpots[x]; ok {
					if maze[p.y][p.x] != 'O' {
						maze[p.y][p.x] = 'O'
						maze[y][x] = '.'
						freeSpots[x] = point{x: p.x, y: p.y + 1}
					} else {
						freeSpots[x] = point{x: p.x, y: p.y + 1}
					}
				}
			}
		}
	}

	// count
	sum := 0
	col := len(maze)
	for y := 0; y < len(maze); y++ {
		rockCount := strings.Count(string(maze[y]), "O")
		// fmt.Println("add: ", rockCount*col)
		sum += rockCount * col
		col--
	}

	fmt.Println(sum)
	// display(maze)
}

func display(maze [][]rune) {
	for _, col := range maze {
		for _, row := range col {
			fmt.Print(string(row))
		}

		fmt.Println()
	}
}
