package main

import (
	"fmt"

	"example.com/utils"
)

const path = "./input"

func main() {
	partOne()
	partTwo()
}

// *** Part One ***

func partOne() {
	file := utils.OpenFile(path)
	defer file.Close()

	grid := loadGrid(file)
	result := visibleCount(grid)

	fmt.Println(result)
}

// *** Part Two ***

func partTwo() {
}

// *** Helpers ***

const xAxis = 0
const yAxis = 1

type grid [][]tree
type tree struct {
	height  int
	visible bool
}

func loadGrid(file *utils.File) *grid {
	var line string
	var trees grid

	for i := 0; file.ReadLine(&line); i++ {
		trees = append(trees, make([]tree, len(line)))
		for j, char := range line {
			height := utils.GetNumber(string(char))
			trees[i][j] = tree{height, false}
		}
	}
	return &trees
}

func visibleCount(grid *grid) (count int) {
	count = 0
	increment := func(isNewVisible bool) {
		if isNewVisible {
			count++
		}
	}
	grid.checkVisible(xAxis, increment)
	grid.checkVisible(yAxis, increment)
	return
}

func (grid *grid) checkVisible(axis int, increment func(isNewVisible bool)) {
	iterateAndCheck := func(reverse bool) {
		iterate := iterator(len(*grid))
		iterate(false, func(i int) {
			isNewVisible := visibilityChecker()
			iterate(reverse, func(j int) {
				tree := grid.tree(i, j, axis)
				increment(isNewVisible(tree))
			})
		})
	}
	iterateAndCheck(false)
	iterateAndCheck(true)
}

func iterator(size int) func(reverse bool, callback func(int)) {
	return func(reverse bool, callback func(int)) {
		if reverse {
			for i := size - 1; i >= 0; i-- {
				callback(i)
			}
		} else {
			for i := 0; i < size; i++ {
				callback(i)
			}
		}
	}
}

func visibilityChecker() func(*tree) bool {
	maxHeight := -1
	return func(tree *tree) bool {
		if tree.height > maxHeight {
			maxHeight = tree.height
			return updateNewVisible(tree)
		} else {
			return false
		}
	}
}

func updateNewVisible(tree *tree) bool {
	if !tree.visible {
		tree.visible = true
		return true
	}
	return false
}

func (grid *grid) tree(i, j, axis int) *tree {
	row, col := rowAndCol(i, j, axis)
	return &(*grid)[row][col]
}

func rowAndCol(i, j, axis int) (int, int) {
	if axis == xAxis {
		return i, j
	} else {
		return j, i
	}
}

func (grid *grid) print() {
	for _, row := range *grid {
		for _, t := range row {
			fmt.Print(t)
		}
		fmt.Println()
	}
	fmt.Println()
}
