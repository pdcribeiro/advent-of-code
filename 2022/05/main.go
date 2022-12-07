package main

import (
	"fmt"
	"strings"

	"example.com/utils"
)

const path = "./input"

func main() {
	file := utils.OpenFile(path)
	defer file.Close()

	stacks := getStacks(file)
	moveCrates(file, stacks)
	result := getTopCrates(stacks)

	fmt.Println(result)
}

type Stack = utils.Stack[string]

func getStacks(file *utils.File) []Stack {
	lines := getStackLines(file)
	return getStacksFromLines(lines)
}

func getStackLines(file *utils.File) (lines []string) {
	for file.ReadLineOld() {
		line := file.GetText()
		if isNumbersLine(line) {
			break
		}
		lines = append(lines, line)
	}
	skipEmptyLine(file)
	return
}

func isNumbersLine(line string) bool {
	return line[:2] == " 1"
}

func skipEmptyLine(file *utils.File) {
	file.ReadLineOld()
}

func getStacksFromLines(lines []string) (stacks []Stack) {
	count := (len(lines[0]) + 1) / 4
	stacks = make([]Stack, count)

	for i := len(lines) - 1; i >= 0; i-- {
		for j := 0; j < count; j++ {
			crate := lines[i][4*j+1]
			if crate != ' ' {
				stacks[j].Push(string(crate))
			}
		}
	}

	return
}

func moveCrates(file *utils.File, stacks []Stack) {
	for file.ReadLineOld() {
		line := file.GetText()
		count, from, to := parseInstruction(line)
		for i := 0; i < count; i++ {
			crate := stacks[from].Pop()
			stacks[to].Push(crate)
		}
	}
}

func parseInstruction(instruction string) (count, from, to int) {
	values := strings.Split(instruction, " ")
	count = utils.GetNumber(values[1])
	from = utils.GetNumber(values[3]) - 1
	to = utils.GetNumber(values[5]) - 1
	return
}

func getTopCrates(stacks []Stack) string {
	var sb strings.Builder
	for _, stack := range stacks {
		sb.WriteString(stack.Pop())
	}
	return sb.String()
}
