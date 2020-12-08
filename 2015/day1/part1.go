package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	chars := utils.InputToSlice("./input.txt", "")

	const UP = "("
	const DOWN = ")"

	part1(chars, UP, DOWN)
	part2(chars, UP, DOWN)
}

func part1(chars []string, UP string, DOWN string) {
	var floor = 0
	for _, c := range chars {
		if c == UP {
			floor++
			continue
		}

		if c == DOWN {
			floor--
			continue
		}
	}

	fmt.Printf("Part 1\nFloor: %d\n\n", floor)
}
