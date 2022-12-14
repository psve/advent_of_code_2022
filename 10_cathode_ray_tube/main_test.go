package main

import (
	"testing"
)

func TestSignalStrength(t *testing.T) {
	res := signalStrength("./input_test")
	if res != 13140 {
		t.Errorf("wrong result %d", res)
	}
}

func TestDraw(t *testing.T) {
	res := draw("./input_test")
	expected := `##..##..##..##..##..##..##..##..##..##..
###...###...###...###...###...###...###.
####....####....####....####....####....
#####.....#####.....#####.....#####.....
######......######......######......####
#######.......#######.......#######.....`
	if res != expected {
		t.Errorf("wrong result")
	}
}
