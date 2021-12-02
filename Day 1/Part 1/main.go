package main

import (
	"fmt"
	"io/ioutil"

	"AOC2021/common"
)

const inputFile string = "../input.txt"

func input() []byte {
	file, err := ioutil.ReadFile(inputFile)
	if err != nil {
		panic(err)
	}

	return file
}

func main() {
	input := input()
	fmt.Println(common.ToSliceInt(input))
}
