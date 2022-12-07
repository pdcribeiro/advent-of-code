package main

import (
	"fmt"

	"example.com/utils"
)

const path = "./test"

func main() {
	partOne()
	partTwo()
}

// *** Part One ***

func partOne() {
	file := utils.OpenFile(path)
	defer file.Close()

	var line string
	result := 0

	for file.ReadLine(&line) {
		fmt.Println(line)
	}

	fmt.Println(result)
}

// *** Part Two ***

func partTwo() {
}

// *** Helpers ***
