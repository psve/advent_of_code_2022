package main

import (
	"testing"
)

func TestTotalPriority(t *testing.T) {
	res := totalPriority("./input_test")
	if res != 157 {
		t.Errorf("wrong result")
	}
}

func TestBadgePriority(t *testing.T) {
	res := badgePriority("./input_test")
	if res != 70 {
		t.Errorf("wrong result")
	}
}
