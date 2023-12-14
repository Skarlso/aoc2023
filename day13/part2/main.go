package main

import (
	"fmt"
	"os"
	"slices"
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
	// With a Smudge, one match will always be true.
	// If there are no more smudges left, we say we didn't match.
	for _, m := range mazes {
		smudge := 1
		s := 0

		if v, ok := checkRowMirror(m, smudge); ok {
			s = v
		} else if v, ok := checkColumnMirror(m, smudge); ok {
			s = v
		}

		sum += s
	}

	fmt.Println(sum)
}

// TODO: Transpose it instead so we can just do a vertical search again.
func checkRowMirror(maze [][]rune, smudge int) (int, bool) {
	for y := 0; y < len(maze); y++ {

		if y+1 < len(maze) {
			if checkIsMirroredFromRow(y, maze, smudge) {
				fmt.Println("found mirrored row at: ", y+1, y+2)

				return (y + 1) * 100, true
			}
		}
	}

	return 0, false
}

func checkIsMirroredFromRow(y int, maze [][]rune, smudge int) bool {
	for up, down := y, y+1; ; up, down = up-1, down+1 {
		if up < 0 {
			break
		}

		if down >= len(maze) {
			break
		}

		// we match each character instead to use up a smudge.
		// if we used up a smudge we can't use any more.
		if v := slices.CompareFunc(maze[up], maze[down], func(r1, r2 rune) int {
			if smudge == 1 && r1 != r2 {
				smudge = 0
				return 0
			}

			return int(r1) - int(r2)
		}); v != 0 {
			return false
		}
		// for left, right :=
		// if string(maze[up]) != string(maze[down]) {
		// 	return false
		// }
	}

	return true
}

func checkColumnMirror(maze [][]rune, smudge int) (int, bool) {
	for x := 0; x < len(maze[0]); x++ {
		if x+1 < len(maze[0]) {
			if checkIsMirroredFromColumn(x, maze, smudge) {
				fmt.Println("mirror point at column at: ", x+1, x+2)
				return x + 1, true
			}
		}
	}

	return 0, false
}

func checkIsMirroredFromColumn(index int, maze [][]rune, smudge int) bool {
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

			if smudge == 1 && maze[y][left] != maze[y][right] {
				smudge = 0
			} else if maze[y][left] != maze[y][right] {
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
