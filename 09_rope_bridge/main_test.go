package main

import (
	"testing"
)

func TestMoveShortRope(t *testing.T) {
	res := moveRope("./input_1_test", 2)
	if res != 13 {
		t.Errorf("wrong result %d", res)
	}
}

func TestMoveLongRope1(t *testing.T) {
	res := moveRope("./input_1_test", 10)
	if res != 1 {
		t.Errorf("wrong result %d", res)
	}
}

func TestMoveLongRope2(t *testing.T) {
	res := moveRope("./input_2_test", 10)
	if res != 36 {
		t.Errorf("wrong result %d", res)
	}
}
