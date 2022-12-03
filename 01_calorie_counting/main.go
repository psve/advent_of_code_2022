package main

import (
	"container/heap"
	"fmt"
	"strconv"

	"helper"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// processElfs runs through the input file, calculates calories totals for each Elf,
// and calls finalize with each total.
func processElfs(path string, finalize func(int)) error {
	var total int
	err := helper.ForEachLine(path, func(line string) error {
		// New Elf, pass total to finalize and prepare for a new one.
		if line == "" {
			finalize(total)
			total = 0
			return nil
		}

		calories, err := strconv.Atoi(line)
		if err != nil {
			return err
		}
		total += calories
		return nil
	})
	finalize(total)

	return err
}

func most(path string) int {
	var most int
	err := processElfs(path, func(total int) {
		most = max(most, total)
	})
	if err != nil {
		panic(err)
	}
	return most
}

type maxHeap []int

func (h maxHeap) Len() int           { return len(h) }
func (h maxHeap) Less(i, j int) bool { return h[i] > h[j] }
func (h maxHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *maxHeap) Push(x any)        { (*h) = append((*h), x.(int)) }
func (h *maxHeap) Pop() any {
	n := len(*h)
	x := (*h)[n-1]
	*h = (*h)[:n-1]
	return x
}

func topThree(path string) int {
	h := &maxHeap{}
	heap.Init(h)
	err := processElfs(path, func(total int) {
		heap.Push(h, total)
	})
	if err != nil {
		panic(err)
	}

	total := heap.Pop(h).(int)
	total += heap.Pop(h).(int)
	total += heap.Pop(h).(int)
	return total
}

func main() {
	fmt.Println(most("./input"))
	fmt.Println("---------")
	fmt.Println(topThree("./input"))
}
