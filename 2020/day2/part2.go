package main

import "fmt"

func part2(passwords []passwordRule) {
	var total int
	for _, p := range passwords {
		if string(p.password[p.min-1]) != string(p.password[p.max-1]) {
			if string(p.password[p.min-1]) == p.letter || string(p.password[p.max-1]) == p.letter {
				total++
			}
		}
	}

	fmt.Printf("Part 2:\nTotal correct: %d", total)
}
