package main

import (
	"fmt"
	"sort"
)

func part2(passIDs []int) {
	sort.Ints(passIDs)
	var missing []int

	valInc := passIDs[0]

	for i, id := range passIDs {
		if id != valInc {
			missing = append(missing, valInc)
			valInc = passIDs[i+1]
			continue
		}
		valInc++
	}
	fmt.Println(missing)
}
