package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strings"
)

func main() {
	groups, err := utils.InputToSlice("./input.txt", "\n\n")
	if err != nil {
		log.Fatal(err)
	}

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
