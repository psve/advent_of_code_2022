package main

import (
	"fmt"
	"helper"
)

type set map[rune]any

func toPriority(item rune) int {
	if item <= 'a' {
		return int(item-'A') + 27
	}
	return int(item-'a') + 1
}

func priority(contents string) int {
	compartment := make(set)
	l := len(contents) / 2
	for _, item := range contents[:l] {
		compartment[item] = nil
	}
	for _, item := range contents[l:] {
		if _, found := compartment[item]; found {
			return toPriority(item)
		}
	}
	panic("no item found")
}

func totalPriority(path string) int {
	var total int
	helper.ForEachLine(path, func(line string) error {
		total += priority(line)
		return nil
	})
	return total
}

func badgePriority(path string) int {
	var total, i int
	carriedBy := make(map[rune]int)
	helper.ForEachLine(path, func(contents string) error {
		for _, item := range contents {
			// Set the bit of the first/second/third Elf in the group.
			// If we've set all three (which can only happen on the third Elf),
			// we're done, and can reset.
			if carriedBy[item] |= 1 << (i % 3); carriedBy[item] == 0b111 {
				total += toPriority(item)
				carriedBy = make(map[rune]int)
				break
			}
		}
		i++
		return nil
	})
	return total
}

func main() {
	fmt.Println(totalPriority("./input"))
	fmt.Println("---------")
	fmt.Println(badgePriority("./input"))
}
