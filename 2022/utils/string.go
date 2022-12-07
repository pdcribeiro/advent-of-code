package utils

import "strconv"

func GetNumber(s string) int {
	n, err := strconv.Atoi(s)
	HandleError(err)
	return n
}
