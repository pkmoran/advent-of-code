package main

import (
	"advent-of-code/utils"
	"fmt"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func reverse(s string) (r string) {
	for i := len(s) - 1; i >= 0; i-- {
		r += string(s[i])
	}
	return
}

func main() {
	program := utils.InputToSlice("./input.txt", "\n")

	part1(program)
	part2(program)
}

func part1(program []string) {
	mem := make(map[string]uint64)
	mask := ""
	for _, instruction := range program {
		isMask, _ := regexp.MatchString(`mask`, instruction)
		if isMask {
			mask = strings.Split(instruction, " = ")[1]
			continue
		}

		re := regexp.MustCompile(`(mem)\[(\d+)\] = (\d+)`)
		matches := re.FindStringSubmatch(instruction)

		loc := matches[2]
		val, _ := strconv.Atoi(matches[3])
		bin := fmt.Sprintf("%0*s", len(mask), strconv.FormatUint(uint64(val), 2))

		b := strings.Split(bin, "")

		for i, m := range mask {
			if m == 'X' {
				continue
			}

			b[i] = string(m)
		}

		masked := strings.Join(b, "")
		maskedInt, err := strconv.ParseUint(masked, 2, 36)
		if err != nil {
			log.Fatal(err)
		}

		mem[loc] = maskedInt
	}

	var total uint64 = 0
	for _, v := range mem {
		total += v
	}

	fmt.Println(total)
}
