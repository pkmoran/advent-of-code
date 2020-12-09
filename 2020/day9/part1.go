package main

import (
	"advent-of-code/utils"
	"fmt"
)

const preambleLength = 25

func main() {
	lines := utils.InputToSlice("./input.txt", "\n")
	values := utils.StringsToInts(lines)

	part1(values)
	part2(values)
}

func part1(values []int) {
	preamble := values[:preambleLength]

	preambleInc := 0
	i := preambleLength
startOver:
	for {
		for _, p1 := range preamble {
			for _, p2 := range preamble[1:] {
				if p1 != p2 && p1+p2 == values[i] {
					// valid
					preambleInc++
					preamble = values[preambleInc : preambleInc+preambleLength]
					i++
					continue startOver
				}
			}
		}

		fmt.Println(values[i])
		break
	}
}
