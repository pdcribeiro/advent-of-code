package main

import (
	"fmt"
	"strconv"
	s "strings"

	"example.com/utils"
)

const path = "./input"

func main() {
	file := utils.OpenFile(path)
	defer file.Close()

	result := 0

	for file.ReadLineOld() {
		line := file.GetText() // '2-4,6-8'
		pairs := getBothPairs(line)
		if checkContains(pairs) {
			result += 1
		}
	}

	fmt.Println(result)
}

type Pair []int

func getBothPairs(line string) []Pair {
	pairs := s.Split(line, ",") // ['2-4','6-8']
	return utils.Map[string, Pair](pairs, getPair)
}

func getPair(pair string) Pair {
	sections := s.Split(pair, "-") // ['2', '4']
	return utils.Map[string, int](sections, getNumber)
}

func getNumber(s string) int {
	n, err := strconv.Atoi(s)
	utils.HandleError(err)
	return n
}

func checkContains(pairs []Pair) bool {
	first := pairs[0]
	second := pairs[1]
	return first.checkContains(second) || second.checkContains(first)
}

func (p Pair) checkContains(other Pair) bool {
	return p[0] <= other[0] && p[1] >= other[1]
}
