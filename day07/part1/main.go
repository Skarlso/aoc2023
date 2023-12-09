package main

import (
	"fmt"
	"os"
	"sort"
	"strings"
)

const (
	FiveOfAKind = iota
	FourOfAKind
	FullHouse
	ThreeOfAKind
	TwoPair
	OnePair
	HighCard
)

// uh this is stupid
// but the runes are different.
var cardValue = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'J': 10,
	'T': 9,
	'9': 8,
	'8': 7,
	'7': 6,
	'6': 5,
	'5': 4,
	'4': 3,
	'3': 2,
	'2': 1,
}

var types = map[int]func(h []rune) bool{
	FiveOfAKind: func(h []rune) bool {
		first := h[0]
		for i := 1; i < len(h); i++ {
			if h[i] != first {
				return false
			}
		}

		return true
	},
	FourOfAKind: func(h []rune) bool {
		m := map[rune]int{}
		for _, r := range h {
			m[r]++
		}

		if len(m) != 2 {
			return false
		}

		diff := []int{}
		for _, v := range m {
			diff = append(diff, v)
		}

		sort.Ints(diff)

		return diff[0] == 1 && diff[1] == 4
	},
	FullHouse: func(h []rune) bool {
		m := map[rune]int{}
		for _, r := range h {
			m[r]++
		}

		if len(m) != 2 {
			return false
		}

		diff := []int{}
		for _, v := range m {
			diff = append(diff, v)
		}

		sort.Ints(diff)

		return diff[0] == 2 && diff[1] == 3
	},
	ThreeOfAKind: func(h []rune) bool {
		m := map[rune]int{}
		for _, r := range h {
			m[r]++
		}

		if len(m) != 3 {
			return false
		}

		diff := []int{}
		for _, v := range m {
			diff = append(diff, v)
		}

		sort.Ints(diff)

		return diff[0] == 1 && diff[1] == 1 && diff[2] == 3
	},
	TwoPair: func(h []rune) bool {
		m := map[rune]int{}
		for _, r := range h {
			m[r]++
		}

		if len(m) != 3 {
			return false
		}

		diff := []int{}
		for _, v := range m {
			diff = append(diff, v)
		}

		sort.Ints(diff)

		return diff[0] == 1 && diff[1] == 2 && diff[2] == 2
	},
	OnePair: func(h []rune) bool {
		m := map[rune]int{}
		for _, r := range h {
			m[r]++
		}

		if len(m) != 4 {
			return false
		}

		diff := []int{}
		for _, v := range m {
			diff = append(diff, v)
		}

		sort.Ints(diff)

		return diff[0] == 1 && diff[1] == 1 && diff[2] == 1 && diff[3] == 2
	},
	HighCard: func(h []rune) bool {
		m := map[rune]struct{}{}
		for _, r := range h {
			m[r] = struct{}{}
		}

		return len(m) == 5
	},
}

type hand struct {
	cards []rune
	bid   int
}

func (h *hand) Type() int {
	for k, v := range types {
		if v(h.cards) {
			return k
		}
	}

	return -1
}

// Compare:
// 1: higher
// 0: equal
// -1: lower
func (h *hand) Compare(other hand) int {
	if h.Type() < other.Type() {
		return 1
	} else if h.Type() > other.Type() {
		return -1
	}

	// the types equal, fall back to comparison.
	for i, v := range h.cards {
		if cardValue[v] > cardValue[other.cards[i]] {
			return 1
		} else if cardValue[v] < cardValue[other.cards[i]] {
			return -1
		}
	}

	return -1
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run part1/main.go [file]")
		os.Exit(1)
	}
	file := os.Args[1]

	content, _ := os.ReadFile(file)
	split := strings.Split(string(content), "\n")
	hands := make([]hand, 0)

	for _, l := range split {
		var (
			h   string
			bid int
		)

		fmt.Sscanf(l, "%s %d", &h, &bid)
		hands = append(hands, hand{
			cards: []rune(h),
			bid:   bid,
		})
	}

	sort.Slice(hands, func(i, j int) bool {
		return hands[i].Compare(hands[j]) != 1
	})

	total := 0

	for i, v := range hands {
		total += (i + 1) * v.bid
	}

	fmt.Println(total)
}
