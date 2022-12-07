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
	max := 0

	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			sum, max = updateMax(sum, max)
		} else {
			count, err := strconv.Atoi(line)
			handleError(err)
			sum += count
		}
	}
	handleError(scanner.Err())

	_, max = updateMax(sum, max)
	fmt.Println(max)
}

func handleError(e error) {
	if e != nil {
		log.Fatal(e)
	}
}

func updateMax(sum int, max int) (int, int) {
	if sum > max {
		return 0, sum
	}
	return 0, max
}
