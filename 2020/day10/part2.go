package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func part2(adapters []int) {

	// Honestly, not totally sure how this works...
	// Did a lot of reading on the subreddit to come up with this
	// but still quite confused :/

	adapters = append([]int{0}, adapters...)
	sort.Ints(adapters)

	trib := make([]int, 20)
	trib[0] = 1
	trib[1] = 1
	trib[2] = 2
	var diffs []int

	for i := 3; i < 20; i++ {
		trib[i] = trib[i-1] + trib[i-2] + trib[i-3]
	}

	for i := 0; i < len(adapters)-1; i++ {
		diffs = append(diffs, adapters[i+1]-adapters[i])
	}

	var asStrings []string

	for _, d := range diffs {
		asStrings = append(asStrings, strconv.Itoa(d))
	}

	total := 1
	for _, s := range strings.Split(strings.Join(asStrings, ""), "3") {
		total *= trib[len(s)]
	}

	fmt.Println(total)
}
