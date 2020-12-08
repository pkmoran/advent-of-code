package main

func part2(data map[string][]ruleSet, color string) int {
	total := 0
	for _, rule := range data[color] {
		total += rule.count + rule.count*part2(data, rule.bagType)
	}
	return total
}
