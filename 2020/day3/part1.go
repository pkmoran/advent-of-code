package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

func main() {
	lines := utils.InputToSlice("./input.txt", "\n")

	part1(lines)
	part2(lines)
}

func part1(lines []string) {
	fmt.Println(getTreesSmashed(lines, 3, 1))
}

func getTreesSmashed(lines []string, right int, down int) int {
	total := 0
	width := len(strings.Split(lines[0], ""))

	pos := 0
	for i := down; i < len(lines); i += down {
		spaces := strings.Split(lines[i], "")
		check := pos + right
		if check > width-1 {
			check = check - width
			pos = check
		} else {
			pos += right
		}

		if spaces[check] == "#" {
			total++
		}

	}
	return total
}
