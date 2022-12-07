package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

const path = "./input"

func main() {
	file, err := os.Open(path)
	handleError(err)
	defer file.Close()

	scanner := bufio.NewScanner(file)
	sum := 0
	maxs := make([]int, 3)

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			updateMaxs(sum, maxs)
			sum = 0
		} else {
			count, err := strconv.Atoi(line)
			handleError(err)
			sum += count
		}
	}
	handleError(scanner.Err())

	updateMaxs(sum, maxs)
	maxsSum := sumSlice(maxs)
	fmt.Println(maxsSum)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func updateMaxs(sum int, maxs []int) {
	fmt.Println("Before update", maxs)
	for i, m := range maxs {
		if sum > m {
			shiftMaxs(maxs, i)
			maxs[i] = sum
			return
		}
	}
	fmt.Println("After update ", maxs, "\n")
}

func shiftMaxs(maxs []int, pos int) {
	for i := len(maxs) - 1; i > pos; i-- {
		maxs[i] = maxs[i-1]
	}
}

func sumSlice(slice []int) int {
	sum := 0
	for _, n := range slice {
		sum += n
	}
	return sum
}
