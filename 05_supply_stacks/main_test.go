package main

import (
	"testing"
)

func TestCrateMover9000(t *testing.T) {
	res := crateMover9000("./input_test")
	if res != "CMZ" {
		t.Errorf("wrong result")
	}
}

func TestCrateMover9001(t *testing.T) {
	res := crateMover9001("./input_test")
	if res != "MCD" {
		t.Errorf("wrong result")
	}
}
