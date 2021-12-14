package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
)

func main() {
	//positions := utils.StringsToFloat64(utils.InputToSlice("test.txt", ","))
	positions := utils.StringsToFloat64(utils.InputToSlice("input.txt", ","))

	part1(positions)
	part2(positions)
}

func part1(positions []float64) {
	minFuel := math.MaxFloat64
	for _, finalPosOpt := range positions {
		var tmp float64
		for _, pos := range positions {
			tmp += math.Abs(finalPosOpt - pos)
		}
		minFuel = math.Min(minFuel, tmp)
	}
	fmt.Println(minFuel)
}

func part2(positions []float64) {
	minFuel := math.MaxFloat64

	maxPosition := getMaxPosition(positions)
	var posOpt float64
	for posOpt <= maxPosition {
		var tmp float64
		for _, pos := range positions {
			steps := math.Abs(posOpt - pos)
			i := float64(1)
			for i <= steps {
				tmp += i
				i++
			}
		}
		minFuel = math.Min(minFuel, tmp)
		posOpt++
	}
	fmt.Println(int32(minFuel))
}

func getMaxPosition(positions []float64) float64 {
	var max float64
	for _, pos := range positions {
		max = math.Max(max, pos)
	}
	return max
}
