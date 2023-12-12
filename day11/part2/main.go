package main

import (
	"fmt"
	"os"
	"strings"
)

const part1Result = 10228230
const part1TestResult = 374

type point struct {
	x, y int
}

type galaxy struct {
	location point
	pair     map[point]struct{}
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	// while looping through the rows, we can already expand them

	// if a galaxy is already located in the pair of a galaxy that it's trying to
	// pair up with don't pair it twice
	maze := [][]rune{}
	for _, l := range split {
		maze = append(maze, []rune(l))
		if !strings.Contains(l, "#") {
			maze = append(maze, []rune(l))
			maze = append(maze, []rune(l))
		}
	}

	// now, expand columns
	for x := 0; x < len(maze[0]); x++ {
		containsGalaxy := false
		for y := 0; y < len(maze); y++ {
			if maze[y][x] == '#' {
				containsGalaxy = true
			}
		}

		if !containsGalaxy {
			for y := 0; y < len(maze); y++ {
				maze[y] = appendToSlice(x, '.', maze[y])
				maze[y] = appendToSlice(x, '.', maze[y])
				// skip the added column otherwise we would be expanding indefinitely.
			}

			// skip it only once.
			x++
			x++
		}
	}

	// second task, pair up galaxies.
	// if a galaxy is already paired with a galaxy, don't pair it twice
	// probably could be done in one step...
	galaxies := make(map[point]galaxy)
	for y := 0; y < len(maze); y++ {
		for x := 0; x < len(maze[y]); x++ {
			if maze[y][x] == '#' {
				p := point{y: y, x: x}
				galaxies[p] = galaxy{location: p, pair: map[point]struct{}{}}
			}
		}
	}

	pairs := 0
	for ck, cv := range galaxies {
		for ok, ov := range galaxies {
			if ck == ok {
				continue
			}

			if _, ok := ov.pair[ck]; !ok {
				cv.pair[ov.location] = struct{}{}
				pairs++
			}
		}
	}

	fmt.Println("pairs: ", pairs)

	sum := 0
	for _, v := range galaxies {
		for pair := range v.pair {
			sum += manhattan(v.location, pair)
			// fmt.Println("done: ", pairsDone)
		}
	}

	fmt.Println("sum: ", sum)
	fmt.Println("testDiff: ", sum-part1TestResult)
	fmt.Println("testResultWith100Iterations: ", part1TestResult+98*(sum-part1TestResult))
	fmt.Println("diff for a mil: ", part1Result+999998*(sum-part1Result))
}

func manhattan(from, to point) int {
	return abs(from.x-to.x) + abs(from.y-to.y)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}

	return x
}

func appendToSlice(i int, a rune, slice []rune) []rune {
	if i+1 == len(slice) {
		return append(slice, a)
	}

	slice = append(slice[:i+1], slice[i:]...)
	slice[i] = '.'

	return slice
}

func display(maze [][]rune) {
	for _, c := range maze {
		for _, r := range c {
			fmt.Print(string(r))
		}

		fmt.Println()
	}
}
