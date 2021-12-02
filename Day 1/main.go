package main

import (
	"flag"
	"fmt"

	"github.com/stephenhoran/aocutils"
)

func Day1(numbers []int) {
	var count, previousValue int

	for index, num := range numbers {
		if index == 0 {
			previousValue = num
			continue
		}

		if num > previousValue {
			count++
		}

		previousValue = num
	}

	fmt.Println(count)
}

func Day2(numbers []int) {
	var count, previousValue int
	for index, _ := range numbers {
		if index == 0 {
			previousValue = slidingWindowSum(numbers, index)
			continue
		}

		if index == len(numbers)-2 {
			break
		}

		n := slidingWindowSum(numbers, index)

		if previousValue < n {
			count++
		}

		previousValue = n
	}

	fmt.Println(count)
}

func slidingWindowSum(numbers []int, index int) int {
	return numbers[index] + numbers[index+1] + numbers[index+2]
}

func main() {
	input := aocutils.InputToSliceInt()
	day := flag.Int("day", 1, "Which Day to Run")
	flag.Parse()

	switch *day {
	case 1:
		Day1(input)
	case 2:
		Day2(input)
	}
}
