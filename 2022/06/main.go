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
// Find first set of 4 different characters
// Return index of last character (with 1-based indexing)

func partOne() {
	result := findMarker(4)

	fmt.Println(result)
}

// *** Part Two ***
// Find first set of 14 different characters

func partTwo() {
	result := findMarker(14)

	fmt.Println(result)
}

// *** Helpers ***

type Buffer []string

func findMarker(size int) int {
	file := utils.OpenFile(path)
	defer file.Close()

	return findMarkerInFile(file, size)
}

func findMarkerInFile(file *utils.File, size int) int {
	buffer := initBuffer(file, size)
	for i := len(*buffer) + 1; advanceBuffer(file, buffer); i++ {
		if checkDifferent(buffer) {
			return i
		}
	}
	panic("Marker not found")
}

func initBuffer(file *utils.File, size int) *Buffer {
	var char string
	var buffer Buffer
	for i := 0; i < size; i++ {
		file.ReadChar(&char)
		buffer = append(buffer, char)
	}
	return &buffer
}

func advanceBuffer(file *utils.File, buffer *Buffer) (ok bool) {
	var char string
	ok = file.ReadChar(&char)
	*buffer = (*buffer)[1:]
	*buffer = append(*buffer, char)
	return
}

func checkDifferent(buffer *Buffer) bool {
	s := utils.NewSet[string](*buffer...)
	return s.Size() == len(*buffer)
}
