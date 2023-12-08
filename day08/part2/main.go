package main

import (
	"fmt"
	"os"
	"regexp"
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
	steps := ""
	match := regexp.MustCompile(`(\w+) = \((\w+), (\w+)\)`)
	nodes := make(map[string]*node)
	for i, l := range split {
		if i == 0 {
			steps = l

			continue
		}

		if l == "" {
			continue
		}

		var (
			n, left, right string
		)

		matches := match.FindAllStringSubmatch(l, -1)

		m := matches[0]
		n = m[1]
		left = m[2]
		right = m[3]

		if v, ok := nodes[n]; ok {
			v.left = &node{value: left}
			v.right = &node{value: right}
		} else {
			nodes[n] = &node{value: n, right: &node{value: right}, left: &node{value: left}}
		}
	}

	movingNodes := []*node{}
	for _, n := range nodes {
		if n.value[len(n.value)-1] == 'A' {
			movingNodes = append(movingNodes, n)
		}
	}

	allSteps := make([]int, 0)
	for _, n := range movingNodes {
		n := n
		s := findSteps(n, nodes, steps)
		allSteps = append(allSteps, s)
	}

	fmt.Println(lcmList(allSteps))
}

// gather all the points at which each of the nodes reach the end.
// then the LCD of those numbers is what we are looking for. This is like a frequency count.
// The numbers obviously move in a pattern. That is the indicator that you don't want to loop and try and bruteforce
// the result.
func findSteps(current *node, nodes map[string]*node, steps string) int {
	index := 0
	taken := 0
	for {
		if index >= len(steps) {
			index = 0
		}
		nextStep := steps[index]
		index++
		taken++

		if nextStep == 'L' {
			current = nodes[current.left.value]
		} else {
			current = nodes[current.right.value]
		}

		if current.value[len(current.value)-1] == 'Z' {
			return taken
		}
	}
}

func lcmList(nums []int) int {
	if len(nums) == 2 {
		return lcm(nums[0], nums[1])
	}

	return lcm(nums[0], lcmList(nums[1:]))
}

// LCM (a,b) = (a x b)/GCD(a,b)
func lcm(a, b int) int {
	return a * b / gcd(a, b)
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}

	return gcd(b, a%b)
}
