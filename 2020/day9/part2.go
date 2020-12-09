package main

import (
	"fmt"
	"sort"
)

func part2(values []int) {
	check := 530627549

startOver:
	for idx, i := range values {
		sum := i
		list := []int{i}
		for _, j := range values[idx+1:] {
			sum += j
			list = append(list, j)

			if sum == check {
				sort.Ints(list)
				fmt.Println(list[0] + list[len(list)-1])
				break startOver
			}

			if j == check || sum > check {
				continue startOver
			}
		}
	}
}
