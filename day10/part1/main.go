package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

type matcher struct {
	point point
	match func(origin, r rune) bool
}

var heading = []matcher{
	// right
	{
		point: point{x: 1, y: 0},
		match: func(origin, r rune) bool {

			return (origin == '-' || origin == 'L' || origin == 'F') && (r == '-' || r == 'J' || r == '7')
		},
	},
	// down
	{
		point: point{x: 0, y: 1},
		match: func(origin, r rune) bool {
			return (origin == '|' || origin == '7' || origin == 'F') && (r == '|' || r == 'J' || r == 'L')
		},
	},
	// left
	{
		point: point{x: -1, y: 0},
		match: func(origin, r rune) bool {
			return (origin == '-' || origin == 'J' || origin == '7') && (r == '-' || r == 'L' || r == 'F')
		},
	},
	// up
	{
		point: point{x: 0, y: -1},
		match: func(origin, r rune) bool {
			return (origin == '|' || origin == 'L' || origin == 'J') && (r == '|' || r == 'F' || r == '7')
		},
	},
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
	startingPosition := point{}
	for y, l := range split {
		x := strings.Index(l, "S")
		if x != -1 {
			startingPosition = point{x: x, y: y}
		}
		maze = append(maze, []rune(l))
	}

	fmt.Println("starting position:", startingPosition)
	// TODO: refactor the hell out of this
	maze[startingPosition.y][startingPosition.x] = changeStartingPipe(maze, startingPosition)
	fmt.Println("starting: ", string(maze[startingPosition.y][startingPosition.x]))

	current := findPath(maze, startingPosition, startingPosition)
	// step is 1 because we already stepped to current from starting position
	steps := 1

	prev := startingPosition
	for {
		if current == startingPosition {
			break
		}

		newPoint := findPath(maze, current, prev)
		prev = current
		current = newPoint
		steps++
	}

	fmt.Println("steps: ", steps/2) // +1 because of the start
}

// we are in a loop so technically we should have only ONE way out except the one from which we came from
func findPath(maze [][]rune, p point, prev point) point {
	for _, d := range heading {
		// fmt.Println("d: ", d)
		newPoint := point{x: p.x + d.point.x, y: p.y + d.point.y}
		if newPoint.x < 0 || newPoint.y < 0 || newPoint.y >= len(maze) || newPoint.x >= len(maze[newPoint.y]) {
			continue
		}

		if d.match(maze[p.y][p.x], maze[newPoint.y][newPoint.x]) && newPoint != prev {
			// fmt.Println("new point match: ", newPoint)
			return newPoint
		}
	}

	// this shouldn't happen
	return point{x: -1, y: -1}
}

func changeStartingPipe(maze [][]rune, p point) rune {
	// Yeah, I'm pretty sure this could be done better recursively, but fuck that. :D
	if v, ok := checkFromUp(maze, p); ok {
		return v
	}

	if v, ok := checkFromRight(maze, p); ok {
		return v
	}

	if v, ok := checkFromDown(maze, p); ok {
		return v
	}

	if v, ok := checkFromLeft(maze, p); ok {
		return v
	}

	return 'N'
}

var (
	fromUp    = []rune{'|', 'F', '7'}
	fromRight = []rune{'-', 'J', '7'}
	fromDown  = []rune{'|', 'J', 'L'}
	fromLeft  = []rune{'-', 'L', 'F'}
)

func checkFromUp(maze [][]rune, p point) (rune, bool) {
	for _, u := range fromUp {
		if maze[p.y-1][p.x] != u {
			continue
		}

		for _, l := range fromLeft {
			if maze[p.y][p.x-1] == l {
				return 'J', true
			}
		}

		for _, d := range fromDown {
			if maze[p.y+1][p.x] == d {
				return '|', true
			}
		}

		for _, d := range fromRight {
			if maze[p.y][p.x+1] == d {
				return 'L', true
			}
		}
	}

	return 'N', false
}

func checkFromLeft(maze [][]rune, p point) (rune, bool) {
	for _, r := range fromLeft {
		if maze[p.y][p.x-1] != r {
			continue
		}

		for _, l := range fromDown {
			if maze[p.y+1][p.x] == l {
				return 'F', true
			}
		}

		for _, d := range fromRight {
			if maze[p.y][p.x+1] == d {
				return '-', true
			}
		}

		for _, d := range fromUp {
			if maze[p.y-1][p.x] == d {
				return 'L', true
			}
		}
	}

	return 'N', false
}

func checkFromDown(maze [][]rune, p point) (rune, bool) {
	for _, r := range fromDown {
		if maze[p.y+1][p.x] != r {
			continue
		}

		for _, l := range fromRight {
			if maze[p.y][p.x+1] == l {
				return 'F', true
			}
		}

		for _, d := range fromUp {
			if maze[p.y-1][p.x] == d {
				return '|', true
			}
		}

		for _, d := range fromLeft {
			if maze[p.y][p.x-1] == d {
				return '7', true
			}
		}
	}

	return 'N', false
}

func checkFromRight(maze [][]rune, p point) (rune, bool) {
	for _, r := range fromRight {
		if maze[p.y][p.x+1] != r {
			continue
		}

		for _, l := range fromDown {
			if maze[p.y+1][p.x] == l {
				return 'F', true
			}
		}

		for _, d := range fromLeft {
			if maze[p.y][p.x-1] == d {
				return '-', true
			}
		}

		for _, d := range fromUp {
			if maze[p.y-1][p.x] == d {
				return 'L', true
			}
		}
	}

	return 'N', false
}

func display(maze [][]rune) {
	for _, c := range maze {
		for _, r := range c {
			fmt.Print(string(r))
		}

		fmt.Println()
	}
}
