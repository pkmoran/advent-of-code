package main

import (
	"fmt"
	"strconv"
	"strings"
)

func part2(dimensions []string) {
	total := 0

	for _, d := range dimensions {
		b := strings.Split(d, "x")

		var gift box
		gift.length, _ = strconv.Atoi(b[0])
		gift.width, _ = strconv.Atoi(b[1])
		gift.height, _ = strconv.Atoi(b[2])

		p1, p2, p3 := perimeter(gift.length, gift.width), perimeter(gift.length, gift.height), perimeter(gift.width, gift.height)

		// volume + smallest perimeter
		a := (gift.length * gift.width * gift.height) + min(p1, p2, p3)

		total += a
	}

	fmt.Printf("Total Part 2: %d", total)
}

func perimeter(a int, b int) int {
	return 2 * (a + b)
}
