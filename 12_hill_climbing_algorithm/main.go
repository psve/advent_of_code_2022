package main

import (
	"fmt"
	"helper"
	"math"
)

type vertex struct {
	label      string
	height     int
	neighbours []*vertex
}

func newVertex(label string, value int) vertex {
	return vertex{
		label:      label,
		height:     value,
		neighbours: make([]*vertex, 0),
	}
}

type graph struct {
	vertices map[string]*vertex
}

func newGraph() graph {
	return graph{make(map[string]*vertex)}
}

// Shortest path from source to all other vertices via Dijkstra's algorithm.
func (g *graph) shortestPath(source string) map[string]int {
	dist := make(map[string]int)
	queue := &PriorityQueue{}
	queue.Push(&Item{
		value:    g.vertices[source],
		priority: 0,
	})

	for queue.Len() != 0 {
		u := queue.Pop().(*Item)
		for _, v := range u.value.neighbours {
			x, ok := dist[v.label]
			if alt := dist[u.value.label] + 1; !ok || alt < x {
				dist[v.label] = alt
				queue.Push(&Item{
					value:    v,
					priority: alt,
				})
			}
		}
	}

	return dist
}

func makeLabel(row, col int) string {
	return fmt.Sprintf("(%d,%d)", row, col)
}

func parse(path string) (graph, string, string) {
	g := newGraph()
	var start, end string
	var rows, cols int
	helper.ForEachLine(path, func(line string) error {
		for i, h := range line {
			switch h {
			case 'S':
				v := newVertex(makeLabel(rows, i), 0)
				g.vertices[v.label] = &v
				start = v.label
			case 'E':
				v := newVertex(makeLabel(rows, i), 25)
				g.vertices[v.label] = &v
				end = v.label
			default:
				v := newVertex(makeLabel(rows, i), int(h)-int('a'))
				g.vertices[v.label] = &v
			}
			cols = i + 1
		}
		rows++
		return nil
	})

	for i := 0; i < rows; i++ {
		for j := 0; j < cols; j++ {
			v := g.vertices[makeLabel(i, j)]

			// Calculate each vertex's neighbours. Note that we are orienting the edges *away* from the
			// end vertex to make part 2 easier, i.e. we find the shortest path from the end to all other
			// vertices.
			if u, ok := g.vertices[makeLabel(i, j+1)]; ok && v.height+1 >= u.height {
				u.neighbours = append(u.neighbours, v)
			}
			if u, ok := g.vertices[makeLabel(i, j-1)]; ok && v.height+1 >= u.height {
				u.neighbours = append(u.neighbours, v)
			}
			if u, ok := g.vertices[makeLabel(i-1, j)]; ok && v.height+1 >= u.height {
				u.neighbours = append(u.neighbours, v)
			}
			if u, ok := g.vertices[makeLabel(i+1, j)]; ok && v.height+1 >= u.height {
				u.neighbours = append(u.neighbours, v)
			}
		}
	}

	return g, start, end
}

func climbHill(path string) int {
	g, start, end := parse(path)
	dist := g.shortestPath(end)
	return dist[start]
}

func hikeTrail(path string) int {
	g, _, end := parse(path)
	dist := g.shortestPath(end)

	// Find the shortest path from the end vertex to a vertex with height 0.
	min := math.MaxInt
	for l, v := range g.vertices {
		if d, ok := dist[l]; ok && v.height == 0 && d < min {
			min = d
		}
	}
	return min
}

func main() {
	fmt.Println(climbHill("./input"))
	fmt.Println("---------")
	fmt.Println(hikeTrail("./input"))
}
