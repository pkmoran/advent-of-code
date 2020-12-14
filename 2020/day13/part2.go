package main

import (
	"fmt"
	"strconv"
)

func part2(busIds []string) {
	var i int64 = 100000000000000
	firstBusId := toInt(busIds[0])
	start := i - (i % firstBusId)

	max := 0
	inc := firstBusId

top:
	for {
		for bi := 0; bi < len(busIds); bi++ {
			if busIds[bi] == "x" {
				continue
			}

			id := toInt(busIds[bi])
			if (start+int64(bi))%id != 0 {
				start += inc
				continue top
			}

			if bi > max {
				max = bi
				inc *= id
			}
		}
		break top
	}
	fmt.Println("I:", start)
}

func toInt(s string) int64 {
	i, _ := strconv.Atoi(s)
	return int64(i)
}
