package main

import (
	"bufio"
	"container/heap"
	"fmt"
	"os"
	"strconv"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// processElfs runs through the input file, calculates calories totals for each Elf,
// and calls finalize with each total.
func processElfs(finalize func(int)) error {
	file, err := os.Open("./input")
	if err != nil {
		return err
	}

	var total int
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// New Elf, pass total to finalize and prepare for a new one.
		if scanner.Text() == "" {
			finalize(total)
			total = 0
			continue
		}

		calories, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return err
		}
		total += calories
	}
	return scanner.Err()
}

func most() {
	var most int
	err := processElfs(func(total int) {
		most = max(most, total)
	})
	if err != nil {
		panic(err)
	}
	fmt.Println(most)
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

func topThree() {
	h := &maxHeap{}
	heap.Init(h)
	err := processElfs(func(total int) {
		heap.Push(h, total)
	})
	if err != nil {
		panic(err)
	}

	total := heap.Pop(h).(int)
	total += heap.Pop(h).(int)
	total += heap.Pop(h).(int)
	fmt.Println(total)
}

func main() {
	most()
	fmt.Println("---------")
	topThree()
}
