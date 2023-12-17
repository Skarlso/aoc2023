package main

import (
	"fmt"
	"os"
	"strings"
)

type point struct {
	x, y int
}

// save the memory of everyone.
var mirrors = map[rune]func(current point, b *beam) []*beam{
	'/': func(current point, b *beam) []*beam {
		newHeading := point{x: -b.heading.y, y: -b.heading.x}
		return []*beam{
			{
				heading: newHeading,
				visited: b.visited,
			},
		}
	},
	'\\': func(current point, b *beam) []*beam {
		newHeading := point{x: b.heading.y, y: b.heading.x}
		return []*beam{
			{
				heading: newHeading,
				visited: b.visited,
			},
		}
	},
	'|': func(current point, b *beam) []*beam {
		switch b.heading {
		case point{x: -1, y: 0}, point{x: 1, y: 0}:
			return []*beam{
				{
					heading: point{x: 0, y: -1},
					// current: point{x: current.x, y: current.y - 1},
					current: point{x: current.x, y: current.y},
					visited: b.visited,
				},
				{
					heading: point{x: 0, y: 1},
					// current: point{x: current.x, y: current.y + 1},
					current: point{x: current.x, y: current.y},
					visited: b.visited,
				},
			}
		case point{x: 0, y: 1}, point{x: 0, y: -1}:
			return []*beam{
				{
					heading: b.heading,
					visited: b.visited,
					// current: point{x: current.x + b.heading.x, y: current.y + b.heading.y},
				},
			}
		}

		return nil
	},
	'-': func(current point, b *beam) []*beam {
		switch b.heading {
		case point{x: 0, y: 1}, point{x: 0, y: -1}:
			return []*beam{
				{
					heading: point{x: -1, y: 0},
					// current: point{x: current.x - 1, y: current.y},
					current: point{x: current.x, y: current.y},
					visited: b.visited,
				},
				{
					heading: point{x: 1, y: 0},
					// current: point{x: current.x + 1, y: current.y},
					current: point{x: current.x, y: current.y},
					visited: b.visited,
				},
			}
		case point{x: 1, y: 0}, point{x: -1, y: 0}:
			return []*beam{
				{
					heading: b.heading,
					// current: point{x: current.x + b.heading.x, y: current.y + b.heading.y},
					visited: b.visited,
				},
			}
		}

		return nil
	},
	'.': func(current point, b *beam) []*beam {
		return []*beam{
			{
				// current: current,
				heading: b.heading,
				visited: b.visited,
			},
		}
	},
	// this is my mark
	'o': func(current point, b *beam) []*beam {
		return []*beam{
			{
				// current: current,
				heading: b.heading,
				visited: b.visited,
			},
		}
	},
}

type impact struct {
	location      point
	fromDirection point
}

type beam struct {
	heading point
	current point
	visited map[impact]struct{}
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

	// from the top
	visited := map[impact]struct{}{}
	maxEnergy := 0
	energized := map[point]struct{}{}
	x := 0
	for {
		if x >= len(maze[0]) {
			break
		}

		start := point{x: x, y: -1}
		queue := []*beam{
			{heading: point{x: 0, y: 1}, current: start, visited: map[impact]struct{}{}},
		}

		// we stop when no beams can move
		var current *beam
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]

			energized[current.current] = struct{}{}
			// if _, ok := visited[impact{location: current.current, fromDirection: current.heading}]; ok {
			// 	// we already reached this point by others...
			// 	break
			// }

			visited[impact{location: current.current, fromDirection: current.heading}] = struct{}{}

			queue = append(queue, moveBeam(maze, current)...)
		}

		if len(energized)-1 > maxEnergy {
			maxEnergy = len(energized) - 1
		}

		energized = map[point]struct{}{}
		x++
	}

	fmt.Println(maxEnergy)

	energized = map[point]struct{}{}

	// // TODO: Put this into a function.
	// start = point{x: -1, y: 0}

	// from the bottom
	// visited = map[impact]struct{}{}
	// maxEnergy := 0
	// energized := map[point]struct{}{}
	x = 0
	for {
		if x >= len(maze[0]) {
			break
		}

		start := point{x: x, y: len(maze)}
		queue := []*beam{
			{heading: point{x: 0, y: -1}, current: start, visited: map[impact]struct{}{}},
		}

		// we stop when no beams can move
		var current *beam
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]

			energized[current.current] = struct{}{}
			// if _, ok := visited[impact{location: current.current, fromDirection: current.heading}]; ok {
			// 	// we already reached this point by others...
			// 	break
			// }

			visited[impact{location: current.current, fromDirection: current.heading}] = struct{}{}

			queue = append(queue, moveBeam(maze, current)...)
		}

		if len(energized)-1 > maxEnergy {
			maxEnergy = len(energized) - 1
		}

		energized = map[point]struct{}{}
		x++
	}

	fmt.Println(maxEnergy)

	energized = map[point]struct{}{}

	// from right
	y := 0
	for {
		if y >= len(maze) {
			break
		}

		start := point{x: -1, y: y}
		queue := []*beam{
			{heading: point{x: 1, y: 0}, current: start, visited: map[impact]struct{}{}},
		}

		// we stop when no beams can move
		var current *beam
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]

			energized[current.current] = struct{}{}
			// if _, ok := visited[impact{location: current.current, fromDirection: current.heading}]; ok {
			// 	// we already reached this point by others...
			// 	break
			// }

			visited[impact{location: current.current, fromDirection: current.heading}] = struct{}{}

			queue = append(queue, moveBeam(maze, current)...)
		}

		if len(energized)-1 > maxEnergy {
			maxEnergy = len(energized) - 1
		}

		energized = map[point]struct{}{}
		y++
	}

	fmt.Println(maxEnergy)

	energized = map[point]struct{}{}

	// start = point{x: -1, y: 0}
	y = 0
	for {
		if y >= len(maze) {
			break
		}

		start := point{x: len(maze[y]), y: y}
		queue := []*beam{
			{heading: point{x: -1, y: 0}, current: start, visited: map[impact]struct{}{}},
		}

		// we stop when no beams can move
		var current *beam
		for len(queue) > 0 {
			current, queue = queue[0], queue[1:]

			energized[current.current] = struct{}{}
			// if _, ok := visited[impact{location: current.current, fromDirection: current.heading}]; ok {
			// 	// we already reached this point by others...
			// 	break
			// }

			visited[impact{location: current.current, fromDirection: current.heading}] = struct{}{}

			queue = append(queue, moveBeam(maze, current)...)
		}

		if len(energized) > maxEnergy {
			maxEnergy = len(energized) - 1
		}

		energized = map[point]struct{}{}
		y++
	}

	fmt.Println(maxEnergy)
}

// we should probably not track individual beams.
func moveBeam(maze [][]rune, b *beam) []*beam {
	// out of bounds
	newPoint := point{x: b.current.x + b.heading.x, y: b.current.y + b.heading.y}
	if newPoint.x < 0 || newPoint.y < 0 || newPoint.y >= len(maze) || newPoint.x >= len(maze[newPoint.y]) {
		return nil
	}

	// we hit the same mirror in the same direction twice (loop)
	if _, ok := b.visited[impact{
		fromDirection: b.heading,
		location:      newPoint,
	}]; ok {
		return nil
	}

	fn := mirrors[maze[newPoint.y][newPoint.x]]
	newBeams := fn(newPoint, b)

	if len(newBeams) > 1 {
		return newBeams
	}

	// Save the angle at which we found the mirror first.
	if maze[newPoint.y][newPoint.x] != '.' {
		b.visited[impact{fromDirection: b.heading, location: newPoint}] = struct{}{}
	}

	b.current = newPoint
	b.heading = newBeams[0].heading

	return []*beam{b}
}

func display(maze [][]rune) {
	for _, c := range maze {
		for _, v := range c {
			fmt.Print(string(v))
		}

		fmt.Println()
	}
}
