package main

import (
	"advent-of-code/utils"
	"fmt"
)

func filter(slice []string, fn func(string) bool) (result []string) {
	for _, s := range slice {
		if fn(s) {
			result = append(result, s)
		}
	}
	return
}

func main() {
	input := utils.InputToSlice("./input.txt", ",")

	part1(utils.StringsToInts(filter(input, func(s string) bool { return s != "x" })))
	part2(input)

}

func part1(busIds []int) {
	// const departureTime = 1005595
	const departureTime = 939
	checks := make(map[int]int)
	for _, busId := range busIds {
		checks[departureTime%busId] = busId
	}

	max := 0
	for k, _ := range checks {
		if k > max {
			max = k
		}
	}

	waitTime := checks[max] - max

	fmt.Println(checks[max] * waitTime)
}
