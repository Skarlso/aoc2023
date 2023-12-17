package main

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type point struct {
	x, y int
}

// func (p point) subtract(other point) point {
// 	return point{x: p.x - other.x, y: p.y - other.y}
// }

// var directions = []point{
// 	// left
// 	{x: -1, y: 0},
// 	// right
// 	{x: 1, y: 0},
// 	// down
// 	{x: 0, y: 1},
// 	// up
// 	{x: 0, y: -1},
// }

var mirrors = map[rune]func(current point, b *beam) []*beam{
	'/': func(current point, b *beam) []*beam {
		newHeading := point{x: -b.heading.y, y: -b.heading.x}
		return []*beam{
			{
				heading: newHeading,
				// current: point{x: current.x + newHeading.x, y: current.y + newHeading.y},
			},
		}
	},
	'\\': func(current point, b *beam) []*beam {
		newHeading := point{x: b.heading.y, y: b.heading.x}
		return []*beam{
			{
				heading: newHeading,
				// current: point{x: current.x + newHeading.x, y: current.y + newHeading.y},
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
					visited: map[point]struct{}{},
				},
				{
					heading: point{x: 0, y: 1},
					// current: point{x: current.x, y: current.y + 1},
					current: point{x: current.x, y: current.y},
					visited: map[point]struct{}{},
				},
			}
		case point{x: 0, y: 1}, point{x: 0, y: -1}:
			return []*beam{
				{
					heading: b.heading,
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
					visited: map[point]struct{}{},
				},
				{
					heading: point{x: 1, y: 0},
					// current: point{x: current.x + 1, y: current.y},
					current: point{x: current.x, y: current.y},
					visited: map[point]struct{}{},
				},
			}
		case point{x: 1, y: 0}, point{x: -1, y: 0}:
			return []*beam{
				{
					heading: b.heading,
					// current: point{x: current.x + b.heading.x, y: current.y + b.heading.y},
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
			},
		}
	},
	// this is my mark
	'o': func(current point, b *beam) []*beam {
		return []*beam{
			{
				// current: current,
				heading: b.heading,
			},
		}
	},
}

type beam struct {
	heading point
	current point
	visited map[point]struct{}
	dead    bool
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

	// queue := []point{{x: 0, y: 0}}
	// var current point
	// energized := map[point]struct{}{}
	start := point{x: 0, y: 0}
	beams := []*beam{
		{heading: point{x: 1, y: 0}, current: start, visited: map[point]struct{}{}},
	}

	// // note: there is no visited because the lights path will just die out.
	// for len(queue) > 0 {
	// 	current, queue = queue[0], queue[1:]
	// 	maze[current.y][current.x] = 'o'

	// 	for _, next := range path(maze, heading, current) {
	// 		if _, ok := energized[next]; !ok {
	// 			energized[next] = struct{}{}
	// 			queue = append(queue, next)
	// 		}
	// 	}

	// 	display(maze)
	// 	fmt.Println()
	// 	time.Sleep(500 * time.Millisecond)
	// }

	// we stop when no beams can move
	energized := map[point]struct{}{}
	for {
		// moved := false
		for _, b := range beams {
			// fmt.Println(b)
			// only add new beans, current beams should be terminated or removed.
			_, newBeans := moveBeam(maze, b)

			// if !moved {
			// 	moved = update
			// }

			// fmt.Println()
			// time.Sleep(200 * time.Millisecond)
			// if update {
			// display(maze, b.current)
			// maze[b.current.y][b.current.x] = 'o'
			// }
			energized[b.current] = struct{}{}
			beams = append(beams, newBeans...)
		}

		allDead := true
		for _, b := range beams {
			// if ! {
			if allDead && !b.dead {
				fmt.Println("not dead: ", b)
				allDead = b.dead
			}
			// }
		}

		if allDead {
			break
		}

		// if !moved {
		// break
		// }
	}

	sum := 0
	for _, b := range beams {
		sum += len(b.visited)
	}

	fmt.Println("sum: ", sum)
}

// we should probably not track individual beams.
func moveBeam(maze [][]rune, b *beam) (bool, []*beam) {
	if b.dead {
		return false, nil
	}

	// out of bounds
	newPoint := point{x: b.current.x + b.heading.x, y: b.current.y + b.heading.y}
	if newPoint.x < 0 || newPoint.y < 0 || newPoint.y >= len(maze) || newPoint.x >= len(maze[newPoint.y]) {
		b.dead = true

		return false, nil
	}

	fn := mirrors[maze[newPoint.y][newPoint.x]]
	newBeams := fn(newPoint, b)

	if len(newBeams) > 1 {
		b.dead = true

		return true, newBeams
	}

	// Don't update current because the next thing can also be a mirror
	// and if we skip that, we are toast.
	if len(newBeams) == 0 {
		log.Fatal("no new beams at: ", newPoint, string(maze[newPoint.y][newPoint.x]))
	}
	b.current = newPoint
	// fmt.Println(newPoint.x, newPoint.y)
	b.heading = newBeams[0].heading
	b.visited[newPoint] = struct{}{}
	// maze[b.current.y][b.current.x] = 'o'

	return true, nil
}

func display(maze [][]rune, current point) {
	for y, c := range maze {
		for x, v := range c {
			if x == current.x && y == current.y {
				fmt.Print("o")
			} else {
				fmt.Print(string(v))
			}
		}

		fmt.Println()
	}
}

// func path(maze [][]rune, heading, p point) []point {
// 	var result []point
// 	for _, d := range directions {
// 		np := point{x: p.x + d.x, y: p.y + d.y}
// 		if np.x < 0 || np.y < 0 || np.y >= len(maze) || np.x >= len(maze[np.y]) {
// 			continue
// 		}

// 		fn := mirrors[maze[np.y][np.x]]
// 		points := fn(np, p)

// 		for _, pt := range points {
// 			if pt.x < 0 || pt.y < 0 || pt.y >= len(maze) || pt.x >= len(maze[pt.y]) {
// 				continue
// 			}

// 			result = append(result, pt)
// 		}
// 	}

// 	return result
// }
