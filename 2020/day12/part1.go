package main

import (
	"advent-of-code/utils"
	"fmt"
	"math"
	"regexp"
	"strconv"
)

type point struct {
	x, y, direction int
}

func (p *point) North(i int) {
	p.y += i
}

func (p *point) South(i int) {
	p.y -= i
}

func (p *point) East(i int) {
	p.x += i
}

func (p *point) West(i int) {
	p.x -= i
}

func (p *point) Right(i int) {
	p.direction += i

	if p.direction >= 360 {
		p.direction -= 360
	}
}

func (p *point) Left(i int) {
	p.direction -= i

	if p.direction < 0 {
		p.direction += 360
	}
}

func (p *point) Forward(i int) {
	switch p.direction {
	case 0:
		p.North(i)
	case 90:
		p.East(i)
	case 180:
		p.South(i)
	case 270:
		p.West(i)
	}
}

func main() {
	instructions := utils.InputToSlice("./input.txt", "\n")

	// instructions := []string{"F10",
	// 	"N3",
	// 	"F7",
	// 	"R90",
	// 	"F11"}

	part1(instructions)
	part2(instructions)
}

func part1(instructions []string) {
	ferry := point{0, 0, 90}

	for _, i := range instructions {
		re := regexp.MustCompile(`([NSEWRLF])(\d+)`)
		parts := re.FindStringSubmatch(i)

		instruction := parts[1]
		amp, _ := strconv.Atoi(parts[2])

		switch instruction {
		case "N":
			ferry.North(amp)
		case "S":
			ferry.South(amp)
		case "E":
			ferry.East(amp)
		case "W":
			ferry.West(amp)
		case "R":
			ferry.Right(amp)
		case "L":
			ferry.Left(amp)
		case "F":
			ferry.Forward(amp)
		}
	}

	fmt.Println(math.Abs(float64(ferry.x)) + math.Abs(float64(ferry.y)))
}
