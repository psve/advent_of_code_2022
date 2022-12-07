package main

import (
	"fmt"
	"helper"
	"strconv"
	"strings"
)

type Dir struct {
	name     string
	fileSize int
	sub      map[string]*Dir
	parent   *Dir
}

func parseDir(path string) *Dir {
	var current, root *Dir
	helper.ForEachLine(path, func(line string) error {
		switch line {
		case "$ ls":
			return nil
		case "$ cd /":
			current = &Dir{name: "/", sub: make(map[string]*Dir)}
			root = current
			return nil
		case "$ cd ..":
			current = current.parent
			return nil
		}
		if strings.HasPrefix(line, "$ cd ") {
			name := line[5:]
			current = current.sub[name]
			return nil
		}
		if strings.HasPrefix(line, "dir ") {
			name := line[4:]
			current.sub[name] = &Dir{name: name, sub: make(map[string]*Dir), parent: current}
			return nil
		}

		fileInfo := strings.Split(line, " ")
		size, _ := strconv.Atoi(fileInfo[0])
		current.fileSize += size
		return nil
	})
	return root
}

// dirSize runs through the directory tree, calculating the total size of each directory.
// The hook function is called with the total size of each directory.
func dirSize(dir *Dir, hook func(int)) int {
	totalSize := dir.fileSize
	for _, sub := range dir.sub {
		totalSize += dirSize(sub, hook)
	}
	hook(totalSize)
	return totalSize
}

func getSizes(path string) int {
	root := parseDir(path)
	res := 0
	dirSize(root, func(size int) {
		if size <= 100000 {
			res += size
		}
	})
	return res
}

func deleteDirectory(path string) int {
	root := parseDir(path)
	totalSize := dirSize(root, func(int) {})
	target := 30000000 - (70000000 - totalSize)
	res := totalSize
	dirSize(root, func(size int) {
		if size >= target && size < res {
			res = size
		}
	})
	return res
}

func main() {
	fmt.Println(getSizes("./input"))
	fmt.Println("---------")
	fmt.Println(deleteDirectory("./input"))
}
