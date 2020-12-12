package main

import (
	"advent-of-code/utils"
	"fmt"
	"sort"
)

func main() {
	adapters := utils.StringsToInts(utils.InputToSlice("./input.txt", "\n"))

	part1(adapters)
	part2(adapters)
}

func part1(adapters []int) {
	sort.Ints(adapters)
	// init 3 to 1 because of the internal device adapter
	diffMap := map[int]int{1: 0, 3: 1}

	joltage := 0
	for _, a := range adapters {
		diff := a - joltage
		joltage = a
		diffMap[diff]++
	}

	fmt.Println(diffMap[1] * diffMap[3])
}
