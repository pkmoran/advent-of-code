package main

import "fmt"

func part2(chars []string, UP string, DOWN string) {
	var floor = 0
	var pos int
	for i, c := range chars {
		if c == UP {
			floor++
			continue
		}

		if c == DOWN {
			floor--

			if floor < 0 {
				pos = i + 1
				break
			}
			continue
		}
	}

	fmt.Printf("Part 2\nFloor: %d\nPosition: %d", floor, pos)

}
