package main

import (
	"testing"
)

func TestGetSizes(t *testing.T) {
	res := getSizes("./input_test")
	if res != 95437 {
		t.Errorf("wrong result")
	}
}

func TestDeleteDirectory(t *testing.T) {
	res := deleteDirectory("./input_test")
	if res != 24933642 {
		t.Errorf("wrong result %d", res)
	}
}
