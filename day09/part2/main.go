package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

type node struct {
	left  *node
	right *node
	value string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	sequences := make([][]int, 0)
	for _, l := range split {
		if l == "" {
			continue
		}

		nums := make([]int, 0)
		numbers := strings.Split(l, " ")
		for _, n := range numbers {
			num, _ := strconv.Atoi(n)
			nums = append(nums, num)
		}

		sequences = append(sequences, nums)
	}

	// fmt.Println(sequences)

	sum := 0
	for _, seq := range sequences {
		pyramid := buildPyramid(seq)
		// fmt.Println(pyramid)
		sum += extrapolate(pyramid)
	}

	fmt.Println(sum)
}

func extrapolate(pyramid [][]int) int {
	// add the first 0
	pyramid[len(pyramid)-1] = append([]int{0}, pyramid[len(pyramid)-1]...)
	// continue from the second to last
	for i := len(pyramid) - 2; i >= 0; i-- {
		// previous row's last value + current row's last value
		v := pyramid[i][0] - pyramid[i+1][0]
		pyramid[i] = append([]int{v}, pyramid[i]...)
	}

	// fmt.Println(pyramid)
	return pyramid[0][0]
}

func buildPyramid(seq []int) [][]int {
	var result [][]int
	result = append(result, seq)

	for {
		diff := make([]int, 0)
		allZeroes := true
		for i := 1; i < len(seq); i++ {
			v := seq[i] - seq[i-1]
			if v != 0 {
				allZeroes = false
			}

			diff = append(diff, v)
		}
		result = append(result, diff)

		if allZeroes {
			return result
		}

		seq = diff
	}
}
