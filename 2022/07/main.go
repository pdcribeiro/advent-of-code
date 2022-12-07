package main

import (
	"fmt"
	"strings"

	"example.com/utils"
)

const path = "./input"

func main() {
	partOne()
	partTwo()
}

// *** Part One ***

const maxSize = 100000

func partOne() {
	file := utils.OpenFile(path)
	defer file.Close()

	result := 0
	skipLines(file, 2)
	dirSize(file, &result)

	fmt.Println(result)
}

// *** Part Two ***

const totalSize = 70000000
const updateSize = 30000000

func partTwo() {
	file := utils.OpenFile(path)
	defer file.Close()

	var sizes []int
	skipLines(file, 2)
	used := dirSize2(file, &sizes)
	unused := totalSize - used
	minToFree := updateSize - unused
	result := findSmallestAbove(&sizes, minToFree)

	fmt.Println(result)
}

// *** Helpers ***

func dirSize(file *utils.File, sum *int) (size int) {
	var line string
	size = 0

	for file.ReadLine(&line) {
		if isExit(line) {
			addSizeIfSmall(size, sum)
			return
		}
		if isOpen(line) {
			skipLines(file, 1)
			size += dirSize(file, sum)
		} else if isFile(line) {
			size += fileSize(line)
		}
	}
	addSizeIfSmall(size, sum)
	return
}

func skipLines(file *utils.File, count int) {
	var line string
	for i := 0; i < count; i++ {
		file.ReadLine(&line)
	}
}

func isExit(line string) bool {
	return line == "$ cd .."
}

func addSizeIfSmall(size int, sum *int) {
	if size <= maxSize {
		*sum += size
	}
}

func isOpen(line string) bool {
	return line[:5] == "$ cd "
}

func isFile(line string) bool {
	return !isDir(line)
}

func isDir(line string) bool {
	return line[:4] == "dir "
}

func fileSize(line string) int {
	items := strings.Split(line, " ")
	return utils.GetNumber(items[0])
}

func dirSize2(file *utils.File, sizes *[]int) (size int) {
	var line string
	size = 0

	for file.ReadLine(&line) {
		if isExit(line) {
			addSize(size, sizes)
			return
		}
		if isOpen(line) {
			skipLines(file, 1)
			size += dirSize2(file, sizes)
		} else if isFile(line) {
			size += fileSize(line)
		}
	}
	addSize(size, sizes)
	return
}

func addSize(size int, sizes *[]int) {
	*sizes = append(*sizes, size)
}

func findSmallestAbove(sizes *[]int, minToFree int) int {
	for _, s := range *sizes {
		if s >= minToFree {
			return s
		}
	}
	panic("Size not found")
}
