package main

import (
	"fmt"
)

func part2(nums []int) {
done:
	for _, i := range nums {
		for _, j := range nums {
			for _, k := range nums {
				if i+j+k == 2020 {
					fmt.Printf("%d * %d * %d = %d\n", i, j, k, i*j*k)
					break done
				}

			}
		}
	}
}
