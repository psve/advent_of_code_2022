package main

import (
	"testing"
)

func TestCheckRow(t *testing.T) {
	sensors := parse("./input_test")
	res := len(checkRow(sensors, 10))
	if res != 26 {
		t.Errorf("wrong result %d", res)
	}
}

func TestFindBeacon(t *testing.T) {
	sensors := parse("./input_test")
	res := findBeacon(sensors, 0, 20)
	if res != 56000011 {
		t.Errorf("wrong result %d", res)
	}
}
