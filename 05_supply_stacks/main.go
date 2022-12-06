package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

type stack []rune
type move struct {
	amount, from, to int
}

func parseInput(path string) ([]stack, []move) {
	stackDesc := make([]string, 0)
	moves := make([]move, 0)

	readingStacks := true
	helper.ForEachLine(path, func(line string) error {
		if line == "" {
			readingStacks = false
			return nil
		}

		if readingStacks {
			stackDesc = append(stackDesc, line)
			return nil
		}

		parts := strings.Split(line, " ")
		amount, _ := strconv.Atoi(parts[1])
		from, _ := strconv.Atoi(parts[3])
		to, _ := strconv.Atoi(parts[5])

		moves = append(moves, move{amount, from - 1, to - 1})
		return nil
	})

	parts := strings.Split(stackDesc[len(stackDesc)-1], " ")
	numStacks, _ := strconv.Atoi(parts[len(parts)-1])
	stacks := make([]stack, numStacks)

	for i := len(stackDesc) - 2; i >= 0; i-- {
		for j := 0; j < numStacks; j++ {
			if 1+4*j >= len(stackDesc[i]) {
				continue
			}
			crate := []rune(stackDesc[i])[1+4*j]
			if crate != rune(' ') {
				stacks[j] = append(stacks[j], crate)
			}
		}
	}

	return stacks, moves
}

func crateMover9000(path string) string {
	stacks, moves := parseInput(path)
	for _, move := range moves {
		for i := 0; i < move.amount; i++ {
			hight := len(stacks[move.from])
			crate := stacks[move.from][hight-1]
			stacks[move.from] = stacks[move.from][:hight-1]
			stacks[move.to] = append(stacks[move.to], crate)
		}
	}

	out := make([]rune, 0, len(stacks))
	for _, stack := range stacks {
		out = append(out, stack[len(stack)-1])
	}
	return string(out)
}

func crateMover9001(path string) string {
	stacks, moves := parseInput(path)
	for _, move := range moves {
		hight := len(stacks[move.from])
		crates := stacks[move.from][hight-move.amount : hight]
		stacks[move.from] = stacks[move.from][:hight-move.amount]
		stacks[move.to] = append(stacks[move.to], crates...)
	}

	out := make([]rune, 0, len(stacks))
	for _, stack := range stacks {
		out = append(out, stack[len(stack)-1])
	}
	return string(out)
}

func main() {
	fmt.Println(crateMover9000("./input"))
	fmt.Println("---------")
	fmt.Println(crateMover9001("./input"))
}
