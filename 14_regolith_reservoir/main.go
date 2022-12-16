package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

func makeLabel(x, y int) string {
	return fmt.Sprintf("(%d,%d)", x, y)
}

type scan map[string]bool

func parse(path string) (scan, int) {
	s := make(scan)
	var maxY int
	helper.ForEachLine(path, func(line string) error {
		parts := strings.Split(line, " -> ")
		for i := 0; i < len(parts)-1; i++ {
			xs, ys, _ := strings.Cut(parts[i], ",")
			xe, ye, _ := strings.Cut(parts[i+1], ",")

			xStart, _ := strconv.Atoi(xs)
			yStart, _ := strconv.Atoi(ys)
			xEnd, _ := strconv.Atoi(xe)
			yEnd, _ := strconv.Atoi(ye)

			// Vertical line
			if xStart == xEnd {
				if yStart > yEnd {
					yStart, yEnd = yEnd, yStart
				}
				for i := yStart; i <= yEnd; i++ {
					s[makeLabel(xStart, i)] = true
				}
			}

			// Horizontal line
			if yStart == yEnd {
				if xStart > xEnd {
					xStart, xEnd = xEnd, xStart
				}
				for i := xStart; i <= xEnd; i++ {
					s[makeLabel(i, yStart)] = true
				}
			}

			if yEnd > maxY {
				maxY = yEnd
			}
		}
		return nil
	})

	return s, maxY
}

func (s scan) dropOne(maxY int, floor bool) {
	x, y := 500, 0

	for {
		if !floor && y > maxY {
			break
		}
		if floor && y == maxY+1 {
			s[makeLabel(x, y)] = true
			break
		}
		if !s[makeLabel(x, y+1)] {
			y++
			continue
		}
		if !s[makeLabel(x-1, y+1)] {
			x, y = x-1, y+1
			continue
		}
		if !s[makeLabel(x+1, y+1)] {
			x, y = x+1, y+1
			continue
		}
		s[makeLabel(x, y)] = true
		break
	}
}

func dropSand(path string, floor bool) int {
	s, maxY := parse(path)
	rocks := len(s)
	sandAndRocks := len(s)
	for {
		s.dropOne(maxY, floor)
		if len(s) == sandAndRocks {
			break
		}
		sandAndRocks = len(s)
	}
	return sandAndRocks - rocks
}

func main() {
	fmt.Println(dropSand("./input", false))
	fmt.Println("---------")
	fmt.Println(dropSand("./input", true))
}
