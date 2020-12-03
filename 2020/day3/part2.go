package main

import "fmt"

func part2(lines []string) {
	fmt.Println(getTreesSmashed(lines, 3, 1) * getTreesSmashed(lines, 1, 1) * getTreesSmashed(lines, 5, 1) * getTreesSmashed(lines, 7, 1) * getTreesSmashed(lines, 1, 2))
}
