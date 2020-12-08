package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type box struct {
	length int
	height int
	width  int
}

func main() {
	lines := utils.InputToSlice("./input.txt", "\n")

	part1(lines)
	part2(lines)
}

func part1(dimensions []string) {
	total := 0

	for _, d := range dimensions {
		b := strings.Split(d, "x")

		var gift box
		gift.length, _ = strconv.Atoi(b[0])
		gift.width, _ = strconv.Atoi(b[1])
		gift.height, _ = strconv.Atoi(b[2])

		s1, s2, s3 := gift.length*gift.width, gift.length*gift.height, gift.width*gift.height

		// surface area + smallest side
		a := (2 * (s1 + s2 + s3)) + min(s1, s2, s3)

		total += a
	}

	fmt.Printf("Total Part 1: %d\n\n", total)
}

func min(nums ...int) int {
	min := nums[0]

	for _, i := range nums {
		if min > i {
			min = i
		}
	}

	return min
}
