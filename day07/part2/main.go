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

var cardFaces = []rune{'A', 'K', 'Q', 'T', '9', '8', '7', '6', '5', '4', '3', '2', 'J'}

// uh this is stupid
// but the runes are different.
var cardValue = map[rune]int{
	'A': 13,
	'K': 12,
	'Q': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
	'J': 1,
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
	typ := 0

	for k, v := range types {
		if v(h.cards) {
			typ = k
			break
		}
	}

	if !strings.Contains(string(h.cards), "J") {
		return typ
	}

	// no point to upscale if there is no better
	if typ == FiveOfAKind {
		return typ
	}

	// backtracking for all types that are greater than the current one.
	for t := 0; t <= typ; t++ {
		// don't modify the original
		cards := make([]rune, len(h.cards))
		copy(cards, h.cards)

		// we start by changing the first occurrence of J.
		// gather all J indexes.
		indexes := []int{}
		for i, c := range cards {
			if c == 'J' {
				indexes = append(indexes, i)
				cards[i] = 'A'
			}
		}
		if tryType(t, cards, indexes, 0) {
			return t
		}
	}

	return -1
}

func tryType(typ int, hand []rune, indexes []int, index int) bool {
	for _, c := range cardFaces {
		hand[indexes[index]] = c

		// try the current one. If we reached the end, we keep trying until we eventually fail.
		// if we fail, we return with failed. If we can increase we increase.
		// if we can increase we increase -> meaning we start at the end and increase until the
		// for loop reaches its end. We return false from there. That return then will continue where
		// it left off and increase the card again and increase the index again. Effectively walking through
		// all the combinations.
		if index+1 >= len(indexes) {
			if types[typ](hand) {
				return true
			}
		} else {
			if tryType(typ, hand, indexes, index+1) {
				return true
			}
		}
	}

	return false
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
