package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
)

func (p *point) Left2(ship point, i int) {
	x := p.x - ship.x
	y := p.y - ship.y
	d := float64(i) * (math.Pi / 180)

	p.x = int(getXPrime(float64(x), float64(y), d) + float64(ship.x))
	p.y = int(getYPrime(float64(x), float64(y), d) + float64(ship.y))
}

func (p *point) Right2(ship point, i int) {
	x := p.x - ship.x
	y := p.y - ship.y
	d := float64(-i) * (math.Pi / 180)

	p.x = int(getXPrime(float64(x), float64(y), d) + float64(ship.x))
	p.y = int(getYPrime(float64(x), float64(y), d) + float64(ship.y))
}

func (p *point) Forward2(waypoint *point, i int) {
	diffX := waypoint.x - p.x
	diffY := waypoint.y - p.y

	p.x += diffX * i
	p.y += diffY * i

	waypoint.x = p.x + diffX
	waypoint.y = p.y + diffY
}

func getXPrime(x, y, d float64) float64 {
	return math.Round((x * math.Cos(d)) - (y * math.Sin(d)))
}
func getYPrime(x, y, d float64) float64 {
	return math.Round((x * math.Sin(d)) + (y * math.Cos(d)))
}

func part2(instructions []string) {
	ferry := point{0, 0, 90}
	waypoint := point{10, 1, 0}

	for _, i := range instructions {
		re := regexp.MustCompile(`([NSEWRLF])(\d+)`)
		parts := re.FindStringSubmatch(i)

		instruction := parts[1]
		amp, _ := strconv.Atoi(parts[2])

		switch instruction {
		case "N":
			waypoint.North(amp)
		case "S":
			waypoint.South(amp)
		case "E":
			waypoint.East(amp)
		case "W":
			waypoint.West(amp)
		case "R":
			waypoint.Right2(ferry, amp)
		case "L":
			waypoint.Left2(ferry, amp)
		case "F":
			ferry.Forward2(&waypoint, amp)
		}
	}

	fmt.Println(math.Abs(float64(ferry.x)) + math.Abs(float64(ferry.y)))
}
