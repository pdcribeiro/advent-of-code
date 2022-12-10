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
	file := utils.OpenFile(path)
	defer file.Close()

	grid := loadGrid(file)
	result := maxScore(grid)

	fmt.Println(result)
}

// *** Helpers ***

const xAxis = 0
const yAxis = 1

type grid [][]tree
type tree struct {
	height  int
	visible bool
	views   []int
}

func loadGrid(file *utils.File) *grid {
	var line string
	var trees grid

	for i := 0; file.ReadLine(&line); i++ {
		trees = append(trees, make([]tree, len(line)))
		for j, char := range line {
			height := utils.GetNumber(string(char))
			views := []int{-1, -1, -1, -1}
			trees[i][j] = tree{height, false, views}
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

func maxScore(grid *grid) int {
	grid.setViews(xAxis)
	grid.setViews(yAxis)
	return grid.maxViewScore()
}

func (grid *grid) setViews(axis int) {
	iterate := iterator(len(*grid))
	iterate(false, func(i int) {
		reverse := false
		iterate(reverse, func(j int) {
			for k := j - 1; k >= 0; k-- {
				updateView(grid, i, j, k, axis, reverse)
			}
		})
	})
	iterate(false, func(i int) {
		reverse := true
		iterate(reverse, func(j int) {
			for k := j + 1; k < len(*grid); k++ {
				updateView(grid, i, j, k, axis, reverse)
			}
		})
	})
}

func updateView(grid *grid, i, j, k, axis int, reverse bool) {
	tree := grid.tree(i, k, axis)
	blocker := grid.tree(i, j, axis)
	view := tree.view(axis, reverse)
	if wasUpdated(view) {
		return
	}
	if blocker.hides(tree) || blockerAtEdge(grid, j, reverse) {
		*view = utils.Abs(j - k)
	}
}

func (tree *tree) view(axis int, reverse bool) *int {
	d := direction(axis, reverse)
	return &tree.views[d]
}

func direction(axis int, reverse bool) (d int) {
	d = 0
	if axis == yAxis {
		d += 1
	}
	if reverse {
		d += 2
	}
	return
}

func wasUpdated(view *int) bool {
	return *view > -1
}

func (t *tree) hides(other *tree) bool {
	return t.height >= other.height
}

func blockerAtEdge(grid *grid, index int, reverse bool) bool {
	if reverse {
		return index == 0
	} else {
		return index == len(*grid)-1
	}
}

func (grid *grid) maxViewScore() (max int) {
	iterate := iterator(len(*grid))
	iterate(false, func(i int) {
		iterate(false, func(j int) {
			score := grid.tree(i, j, xAxis).score()
			if score > max {
				max = score
			}
		})
	})
	return
}

func (tree *tree) score() (score int) {
	score = 1
	for _, view := range tree.views {
		if view == -1 {
			view = 0
		}
		score *= view
	}
	return
}
