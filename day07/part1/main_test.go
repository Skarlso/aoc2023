package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTypes(t *testing.T) {
	testcases := []struct {
		name         string
		expectedType int
		hand         []rune
	}{
		{
			name:         "four of a kind",
			hand:         []rune{'A', 'A', '8', 'A', 'A'},
			expectedType: FourOfAKind,
		},
		{
			name:         "five of a kind",
			hand:         []rune{'A', 'A', 'A', 'A', 'A'},
			expectedType: FiveOfAKind,
		},
		{
			name:         "three of a kind",
			hand:         []rune{'A', 'A', 'A', '8', '9'},
			expectedType: ThreeOfAKind,
		},
		{
			name:         "two pair",
			hand:         []rune{'2', '2', '3', '3', 'A'},
			expectedType: TwoPair,
		},
		{
			name:         "two pair",
			hand:         []rune{'2', '8', '2', '6', '8'},
			expectedType: TwoPair,
		},
		{
			name:         "on pair",
			hand:         []rune{'A', '2', '3', 'A', '4'},
			expectedType: OnePair,
		},
		{
			name:         "high card",
			hand:         []rune{'A', 'T', '2', '3', '4'},
			expectedType: HighCard,
		},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()
			// test that everything else IS FALSE
			for k, v := range types {
				if k == tt.expectedType {
					assert.True(t, v(tt.hand), "expected type %d to be true", k)
				} else {
					assert.False(t, v(tt.hand), "expected type %d to be false", k)
				}
			}
		})
	}
}

func TestCompare(t *testing.T) {
	testcases := []struct {
		name    string
		hand    hand
		other   hand
		outcome int
	}{
		{
			name:    "five beats four",
			hand:    hand{cards: []rune{'A', 'A', 'A', 'A', 'A'}},
			other:   hand{cards: []rune{'A', 'A', 'A', 'A', 'T'}},
			outcome: 1,
		},
		{
			name:    "33332 beats 2AAAA",
			hand:    hand{cards: []rune{'3', '3', '3', '3', '2'}},
			other:   hand{cards: []rune{'2', 'A', 'A', 'A', 'A'}},
			outcome: 1,
		},
		{
			name:    "77888 beats 77788",
			hand:    hand{cards: []rune{'7', '7', '8', '8', '8'}},
			other:   hand{cards: []rune{'7', '7', '7', '8', '8'}},
			outcome: 1,
		},
	}

	for _, tt := range testcases {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Helper()

			assert.Equal(t, tt.outcome, tt.hand.Compare(tt.other))
		})
	}
}
