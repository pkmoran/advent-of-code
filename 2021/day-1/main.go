package main

import (
	"advent-of-code/utils"
	"fmt"
)

func main() {
	lines := utils.InputToSlice("./input.txt", "\n")

	nums := utils.StringsToInts(lines)

	var count int
	//for i := 1; i < len(nums); i++ {
	//		if nums[i]-nums[i-1] > 0 {
	//			count++
	//		}
	//}
	//fmt.Println(count)

	for i := 2; i < len(nums); i++ {
		if i+1 == len(nums) {
			break
		}
		nextSum := nums[i+1]+nums[i]+nums[i-1]
		currSum := nums[i]+nums[i-1]+nums[i-2]
			if nextSum-currSum > 0 {
				count++
			}
	}
	fmt.Println(count)
}