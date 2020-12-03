package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strings"
)

type passwordRule struct {
	min, max         int
	letter, password string
}

func main() {
	lines, err := utils.InputToSlice("./input.txt", "\n")
	if err != nil {
		log.Fatal(err)
	}

	passwords := convertToPasswordRules(lines)

	part1(passwords)
	part2(passwords)

}

func part1(passwords []passwordRule) {
	var total int

	// check validity
	for _, p := range passwords {
		count := strings.Count(p.password, p.letter)

		if count <= p.max && count >= p.min {
			total++
		}
	}

	fmt.Printf("Part 1:\nTotal correct: %d\n\n", total)
}

func convertToPasswordRules(lines []string) (passwords []passwordRule) {
	for _, l := range lines {
		parts := strings.Split(l, ":")
		rules := strings.Split(parts[0], " ")
		min, max := getMinMax(rules[0])
		p := passwordRule{min, max, rules[1], strings.Trim(parts[1], " ")}
		passwords = append(passwords, p)
	}

	return
}

func getMinMax(s string) (min, max int) {
	m := strings.Split(s, "-")
	mInt, err := utils.StringsToInts(m)
	if err != nil {
		log.Fatal(err)
	}
	return mInt[0], mInt[1]
}
