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
	second := int(line[2] - 'X')
	return getOutcomeScore(first, second) + (second + 1)
}

func getOutcomeScore(first int, second int) int {
	if second == first+1 || (second == 0 && first == 2) {
		return 6
	} else if second == first {
		return 3
	} else {
		return 0
	}
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
