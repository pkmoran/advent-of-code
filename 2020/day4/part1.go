package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"strings"
)

// type passport map[string]string

const birthYear = "byr"
const issueYear = "iyr"
const expirationYear = "eyr"
const height = "hgt"
const hairColor = "hcl"
const eyeColor = "ecl"
const passportID = "pid"
const countryID = "cid"

func main() {
	lines, err := utils.InputToSlice("./input.txt", "\n\n")
	if err != nil {
		log.Fatal(err)
	}

	var passports []map[string]string

	for _, l := range lines {
		fields := strings.Fields(l)

		var p = make(map[string]string)

		for _, f := range fields {
			pair := strings.Split(f, ":")

			if pair[0] == countryID {
				continue
			}
			p[pair[0]] = pair[1]
		}

		passports = append(passports, p)
	}

	part1(passports)
	part2(passports)
}

func part1(passports []map[string]string) {
	total := 0
	for _, p := range passports {
		if len(p) == 7 {
			total++
		}
	}

	fmt.Println(total)
}
