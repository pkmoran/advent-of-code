package main

import "fmt"

func printGrid(grid [][]string) {
	for _, g := range grid {
		fmt.Println(g)
	}
	fmt.Println("\n\n")
}

func part2(gridMap map[coord]*seat, grid [][]string) {
getOut:
	for {
		numFlipped := 0
		for c, s := range gridMap {
			directionStarts := map[string]struct {
				c              coord
				rowInc, colInc int
			}{"nw": {coord{row: c.row - 1, col: c.col - 1}, -1, -1},
				"n":  {coord{row: c.row - 1, col: c.col}, -1, 0},
				"ne": {coord{row: c.row - 1, col: c.col + 1}, -1, 1},
				"e":  {coord{row: c.row, col: c.col + 1}, 0, 1},
				"se": {coord{row: c.row + 1, col: c.col + 1}, 1, 1},
				"s":  {coord{row: c.row + 1, col: c.col}, 1, 0},
				"sw": {coord{row: c.row + 1, col: c.col - 1}, 1, -1},
				"w":  {coord{row: c.row, col: c.col - 1}, 0, -1}}

			var checks []coord
			for _, s := range directionStarts {
				if s, ok := getVisibleSeat(s.c, s.rowInc, s.colInc, gridMap); ok {
					checks = append(checks, s)
				}
			}

			flip := false
			if s.value == empty {
				flip = shouldFlipEmpty(checks, gridMap)
			} else if s.value == occupied {
				flip = shouldFlipOccupied(checks, gridMap, 5)
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

func getVisibleSeat(start coord, rowInc, colInc int, gridMap map[coord]*seat) (coord, bool) {
	check := start
out:
	for {
		if c, ok := gridMap[check]; ok {
			if c.value == empty || c.value == occupied {
				break out
			}
			check.row += rowInc
			check.col += colInc
			continue out
		}
		return check, false
	}
	return check, true
}
