package main

import (
	"advent-of-code/utils"
	"fmt"
	"strings"
)

type coord struct {
	row, col int
}

type seat struct {
	value string
	flip  bool
}

const floor = "."
const empty = "L"
const occupied = "#"

func main() {
	rows := utils.InputToSlice("./input.txt", "\n")

	// rows := []string{"L.LL.LL.LL",
	// 	"LLLLLLL.LL",
	// 	"L.L.L..L..",
	// 	"LLLL.LL.LL",
	// 	"L.LL.LL.LL",
	// 	"L.LLLLL.LL",
	// 	"..L.L.....",
	// 	"LLLLLLLLLL",
	// 	"L.LLLLLL.L",
	// 	"L.LLLLL.LL"}
	grid := make([][]string, len(rows))
	gridMap := make(map[coord]*seat)
	for row, r := range rows {
		grid[row] = strings.Split(r, "")
		for col, s := range grid[row] {
			c := coord{row, col}
			gridMap[c] = &seat{value: s, flip: false}
		}
	}

	// don't run these together - they mutate the same objects and you'll get an incorrect value for part 2
	// part1(gridMap, grid)
	part2(gridMap, grid)
}

func part1(gridMap map[coord]*seat, grid [][]string) {
getOut:
	for {
		numFlipped := 0
		for c, s := range gridMap {
			if s.value == floor {
				continue
			}

			checks := []coord{
				coord{row: c.row - 1, col: c.col - 1}, // nw
				coord{row: c.row - 1, col: c.col},     // n
				coord{row: c.row - 1, col: c.col + 1}, // ne
				coord{row: c.row, col: c.col + 1},     // e
				coord{row: c.row + 1, col: c.col + 1}, // se
				coord{row: c.row + 1, col: c.col},     // s
				coord{row: c.row + 1, col: c.col - 1}, // sw
				coord{row: c.row, col: c.col - 1}}     // w

			flip := false
			if s.value == empty {
				flip = shouldFlipEmpty(checks, gridMap)
			} else {
				flip = shouldFlipOccupied(checks, gridMap, 4)
			}

			if flip {
				// need the flip flag because we apply rule to all seats at once
				// if we flipped it now, the later seats wouldn't be accurate
				s.flip = flip
				numFlipped++
			}
		}

		updateGridMap(gridMap, grid)

		if numFlipped == 0 {
			totalOccupied := 0
			for _, s := range gridMap {
				if s.value == occupied {
					totalOccupied++
				}
			}
			fmt.Println(totalOccupied)
			break getOut
		}
	}
}

func shouldFlipEmpty(checks []coord, gridMap map[coord]*seat) bool {
	allSeatsEmpty := true
checkLoop:
	for _, c := range checks {
		if check, ok := gridMap[c]; ok {
			if check.value == occupied {
				allSeatsEmpty = false
				break checkLoop
			}
		}
	}

	return allSeatsEmpty
}

func shouldFlipOccupied(checks []coord, gridMap map[coord]*seat, limit int) (shouldFlip bool) {
	numOccupied := 0
	for _, c := range checks {
		if check, ok := gridMap[c]; ok {
			if check.value == occupied {
				numOccupied++
			}
		}
	}

	if numOccupied >= limit {
		shouldFlip = true
	}
	return
}

func updateGridMap(gridMap map[coord]*seat, grid [][]string) {
	for c, s := range gridMap {
		if s.value == floor {
			continue
		}

		if s.flip {
			if s.value == empty {
				s.value = occupied
			} else if s.value == occupied {
				s.value = empty
			}
		}

		s.flip = false

		grid[c.row][c.col] = s.value
	}
}
