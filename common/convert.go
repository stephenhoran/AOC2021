package common

import (
	"strconv"
	"strings"
)

func ToSliceInt(bytes []byte) []int {
	ints := make([]int, 0)

	stringInts := strings.Split(string(bytes), "\n")
	for _, number := range stringInts {
		i, _ := strconv.Atoi(number)
		ints = append(ints, i)
	}

	return ints
}