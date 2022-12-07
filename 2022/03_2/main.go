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

	for {
		lines, eof := readThreeLines(file)
		if eof {
			break
		}
		common := getCommon(lines)
		result += getPriority(common)
	}

	fmt.Println(result)
}

func readThreeLines(file *utils.File) (lines []string, eof bool) {
	for i := 0; i < 3; i++ {
		eof = !file.ReadLineOld()
		if eof {
			return
		}
		lines = append(lines, file.GetText())
	}
	return
}

func getCommon(lines []string) rune {
	s := createSet(lines[0])
	s = s.Intersection(createSet(lines[1]))
	s = s.Intersection(createSet(lines[2]))
	for item := range *s {
		return item
	}
	panic("Common element not found")
}

func createSet(items string) *utils.Set[rune] {
	return utils.NewSet[rune]([]rune(items)...)
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
