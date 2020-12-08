package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	groups := utils.InputToSlice("./input.txt", "\n\n")

	total := 0
	total2 := 0

	for _, g := range groups {
		total += part1(g)
		total2 += part2(g)
	}

	fmt.Println(total)
	fmt.Println(total2)
}

func part1(group string) (groupTotal int) {
	checkMap := make(map[string]bool)

	for _, r := range strings.ReplaceAll(group, "\n", "") {
		if _, seen := checkMap[string(r)]; seen {
			continue
		}
		checkMap[string(r)] = true
		groupTotal++
	}
	return
}
