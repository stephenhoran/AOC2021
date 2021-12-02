package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"

	"github.com/stephenhoran/aocutils"
)

type Direction string

type Move struct {
	Direction string
	Units     int
}

func ParseMoves(input []byte) []Move {
	moves := make([]Move, 0)

	stringMoves := strings.Split(string(input), "\n")
	for _, n := range stringMoves {
		move := strings.Split(n, " ")
		unit, _ := strconv.Atoi(move[1])
		moves = append(moves, Move{
			Direction: move[0],
			Units:     unit,
		})
	}

	return moves
}

func Day1(moves []Move) {
	var depth, distance int

	for _, move := range moves {
		switch move.Direction {
		case "forward":
			distance += move.Units
		case "down":
			depth += move.Units
		case "up":
			depth -= move.Units
		}
	}

	fmt.Println(depth * distance)
}

func Day2(moves []Move) {
	var depth, distance, aim int

	for _, move := range moves {
		switch move.Direction {
		case "forward":
			distance += move.Units
			if aim != 0 {
				depth += move.Units * aim
			}
		case "down":
			aim += move.Units
		case "up":
			aim -= move.Units
		}
	}

	fmt.Println(depth * distance)
}

func main() {
	input := aocutils.InputToSliceByte()
	parsedInput := ParseMoves(input)
	day := flag.Int("day", 1, "Which Day to Run")
	flag.Parse()

	switch *day {
	case 1:
		Day1(parsedInput)
	case 2:
		Day2(parsedInput)
	}
}
