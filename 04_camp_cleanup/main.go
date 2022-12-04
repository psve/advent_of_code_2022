package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

type interval [2]int

func newInterval(in string) interval {
	parts := strings.Split(in, "-")
	start, _ := strconv.Atoi(parts[0])
	end, _ := strconv.Atoi(parts[1])
	return interval{start, end}
}

func getIntervals(in string) [2]interval {
	intervals := strings.Split(in, ",")
	return [2]interval{newInterval(intervals[0]), newInterval(intervals[1])}
}

// contains returns true if a is fully contained in b
func contains(a, b interval) bool {
	return a[0] >= b[0] && a[1] <= b[1]
}

// disjoint contains true if a and b don't intersect at all
func disjoint(a, b interval) bool {
	return (a[1] < b[0]) || (b[1] < a[0])
}

func fullyContains(path string) int {
	var count int
	helper.ForEachLine(path, func(line string) error {
		intervals := getIntervals(line)
		if contains(intervals[0], intervals[1]) || contains(intervals[1], intervals[0]) {
			count++
		}
		return nil
	})
	return count
}

func overlaps(path string) int {
	var count int
	helper.ForEachLine(path, func(line string) error {
		intervals := getIntervals(line)
		if !disjoint(intervals[0], intervals[1]) {
			count++
		}
		return nil
	})
	return count
}

func main() {
	fmt.Println(fullyContains("./input"))
	fmt.Println("---------")
	fmt.Println(overlaps("./input"))
}
