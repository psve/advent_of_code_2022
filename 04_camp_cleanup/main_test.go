package main

import (
	"testing"
)

func TestFullyContains(t *testing.T) {
	res := fullyContains("./input_test")
	if res != 2 {
		t.Errorf("wrong result")
	}
}

func TestOverlaps(t *testing.T) {
	res := overlaps("./input_test")
	if res != 4 {
		t.Errorf("wrong result")
	}
}
