package main

import (
	"testing"
)

func TestClimbHill(t *testing.T) {
	res := climbHill("./input_test")
	if res != 31 {
		t.Errorf("wrong result %d", res)
	}
}

func TestHikeTrail(t *testing.T) {
	res := hikeTrail("./input_test")
	if res != 29 {
		t.Errorf("wrong result %d", res)
	}
}
