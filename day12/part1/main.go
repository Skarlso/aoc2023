package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/Skarlso/cache"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	c := cache.New[int]()
	result := 0
	for _, l := range split {
		arg := strings.Split(l, " ")
		conditions, rules := arg[0], arg[1]
		rulesList := strings.Split(rules, ",")
		result += c.WithCache(calculate(conditions, rulesList), conditions, rulesList)
	}

	fmt.Println("result: ", result)
}

func calculate(conditions string, rules []string) cache.Cacheable[int] {
	return func() int {
		//if v, ok := cache[args{conditions: conditions, rules: strings.Join(rules, ",")}]; ok {
		//	return v
		//}

		if len(rules) == 0 {
			if strings.Contains(conditions, "#") {
				//cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 0
				return 0
			}

			//cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 1
			return 1
		}
		if len(conditions) == 0 {
			if len(rules) == 0 {
				//cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 1
				return 1
			}

			//cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = 0
			return 0
		}

		var result int

		if conditions[0] == '.' || conditions[0] == '?' {
			result += calculate(conditions[1:], rules)()
		}

		if conditions[0] == '#' || conditions[0] == '?' {
			index, _ := strconv.Atoi(rules[0])
			if index <= len(conditions) && !strings.Contains(conditions[:index], ".") && (index == len(conditions) || conditions[index] != '#') {
				if index+1 >= len(conditions) {
					result += calculate("", rules[1:])()
				} else {
					result += calculate(conditions[index+1:], rules[1:])()
				}
			}
		}

		//cache[args{conditions: conditions, rules: strings.Join(rules, "")}] = result
		return result
	}
}
