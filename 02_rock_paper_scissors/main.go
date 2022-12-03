package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// round represents the two moves. 0 for rock, 1 for paper, 2 for scissors.
type round [2]int

// Calculate score by observing that you win if your move is one larger than
// the opponents and lose if its two larger, modulo 3.
func score(r round) int {
	switch r[1] {
	case (r[0] + 1) % 3: // Win
		return 6 + r[1] + 1
	case r[0]: // Draw
		return 3 + r[1] + 1
	case (r[0] + 2) % 3: // Lose
		return 0 + r[1] + 1
	default:
		panic(fmt.Sprintf("unknown score %v", r))
	}
}

func fromMoves(moves string) round {
	opponent := map[string]int{"A": 0, "B": 1, "C": 2}
	you := map[string]int{"X": 0, "Y": 1, "Z": 2}
	m := strings.Split(moves, " ")
	return round{opponent[m[0]], you[m[1]]}
}

// fromStrategy converst a strategy to a round by using the same observation
// as in `score`.
func fromStrategy(strategy string) round {
	opponent := map[string]int{"A": 0, "B": 1, "C": 2}
	you := map[string]int{"X": 2, "Y": 0, "Z": 1}
	s := strings.Split(strategy, " ")
	return round{opponent[s[0]], (opponent[s[0]] + you[s[1]]) % 3}
}

func totalScore(toRound func(string) round) {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}

	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		total += score(toRound(scanner.Text()))
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}

	fmt.Println(total)
}

func main() {
	totalScore(fromMoves)
	fmt.Println("---------")
	totalScore(fromStrategy)
}
