package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
)

type point struct {
	x, y int
}

// a...b
// .....
// c...d
type rectangle struct {
	a, b, c, d point
}

type number struct {
	value  int
	border []point
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)

	split := strings.Split(string(content), "\n")
	// ids := 0
	numbers := make([]number, 0)
	matrix := [][]rune{}
	for y, l := range split {
		row := make([]rune, 0, len(l))
		var num string
		for x, c := range l {
			// is number or `-`
			// if unicode.IsNumber(c) || c == '-' {
			if unicode.IsNumber(c) {
				num += string(c)
				if !lookAheadIsNumber(l, x) {
					// fmt.Println("num: ", num)
					n, _ := strconv.Atoi(num)
					points := make([]point, 0)
					for j := y - 1; j <= y+1; j++ {
						for i := x - len(num); i <= x+1; i++ {
							points = append(points, point{y: j, x: i})
						}
					}
					numbers = append(numbers, number{
						value:  n,
						border: points,
						// container: rectangle{
						// 	a: point{x: x - len(num), y: y - 1},
						// 	b: point{x: x + 1, y: y - 1},
						// 	c: point{x: x - len(num), y: y + 1},
						// 	d: point{x: x + 1, y: y + 1},
						// },
					})
					num = ""
				}
			}

			row = append(row, c)
		}

		matrix = append(matrix, row)
	}

	sum := 0
	for _, n := range numbers {
		if isSymbolOnBorder(matrix, n) {
			sum += n.value
		}
	}

	fmt.Println(sum)
}

func isSymbolOnBorder(matrix [][]rune, n number) bool {
	for _, p := range n.border {
		if p.x < 0 || p.y < 0 || p.y >= len(matrix) || p.x >= len(matrix[p.y]) {
			continue
		}

		if !unicode.IsNumber(matrix[p.y][p.x]) && matrix[p.y][p.x] != '.' {
			return true
		}
	}

	return false
}

func lookAheadIsNumber(l string, x int) bool {
	if x+1 >= len(l) {
		return false
	}

	return unicode.IsNumber(rune(l[x+1]))
}

func print(matrix [][]string) {
	for _, col := range matrix {
		for _, row := range col {
			fmt.Print(row)
		}
		fmt.Println()
	}
}
