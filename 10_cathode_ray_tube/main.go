package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

type cpu struct {
	x, cycle int
	pipeline []int
}

func newCPU(path string) cpu {
	c := cpu{x: 1, pipeline: make([]int, 0)}
	helper.ForEachLine(path, func(line string) error {
		parts := strings.Split(line, " ")
		switch len(parts) {
		case 1:
			c.addInstruction(0)
		case 2:
			val, _ := strconv.Atoi(parts[1])
			c.addInstruction(val)
		}
		return nil
	})
	return c
}

// Add an instruction. 0 represents a noop, all other inputs represent an addx.
func (c *cpu) addInstruction(i int) {
	switch i {
	case 0:
		c.pipeline = append(c.pipeline, 0)
	default:
		// Simulate an add by first doing a noop and then doing the actual add.
		c.pipeline = append(c.pipeline, 0)
		c.pipeline = append(c.pipeline, i)
	}
}

// Run the next step in the pipeline, returning the old value of the register and the cycle count.
func (c *cpu) step() (int, int, error) {
	if len(c.pipeline) == 0 {
		return 0, c.cycle, fmt.Errorf("no instructions left")
	}
	out := c.x
	c.x += c.pipeline[0]
	c.cycle++
	c.pipeline = c.pipeline[1:]
	return out, c.cycle, nil
}

func signalStrength(path string) int {
	c := newCPU(path)
	var strength int
	for x, cycle, err := c.step(); err == nil; x, cycle, err = c.step() {
		if (cycle-20)%40 == 0 {
			strength += cycle * x
		}
	}

	return strength
}

func draw(path string) string {
	c := newCPU(path)
	var monitor string
	for x, cycle, err := c.step(); err == nil; x, cycle, err = c.step() {
		pos := (cycle - 1) % 40
		if cycle != 1 && pos == 0 {
			monitor += "\n"
		}
		if x-1 == pos || x == pos || x+1 == pos {
			monitor += "#"
		} else {
			monitor += "."
		}
	}
	return monitor
}

func main() {
	fmt.Println(signalStrength("./input"))
	fmt.Println("---------")
	fmt.Println(draw("./input"))
}
