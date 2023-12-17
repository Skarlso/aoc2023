package main

import (
	"testing"
)

func TestHash(t *testing.T) {
	s := "HASH"
	sum := 0
	for _, r := range s {
		sum += int(r)
		sum *= 17
		sum %= 256
	}

	if sum != 52 {
		t.Fatalf("result was: %d", sum)
	}
}
