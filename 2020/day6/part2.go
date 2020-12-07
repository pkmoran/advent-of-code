package main

import "strings"

func part2(group string) (groupTotal int) {
	num := len(strings.Split(group, "\n"))
	checkMap := make(map[string]int)

	for _, l := range strings.ReplaceAll(group, "\n", "") {
		if _, seen := checkMap[string(l)]; seen {
			checkMap[string(l)]++
			continue
		}
		checkMap[string(l)] = 1
	}

	for _, v := range checkMap {
		if v == num {
			groupTotal++
		}
	}
	return
}
