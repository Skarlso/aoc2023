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
	len    int
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
			if unicode.IsNumber(c) {
				num += string(c)
				if !lookAheadIsNumber(l, x) {
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
					})
					num = ""
				}
			}

			row = append(row, c)
		}

		matrix = append(matrix, row)
	}

	// sum := 0
	gears := make(map[point][]number)
	for _, n := range numbers {
		registerNumberForGear(matrix, n, gears)
	}

	sum := 0
	for _, v := range gears {
		if len(v) == 2 {
			sum += v[0].value * v[1].value
		}
	}

	fmt.Println(sum)
}

func registerNumberForGear(matrix [][]rune, n number, gears map[point][]number) {
	for _, p := range n.border {
		if p.x < 0 || p.y < 0 || p.y >= len(matrix) || p.x >= len(matrix[p.y]) {
			continue
		}

		if matrix[p.y][p.x] == '*' {
			gears[point{y: p.y, x: p.x}] = append(gears[point{y: p.y, x: p.x}], n)
		}
	}
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
