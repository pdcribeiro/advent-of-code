package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

const path = "./input"

func main() {
	file := open(path)
	defer file.close()

	result := 0

	for file.read() {
		line := file.GetText()
		result += getScore(line)
	}

	fmt.Println(result)
}

func getScore(line string) int {
	first := int(line[0] - 'A')
	outcome := int(line[2] - 'X')
	outcomeScore := outcome * 3
	return getShapeScore(first, outcome) + outcomeScore
}

func getShapeScore(first int, outcome int) int {
	var second int
	if outcome == 0 {
		second = first - 1
	} else if outcome == 1 {
		second = first
	} else {
		second = first + 1
	}
	return boundValue(second) + 1
}

func boundValue(value int) int {
	if value == -1 {
		return 2
	}
	if value == 3 {
		return 0
	}
	return value
}

// Files

type File struct {
	file    *os.File
	scanner *bufio.Scanner
}

func open(path string) *File {
	file, err := os.Open(path)
	handleError(err)
	scanner := bufio.NewScanner(file)
	return &File{file, scanner}
}

func (f File) read() bool {
	ok := f.scanner.Scan()
	if !ok {
		handleError(f.scanner.Err())
	}
	return ok
}

func (f *File) GetText() string {
	return f.scanner.Text()
}

func (f *File) close() error {
	return f.file.Close()
}

// Misc

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}
