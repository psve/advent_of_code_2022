package main

import (
	"fmt"
	"helper"
	"strconv"
)

type tree struct {
	height  int
	visible bool
}

type grid struct {
	trees   [][]tree
	visible int
}

func (g *grid) setVisible(row, col int) {
	tree := &g.trees[row][col]
	if tree.visible {
		return
	}
	tree.visible = true
	g.visible++
}

func (g *grid) getTree(row, col int) *tree {
	return &g.trees[row][col]
}

func (g *grid) size() int {
	return len(g.trees)
}

func (g *grid) scenicScore(row, col int) int {
	directions := []struct {
		rowMod, colMod int
	}{
		{0, 1}, {0, -1}, {1, 0}, {-1, 0},
	}
	inBounds := func(row, col int) bool {
		return row >= 0 && row < g.size() && col >= 0 && col < g.size()
	}

	tree := g.getTree(row, col)
	score := 1
	for _, d := range directions {
		distance := 0
		for i, j := row+d.rowMod, col+d.colMod; inBounds(i, j); i, j = i+d.rowMod, j+d.colMod {
			distance++
			if g.getTree(i, j).height >= tree.height {
				break
			}
		}
		score *= distance
	}
	return score
}

func readGrid(path string) grid {
	g := grid{trees: make([][]tree, 0)}
	helper.ForEachLine(path, func(line string) error {
		row := make([]tree, 0, len(line))
		for _, r := range line {
			h, _ := strconv.Atoi(string(r))
			row = append(row, tree{height: h})
		}
		g.trees = append(g.trees, row)
		return nil
	})
	return g
}

func countVisible(path string) int {
	g := readGrid(path)
	for i := 0; i < g.size(); i++ {
		tallestRight, tallestLeft, tallestDown, tallestUp := -1, -1, -1, -1
		for d := 0; d < g.size(); d++ {
			// Going left along a row i
			if tree := g.getTree(i, d); tree.height > tallestRight {
				g.setVisible(i, d)
				tallestRight = tree.height
			}
			// Going right along a row i
			if tree := g.getTree(i, g.size()-1-d); tree.height > tallestLeft {
				g.setVisible(i, g.size()-1-d)
				tallestLeft = tree.height
			}
			// Going down along column i
			if tree := g.getTree(d, i); tree.height > tallestDown {
				g.setVisible(d, i)
				tallestDown = tree.height
			}
			// Going up along column i
			if tree := g.getTree(g.size()-1-d, i); tree.height > tallestUp {
				g.setVisible(g.size()-1-d, i)
				tallestUp = tree.height
			}
		}
	}
	return g.visible
}

func bestScenicScore(path string) int {
	g := readGrid(path)
	best := 0
	for i := 1; i < g.size()-1; i++ {
		for j := 1; j < g.size()-1; j++ {
			if score := g.scenicScore(i, j); score > best {
				best = score
			}
		}
	}
	return best
}

func main() {
	fmt.Println(countVisible("./input"))
	fmt.Println("---------")
	fmt.Println(bestScenicScore("./input"))
}
