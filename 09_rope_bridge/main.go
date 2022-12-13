package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

func sign(x int) int {
	if x < 0 {
		return -1
	}
	return 1
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type location struct {
	x, y int
}

type set map[location]any

type rope struct {
	knots []location
}

func newRope(knots int) rope {
	return rope{make([]location, knots)}
}

func (r *rope) moveHead(move rune) {
	switch move {
	case 'R':
		r.knots[0].x++
	case 'L':
		r.knots[0].x--
	case 'U':
		r.knots[0].y++
	case 'D':
		r.knots[0].y--
	}

	for i := 1; i < len(r.knots); i++ {
		diff := location{
			x: r.knots[i-1].x - r.knots[i].x,
			y: r.knots[i-1].y - r.knots[i].y,
		}

		// If head is the same row/column as the tail, move in that direction.
		if diff.x == 0 && abs(diff.y) > 1 {
			r.knots[i].y += sign(diff.y)
			continue
		}
		if diff.y == 0 && abs(diff.x) > 1 {
			r.knots[i].x += sign(diff.x)
			continue
		}

		// If head is diagonal but not touching, move one step diagonally.
		if abs(diff.x) > 1 || abs(diff.y) > 1 {
			r.knots[i].y += sign(diff.y)
			r.knots[i].x += sign(diff.x)
		}
	}
}

func moveRope(input string, knots int) int {
	locations := make(set)
	r := newRope(knots)
	helper.ForEachLine(input, func(line string) error {
		parts := strings.Split(line, " ")
		d := []rune(parts[0])[0]
		n, _ := strconv.Atoi(parts[1])
		for i := 0; i < n; i++ {
			r.moveHead(d)
			// fmt.Printf("%s => %v\n", string(d), r.knots)
			locations[r.knots[knots-1]] = nil
		}
		return nil
	})
	return len(locations)
}

func main() {
	fmt.Println(moveRope("./input", 2))
	fmt.Println("---------")
	fmt.Println(moveRope("./input", 10))
}
