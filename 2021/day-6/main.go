package main

import (
	"advent-of-code/utils"
	"fmt"
)

const days = 80
const part2days = 256

func main() {
	//ages := utils.StringsToInts(utils.InputToSlice("test.txt", ","))
	ages := utils.StringsToInts(utils.InputToSlice("input.txt", ","))
	//part1(ages)
	part2(ages)
}

func part1(ages []int) {
	for i := 0; i < days; i++ {
		var nextDay []int

		numNewFish := 0
		for _, age := range ages {
			if age == 0 {
				numNewFish++
				nextDay = append(nextDay, 6)
				continue
			}

			nextDay = append(nextDay, age-1)
		}
		newFish := make([]int, numNewFish)
		for j := range newFish {
			newFish[j] = 8
		}
		nextDay = append(nextDay, newFish...)
		ages = nextDay
	}
	fmt.Println(len(ages))
}

func part2(ages []int) {
	fisheys := make(map[int]int)
	for _, age := range ages {
		fisheys[age-1]++
	}

	currentDay := 1
	for currentDay < part2days {
		tmp := make(map[int]int)
		for age, numFish := range fisheys {
			_, ok := tmp[age-1]
			if ok {
				// new age is a 6
				tmp[age-1] += numFish
			} else {
				if age == 0 {
					tmp[8] = numFish
					tmp[6] += numFish
				} else {
					tmp[age-1] = numFish
				}

			}
		}

		fisheys = tmp
		currentDay++
	}

	sum := 0
	for _, v := range fisheys {
		sum += v
	}
	fmt.Println(sum)
}
