package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
)

func main() {
	nums, err := utils.InputToSlice("./input.txt")
	if err != nil {
		log.Fatal(err)
	}

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
