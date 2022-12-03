package main

import (
	"testing"
)

func TestFromMoves(t *testing.T) {
	res := totalScore("./input_test", fromMoves)
	if res != 15 {
		t.Errorf("wrong result")
	}
}

func TestFromStrategy(t *testing.T) {
	res := totalScore("./input_test", fromStrategy)
	if res != 12 {
		t.Errorf("wrong result")
	}
}
