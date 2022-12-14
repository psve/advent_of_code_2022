package main

import (
	"testing"
)

func TestMonkeyBusinessOne(t *testing.T) {
	res := monkeyBusiness("./input_test", true, 20)
	if res != 10605 {
		t.Errorf("wrong result %d", res)
	}
}

func TestMonkeyBusinessTwo(t *testing.T) {
	res := monkeyBusiness("./input_test", false, 10000)
	if res != 2713310158 {
		t.Errorf("wrong result %d", res)
	}
}
