package main

import (
	"fmt"

	"example.com/utils"
)

const path = "./input"

func main() {
	file := utils.OpenFile(path)
	defer file.Close()

	result := 0

	for file.ReadLineOld() {
		line := file.GetText()

		common := getCommon(line)
		result += getPriority(common)
	}

	fmt.Println(result)
}

func getCommon(items string) rune {
	size := len(items) / 2
	s := utils.NewSet[byte]()

	for i := 0; i < size; i++ {
		s.Add(items[i])
	}
	for i := size; i < len(items); i++ {
		if s.Has(items[i]) {
			return rune(items[i])
		}
	}

	panic("Common element not found")
}

func getPriority(item rune) int {
	if item >= 'a' && item <= 'z' {
		return 1 + int(item-'a')
	}
	if item >= 'A' && item <= 'Z' {
		return 27 + int(item-'A')
	}
	panic("Invalid character")
}
