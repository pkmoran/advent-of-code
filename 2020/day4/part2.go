package main

import (
	"fmt"
	"regexp"
	"strconv"
)

func part2(passports []map[string]string) {
	total := 0
	for _, p := range passports {
		validBirthYear := validateBirthYear(p)
		validIssueYear := validateIssueYear(p)
		validExpirationYear := validateExpirationYear(p)
		validHeight := validateHeight(p)
		validHairColor := validateHairColor(p)
		validEyeColor := validateEyeColor(p)
		validPassportID := validatePassportID(p)

		validPassport := validBirthYear && validIssueYear && validExpirationYear && validHeight && validHairColor && validEyeColor && validPassportID
		if validPassport {
			total++
		}
	}

	fmt.Println(total)
}

func validateBirthYear(p map[string]string) bool {
	by, _ := strconv.Atoi(p[birthYear])
	return by >= 1920 && by <= 2002
}

func validateIssueYear(p map[string]string) bool {
	iy, _ := strconv.Atoi(p[issueYear])
	return iy >= 2010 && iy <= 2020
}
func validateExpirationYear(p map[string]string) bool {
	ey, _ := strconv.Atoi(p[expirationYear])
	return ey >= 2020 && ey <= 2030
}

func validateHeight(p map[string]string) (validHeight bool) {
	heightRegex := regexp.MustCompile(`(\d+)((cm)|(in))`)
	hMatches := heightRegex.FindStringSubmatch(p[height])

	if len(hMatches) > 0 {
		if hMatches[2] == "in" {
			validHeight = validateHeightIn(hMatches[1])
		}

		if hMatches[2] == "cm" {
			validHeight = validateHeightCm(hMatches[1])
		}
	}
	return
}

func validateHeightIn(s string) bool {
	h, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return h >= 59 && h <= 76
}

func validateHeightCm(s string) bool {
	h, err := strconv.Atoi(s)
	if err != nil {
		return false
	}

	return h >= 150 && h <= 193
}

func validateHairColor(p map[string]string) bool {
	matched, _ := regexp.MatchString(`#(\d|[a-f]){6}`, p[hairColor])
	return matched
}

func validateEyeColor(p map[string]string) (valid bool) {
	supported := []string{"amb", "blu", "brn", "gry", "grn", "hzl", "oth"}

	for _, s := range supported {
		if p[eyeColor] == s {
			valid = true
		}
	}

	return
}

func validatePassportID(p map[string]string) (matched bool) {
	if len(p[passportID]) == 9 {
		matched, _ := regexp.MatchString(`\d+`, p[passportID])
		return matched
	}
	return false
}
