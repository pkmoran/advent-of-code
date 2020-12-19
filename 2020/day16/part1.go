package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type rule struct {
	min, max int
}

type ticketCategory struct {
	rules   []rule
	invalid bool
}

func main() {
	noteSections := utils.InputToSlice("./input.example.txt", "\n\n")
	ticketCategories := noteSections[0]
	myTicket := noteSections[1]
	nearbyTickets := noteSections[2]

	// part1(ticketCategories, nearbyTickets)
	part2(ticketCategories, nearbyTickets, myTicket)
}

func part2(ticketCategories, nearbyTickets string, myTicket string) {
	ruleMap := createRuleMap(strings.Split(ticketCategories, "\n"))

	nt := strings.Split(nearbyTickets, "\n")[1:]
	validTickets := getValidTickets(nt, ruleMap)

	// rule.invalid for each column
	// so all start as valid (!NOTValid)
	// remove once a value does not match a rule
	// there should only be one remaining valid value

	// categoryPositions := make(map[string]int)
	ticketCategoryIndex := 0
	for {
		for _, ticket := range validTickets {
			value := ticket[ticketCategoryIndex]

			// ruleCheck:
			for categoryName, ruleCategory := range ruleMap {
				for _, rule := range ruleCategory.rules {
					if value >= rule.min && value <= rule.max {
						ruleCategory.invalid = true
						// break ruleCheck
					}
				}
				fmt.Println(categoryName, ruleCategory.invalid)
			}
		}
		fmt.Println(ruleMap)
		// resetCategoryValidity(ruleMap)
		ticketCategoryIndex++
		if ticketCategoryIndex == len(validTickets[0]) {
			break
		}
	}

	// categoryPositions := getCategoryPositions(validTickets[0], ruleMap)
	// mt := utils.StringsToInts(strings.Split(strings.Split(myTicket, "\n")[1], ","))

	// var total int64 = 1
	// for category, position := range categoryPositions {
	// 	if strings.HasPrefix(category, "departure") {
	// 		fmt.Println(category, mt[position])
	// 		total *= int64(mt[position])
	// 	}
	// }

	// fmt.Println(total)
	fmt.Println(ruleMap)
}

func part1(ticketCategories, nearbyTickets string) {
	ruleMap := createRuleMap(strings.Split(ticketCategories, "\n"))

	nt := strings.Split(nearbyTickets, "\n")[1:]
	invalidValues := getInvalidValues(nt, ruleMap)

	fmt.Println(sum(invalidValues))
}

func sum(values []int) (total int) {
	for _, v := range values {
		total += v
	}
	return
}

func (c *ticketCategory) addRule(rule rule) {
	c.rules = append(c.rules, rule)
}

func createRuleMap(categories []string) map[string]*ticketCategory {
	ruleMap := make(map[string]*ticketCategory)
	for _, c := range categories {
		split := strings.Split(c, ": ")
		category := split[0]
		ruleString := split[1]

		rules := strings.Split(ruleString, " or ")

		for _, r := range rules {
			minMax := strings.Split(r, "-")
			min, _ := strconv.Atoi(minMax[0])
			max, _ := strconv.Atoi(minMax[1])
			ru := rule{min, max}
			ruleMap[category] = &ticketCategory{rules: []rule{}, invalid: false}
			ruleMap[category].rules = append(ruleMap[category].rules, ru)
		}
	}

	return ruleMap
}

func getInvalidValues(values []string, ruleMap map[string]*ticketCategory) (invalid []int) {
	for _, t := range values {
		vals := utils.StringsToInts(strings.Split(t, ","))

		for _, val := range vals {
			valid := false

		ruleCheck:
			for _, ruleCategory := range ruleMap {
				for _, rule := range ruleCategory.rules {
					if val >= rule.min && val <= rule.max {
						valid = true
						break ruleCheck
					}
				}

			}

			if !valid {
				invalid = append(invalid, val)
			}
		}
	}
	return
}

func getValidTickets(tickets []string, ruleMap map[string]*ticketCategory) (validTickets [][]int) {
	for _, t := range tickets {
		vals := utils.StringsToInts(strings.Split(t, ","))
		var matchingValues []int
		for _, val := range vals {
		ruleCheck:
			for _, ruleCategory := range ruleMap {
				for _, rule := range ruleCategory.rules {
					if val >= rule.min && val <= rule.max {
						matchingValues = append(matchingValues, val)
						break ruleCheck
					}
				}
			}
		}
		if len(matchingValues) == len(vals) {
			validTickets = append(validTickets, vals)
		}
	}
	return
}

func getCategoryPositions(ticket []int, ruleMap map[string]*ticketCategory) map[string]int {
	categoryPositions := make(map[string]int)
	for position, value := range ticket {
	ruleCheck:
		for categoryName, category := range ruleMap {
			for _, rule := range category.rules {
				if value >= rule.min && value <= rule.max {
					if _, ok := categoryPositions[categoryName]; ok {
						continue ruleCheck
					}
					categoryPositions[categoryName] = position
					break ruleCheck
				}
			}
		}
	}
	return categoryPositions
}
