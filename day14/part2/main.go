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

	// get the point where it will start to be the same.
	// sum and iterations at which it occurred
	values := map[int][]int{}
	end := 10000
	begin := 0
	for begin < end {
		// slide north
		// fmt.Println("tilt north")
		slideNorth(maze)
		// display(maze)

		// fmt.Println()

		// slide west
		// fmt.Println("tilt west")
		slideWest(maze)
		// display(maze)

		// fmt.Println()

		// slide south
		// fmt.Println("tilt south")
		slideSouth(maze)
		// display(maze)

		// fmt.Println()

		// slide east
		slide(func(maze [][]rune) int {
			return len(maze[0]) - 1
		}, func(maze [][]rune) int {
			return 0
		}, func(maze [][]rune, y *int) bool {
			if *y < len(maze) {
				return false
			}

			*y++

			return true
		}, func(maze [][]rune, x *int) bool {
			// for x := len(maze[y]) - 1; x >= 0; x-- {
			if *x >= 0 {
				return false
			}

			*x--

			return true
		}, maze, point{x: -1, y: 0})

		// slideEast(maze)

		// display(maze)
		begin++

		// fmt.Println()
		// count
		sum := 0
		col := len(maze)
		for y := 0; y < len(maze); y++ {
			rockCount := strings.Count(string(maze[y]), "O")
			// fmt.Println("add: ", rockCount*col)
			sum += rockCount * col
			col--
		}

		// fmt.Println(sum)
		// if len(values[sum]) < 4 {
		values[sum] = append(values[sum], begin)
		if len(values[sum]) > 1 {
			break
		}
		// }
	}

	// ah it's the frequency using % and the loop value

	// fmt.Println(lowestFirstReoccurrence)
	fmt.Println(values)
	// display(maze)
}

func slideNorth(maze [][]rune) {
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
}

func slideWest(maze [][]rune) {
	// y
	freeSpots := map[int]point{}

	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == '.' {
				if _, ok := freeSpots[y]; !ok {
					freeSpots[y] = point{x: x, y: y}
				}
			}

			if maze[y][x] == '#' {
				freeSpots[y] = point{x: x + 1, y: y}
			}

			if maze[y][x] == 'O' {
				if p, ok := freeSpots[y]; ok {
					if maze[p.y][p.x] != 'O' {
						maze[p.y][p.x] = 'O'
						maze[y][x] = '.'
						freeSpots[y] = point{x: p.x + 1, y: p.y}
					} else {
						freeSpots[y] = point{x: p.x + 1, y: p.y}
					}
				}
			}
		}
	}
}

func slide(startX func(maze [][]rune) int, startY func(maze [][]rune) int, forY func(maze [][]rune, y *int) bool, forX func(maze [][]rune, x *int) bool, maze [][]rune, direction point) {
	// y
	freeSpots := map[int]point{}
	var (
		x, y = startX(maze), startY(maze)
	)

	// for y := 0; y < len(maze); y++ {
	// for x := len(maze[y]) - 1; x >= 0; x-- {
	for forY(maze, &y) {
		for forX(maze, &x) {
			if maze[y][x] == '.' {
				if _, ok := freeSpots[y]; !ok {
					freeSpots[y] = point{x: x, y: y}
				}
			}

			if maze[y][x] == '#' {
				// freeSpots[y] = point{x: x - 1, y: y}
				freeSpots[y] = point{x: x + direction.x, y: y + direction.y}
			}

			if maze[y][x] == 'O' {
				if p, ok := freeSpots[y]; ok {
					if maze[p.y][p.x] != 'O' {
						maze[p.y][p.x] = 'O'
						maze[y][x] = '.'
						freeSpots[y] = point{x: p.x + direction.x, y: p.y + direction.y}
					} else {
						freeSpots[y] = point{x: p.x + direction.x, y: p.y + direction.y}
					}
				}
			}
		}
	}
}

func slideSouth(maze [][]rune) {
	// x
	freeSpots := map[int]point{}

	for y := len(maze) - 1; y >= 0; y-- {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == '.' {
				if _, ok := freeSpots[x]; !ok {
					freeSpots[x] = point{x: x, y: y}
				}
			}

			if maze[y][x] == '#' {
				freeSpots[x] = point{x: x, y: y - 1}
			}

			if maze[y][x] == 'O' {
				if p, ok := freeSpots[x]; ok {
					if maze[p.y][p.x] != 'O' {
						maze[p.y][p.x] = 'O'
						maze[y][x] = '.'
						freeSpots[x] = point{x: p.x, y: p.y - 1}
					} else {
						freeSpots[x] = point{x: p.x, y: p.y - 1}
					}
				}
			}
		}
	}
}

func display(maze [][]rune) {
	for _, col := range maze {
		for _, row := range col {
			fmt.Print(string(row))
		}

		fmt.Println()
	}
}
