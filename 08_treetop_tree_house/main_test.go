package main

import "testing"

func TestCountVisible(t *testing.T) {
	res := countVisible("./input_test")
	if res != 21 {
		t.Errorf("wrong result")
	}
}

func TestBestScenicScore(t *testing.T) {
	res := bestScenicScore("./input_test")
	if res != 8 {
		t.Errorf("wrong result")
	}
}
