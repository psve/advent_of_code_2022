package main

import (
	"testing"
)

func TestDropSandVoid(t *testing.T) {
	res := dropSand("./input_test", false)
	if res != 24 {
		t.Errorf("wrong result %d", res)
	}
}

func TestDropSandFloor(t *testing.T) {
	res := dropSand("./input_test", true)
	if res != 93 {
		t.Errorf("wrong result %d", res)
	}
}
