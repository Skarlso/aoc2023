package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// args for the function to cache.
type args struct {
	conditions string
	rules      string
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	cache := make(map[args]int64)
	var result int64
	for _, l := range split {
		arg := strings.Split(l, " ")
		conditions, rules := arg[0], arg[1]
		var (
			conds string
			rs    string
		)
		for i := 0; i < 5; i++ {
			conds += conditions + "?"
			rs += rules + ","
		}
		conditions = conds[:len(conds)-1]
		rulesList := strings.Split(rs[:len(rs)-1], ",")
		// fmt.Println(conditions, rulesList)

		// fmt.Println(conditions, rulesList)
		result += calculate(cache, conditions, rulesList)
	}

	fmt.Println("result:", result)
}

func calculate(cache map[args]int64, conditions string, rules []string) int64 {
	// fmt.Println(conditions, rules)
	if v, ok := cache[args{conditions: conditions, rules: strings.Join(rules, "")}]; ok {
		return v
	}

	if len(rules) == 0 {
		if strings.Contains(conditions, "#") {
			cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 0
			return 0
		}

		cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 1
		return 1
	}
	if len(conditions) == 0 {
		if len(rules) == 0 {
			cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 1
			return 1
		}

		cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 0
		return 0
	}

	var result int64

	if conditions[0] == '.' || conditions[0] == '?' {
		result += calculate(cache, conditions[1:], rules)
	}

	if conditions[0] == '#' || conditions[0] == '?' {
		index, _ := strconv.Atoi(rules[0])
		if index <= len(conditions) && !strings.Contains(conditions[:index], ".") && (index == len(conditions) || conditions[index] != '#') {
			if index+1 >= len(conditions) {
				result += calculate(cache, "", rules[1:])
			} else {
				result += calculate(cache, conditions[index+1:], rules[1:])
			}
		}
	}

	cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = result
	return result
}
