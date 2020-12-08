package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

type ruleSet struct {
	count   int
	bagType string
}

const testColor = "shiny gold"

func main() {
	rules, err := utils.InputToSlice("./input.txt", ".\n")
	if err != nil {
		log.Fatal(err)
	}

	ruleData := make(map[string][]ruleSet)
	for _, r := range rules {
		buildRuleData(r, ruleData)
	}

	fmt.Println(part1(ruleData))
	fmt.Println(part2(ruleData, testColor))
}

func buildRuleData(rule string, data map[string][]ruleSet) {
	split := strings.Split(rule, "bags contain")

	mainType := strings.TrimSpace(split[0])

	if split[1] == "no other bags" {
		data[mainType] = make([]ruleSet, 0)
		return
	}

	bagTypes := getBagTypes(split[1], mainType)

	data[mainType] = bagTypes
}

func getBagTypes(s string, parent string) (set []ruleSet) {
	re := regexp.MustCompile(`(\d) (\w+ \w+) bags?,?`)
	matches := re.FindAllStringSubmatch(s, -1)

	for _, m := range matches {
		count, err := strconv.Atoi(m[1])
		if err != nil {
			log.Fatal(err)
		}

		set = append(set, ruleSet{count: count, bagType: m[2]})
	}
	return
}

func bagContainsColor(rules []ruleSet, data map[string][]ruleSet) bool {

	for _, rule := range rules {
		if rule.bagType == testColor {
			return true
		}
		if bagContainsColor(data[rule.bagType], data) {
			return true
		}
	}
	return false
}

func part1(ruleData map[string][]ruleSet) int {
	total := 0
	for color, rules := range ruleData {
		if color == testColor || len(rules) == 0 {
			continue
		}

		if contains := bagContainsColor(rules, ruleData); contains {
			total++
		}
	}
	return total
}
