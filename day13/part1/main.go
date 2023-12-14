package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")

	mazes := [][][]rune{}
	maze := [][]rune{}
	for _, l := range split {
		if l == "" {
			mazes = append(mazes, maze)
			maze = [][]rune{}
			continue
		}

		maze = append(maze, []rune(l))
	}

	sum := 0
	for _, m := range mazes {
		if v, ok := checkColumnMirror(m); ok {
			sum += v
		} else if v, ok := checkRowMirror(m); ok {
			sum += v
		}
	}

	fmt.Println(sum)
}

func checkRowMirror(maze [][]rune) (int, bool) {
	for y := 0; y < len(maze); y++ {
		if checkIsMirroredFromRow(y, maze) {
			return (y + 1) * 100, true
		}
	}

	return 0, false
}

func checkIsMirroredFromRow(y int, maze [][]rune) bool {
	for up, down := y, y+1; ; up, down = up-1, down+1 {
		if up < 0 {
			break
		}

		if down >= len(maze) {
			break
		}

		if string(maze[up]) != string(maze[down]) {
			return false
		}
	}

	return true
}

func checkColumnMirror(maze [][]rune) (int, bool) {
	for x := 0; x < len(maze[0]); x++ {
		if x+1 < len(maze[0]) {
			if checkIsMirroredFromColumn(x, maze) {
				return x + 1, true
			}
		}
	}

	return 0, false
}

func checkIsMirroredFromColumn(index int, maze [][]rune) bool {
	match := true
	for y := 0; y < len(maze); y++ {
		// Maybe string(maze[y])[0:index ] == string(maze[y])[index+1:] -> but you'd have to cut it
		// in half based on lenght of another since they can get out of bounds...
		for left, right := index, index+1; ; left, right = left-1, right+1 {
			// if we reached the end of any of the loops, break
			if left < 0 {
				break
			}

			if right >= len(maze[y]) {
				break
			}

			if maze[y][left] != maze[y][right] {
				match = false

				break
			}
		}

		if !match {
			break
		}
	}

	return match
}

func display(mazes [][][]rune) {
	for _, maze := range mazes {
		for _, v := range maze {
			for _, row := range v {
				fmt.Print(string(row))
			}

			fmt.Println()
		}

		fmt.Println()
	}
}
