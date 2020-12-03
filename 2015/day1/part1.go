package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
)

func main() {
	chars, err := utils.InputToSlice("./input.txt", "")
	if err != nil {
		log.Fatal(err)
	}

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
