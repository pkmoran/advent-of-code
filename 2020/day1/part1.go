package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines := utils.InputToSlice("./input.txt", "\n")

	nums := utils.StringsToInts(lines)

	part1(nums)
	part2(nums)
}

func part1(nums []int) {
done:
	for _, i := range nums {
		for _, j := range nums {
			if i+j == 2020 {
				fmt.Printf("%d * %d = %d\n", i, j, i*j)
				break done
			}
		}
	}
}
