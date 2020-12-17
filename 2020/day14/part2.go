package main

import (
	"fmt"
	"math"
	"regexp"
	"strconv"
	"strings"
)

func part2(program []string) {
	mem := make(map[string]int)
	mask := ""
	for _, instruction := range program {
		isMask, _ := regexp.MatchString(`mask`, instruction)
		if isMask {
			mask = strings.Split(instruction, " = ")[1]
			continue
		}

		re := regexp.MustCompile(`(mem)\[(\d+)\] = (\d+)`)
		matches := re.FindStringSubmatch(instruction)

		loc, _ := strconv.Atoi(matches[2])
		bin := fmt.Sprintf("%0*s", len(mask), strconv.FormatUint(uint64(loc), 2))
		val, _ := strconv.Atoi(matches[3])

		b := strings.Split(bin, "")

		for i, m := range mask {
			if m == '0' {
				continue
			}

			b[i] = string(m)
		}

		masked := strings.Join(b, "")
		numX := strings.Count(masked, "X")

		writeCount := int(math.Pow(float64(2), float64(numX)))
		for i := 0; i < writeCount; i++ {
			option := fmt.Sprintf("%0*s", numX, strconv.FormatUint(uint64(i), 2))
			next := 0
			withOption := strings.Map(func(r rune) rune {
				if r == 'X' {
					replaced := rune(option[next])
					next++
					return replaced
				}
				return r
			}, masked)

			mem[withOption] = val
		}
	}

	total := 0
	for _, v := range mem {
		total += v
	}

	fmt.Println(total)
}
