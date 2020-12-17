package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	read := utils.InputToSlice("./input.txt", ",")
	input := utils.StringsToInts(read)

	part1(input)
	part2(input)
}

func part1(input []int) {
	const target = 2020

	getAnswer(input, target)
}

func getAnswer(input []int, target int) {
	spoken := make(map[int][]int)

	// initialize map
	for i, v := range input {
		spoken[v] = append(spoken[v], i)
	}

	for i := len(spoken); i <= target; i++ {
		if lastIndecies, ok := spoken[input[i-1]]; ok && len(lastIndecies) > 1 {
			diff := i - 1 - lastIndecies[len(lastIndecies)-2]
			input = append(input, diff)
			if _, ok := spoken[diff]; ok {
				spoken[diff] = append(spoken[diff][len(spoken[diff])-1:], i)
				continue
			}
			spoken[diff] = []int{i}
			continue
		}

		input = append(input, 0)

		if _, ok := spoken[0]; ok {
			spoken[0] = append(spoken[0][len(spoken[0])-1:], i)
			continue
		}
		spoken[0] = []int{i}
	}

	for k, v := range spoken {
		if v[len(v)-1] == target-1 {
			fmt.Println(k)
		}
	}
}
