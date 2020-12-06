package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
)

type planeSeat struct {
	row, col int
}

func main() {
	passes, err := utils.InputToSlice("./input.txt", "\n")
	if err != nil {
		log.Fatal(err)
	}

	var passIDs []int

	for _, p := range passes {
		var seat planeSeat
		planeRows := make([]int, 128)
		for i := range planeRows {
			planeRows[i] = i
		}
		planeCols := make([]int, 8)
		for i := range planeCols {
			planeCols[i] = i
		}

		rMoves := p[:7]
		cMoves := p[7:]

		for _, rMove := range rMoves {
			switch rMove {
			case 'F':
				planeRows = planeRows[:len(planeRows)/2]
				break
			case 'B':
				planeRows = planeRows[len(planeRows)/2:]
			}
			seat.row = planeRows[0]
		}
		for _, cMove := range cMoves {
			switch cMove {
			case 'L':
				planeCols = planeCols[:len(planeCols)/2]
				break
			case 'R':
				planeCols = planeCols[len(planeCols)/2:]
			}
			seat.col = planeCols[0]
		}

		passIDs = append(passIDs, (seat.row*8)+seat.col)
	}

	part1(passIDs)
	part2(passIDs)
}

func part1(passIDs []int) {
	maxSeatID := 0

	for _, id := range passIDs {
		if id > maxSeatID {
			maxSeatID = id
		}
	}
	fmt.Println(maxSeatID)
}
