package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type node struct {
	name    string
	value   int
	visited bool
}

type insMap map[int]*node

func main() {
	lines := utils.InputToSlice("./input.txt", "\n")

	instructionMap := make(insMap)
	for i, l := range lines {
		fields := strings.Fields(l)
		name := fields[0]
		value, _ := strconv.Atoi(fields[1])

		instructionMap[i] = &node{name: name, value: value, visited: false}
	}

	part1(instructionMap)
	part2(instructionMap)
}

func part1(instructions insMap) {
	acc := 0

	i := 0
	for {
		if instructions[i].visited {
			break
		}

		instructions[i].visited = true

		switch instructions[i].name {
		case "acc":
			acc += instructions[i].value
			i = i + 1
		case "jmp":
			i = i + instructions[i].value
		case "nop":
			i = i + 1
		}
	}

	fmt.Println(acc)
}
