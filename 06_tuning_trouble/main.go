package main

import (
	"fmt"
	"io"
	"os"
)

type multiSet map[rune]int

func (m *multiSet) insert(e rune) {
	(*m)[e]++
}

func (m *multiSet) remove(e rune) {
	count, ok := (*m)[e]
	if !ok {
		panic("tried to remove non-existing element")
	}
	if count == 1 {
		delete(*m, e)
		return
	}
	(*m)[e]--
}

func (m *multiSet) size() int {
	return len(*m)
}

func findMarker(data []rune) int {
	m := make(multiSet)
	for i := 0; i < 3; i++ {
		m.insert(data[i])
	}
	for i := 3; i < len(data); i++ {
		m.insert(data[i])
		if m.size() == 4 {
			return i + 1
		}
		m.remove(data[i-3])
	}
	panic("no marker found")
}

func findMessage(data []rune) int {
	m := make(multiSet)
	for i := 0; i < 13; i++ {
		m.insert(data[i])
	}
	for i := 13; i < len(data); i++ {
		m.insert(data[i])
		if m.size() == 14 {
			return i + 1
		}
		m.remove(data[i-13])
	}
	panic("no marker found")
}

func main() {
	file, err := os.Open("./input")
	if err != nil {
		panic(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		panic(err)
	}

	fmt.Println(findMarker([]rune(string(data))))
	fmt.Println("---------")
	fmt.Println(findMessage([]rune(string(data))))
}
