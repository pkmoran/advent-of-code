package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
	"unicode"
)

type point struct {
	x int
	y int
}

type vent struct {
	start point
	end   point
}

func main() {
	//lines := utils.InputToSlice("test.txt", "\n")
	lines := utils.InputToSlice("input.txt", "\n")
	vents := getVents(lines)

	part1(vents)
	part2(vents)
}

func getVents(lines []string) []vent {
	vents := make([]vent, len(lines))
	for i, line := range lines {
		tmp := strings.Split(line, " -> ")
		start, end := tmp[0], tmp[1]
		fn := func(c rune) bool {
			return !unicode.IsNumber(c)
		}
		startCoords := strings.FieldsFunc(start, fn)
		endCoords := strings.FieldsFunc(end, fn)
		x1, _ := strconv.Atoi(startCoords[0])
		y1, _ := strconv.Atoi(startCoords[1])
		x2, _ := strconv.Atoi(endCoords[0])
		y2, _ := strconv.Atoi(endCoords[1])
		vents[i] = vent{start: point{x: x1, y: y1}, end: point{x: x2, y: y2}}
	}
	return vents
}

func part1(vents []vent) {
	var count int
	grid := make(map[point]int)
	for _, vent := range vents {
		// only horizontal and vertical vents
		horizontal := vent.start.x == vent.end.x
		vertical := vent.start.y == vent.end.y
		if horizontal {
			if vent.start.y > vent.end.y {
				for i := vent.start.y; i >= vent.end.y; i-- {
					grid[point{x: vent.start.x, y: i}]++
				}
			} else {
				for i := vent.start.y; i <= vent.end.y; i++ {
					grid[point{x: vent.start.x, y: i}]++
				}
			}
			continue
		}
		if vertical {
			if vent.start.x > vent.end.x {
				for i := vent.start.x; i >= vent.end.x; i-- {
					grid[point{y: vent.start.y, x: i}]++
				}
			} else {
				for i := vent.start.x; i <= vent.end.x; i++ {
					grid[point{y: vent.start.y, x: i}]++
				}
			}
			continue
		}
	}
	for _, val := range grid {
		if val > 1 {
			count++
		}
	}
	fmt.Println(count)
}

func part2(vents []vent) {
	var count int
	grid := make(map[point]int)
	for _, vent := range vents {
		// only horizontal and vertical vents
		horizontal := vent.start.x == vent.end.x
		vertical := vent.start.y == vent.end.y
		if horizontal {
			if vent.start.y > vent.end.y {
				for i := vent.start.y; i >= vent.end.y; i-- {
					grid[point{x: vent.start.x, y: i}]++
				}
			} else {
				for i := vent.start.y; i <= vent.end.y; i++ {
					grid[point{x: vent.start.x, y: i}]++
				}
			}
		} else if vertical {
			if vent.start.x > vent.end.x {
				for i := vent.start.x; i >= vent.end.x; i-- {
					grid[point{y: vent.start.y, x: i}]++
				}
			} else {
				for i := vent.start.x; i <= vent.end.x; i++ {
					grid[point{y: vent.start.y, x: i}]++
				}
			}
		} else {
			xOp := "ADD"
			yOp := "ADD"
			currX := vent.start.x
			currY := vent.start.y

			if vent.start.x > vent.end.x {
				xOp = "SUBTRACT"
			}

			if vent.start.y > vent.end.y {
				yOp = "SUBTRACT"
			}

			for currY != vent.end.y && currX != vent.end.x {
				grid[point{x: currX, y: currY}]++
				if xOp == "ADD" {
					currX++
				} else {
					currX--
				}
				if yOp == "ADD" {
					currY++
				} else {
					currY--
				}
			}
			grid[point{x: vent.end.x, y: vent.end.y}]++
		}
	}

	for _, val := range grid {
		if val > 1 {
			count++
		}
	}
	fmt.Println(count)
}
