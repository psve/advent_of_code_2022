package main

import (
	"testing"
)

func TestMost(t *testing.T) {
	res := most("./input_test")
	if res != 24000 {
		t.Errorf("wrong result")
	}
}

func TestTopThree(t *testing.T) {
	res := topThree("./input_test")
	if res != 45000 {
		t.Errorf("wrong result")
	}
}
