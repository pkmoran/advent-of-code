package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.InputToSlice("input.txt", "\n")
	//gamma, _ := strconv.ParseInt(generateGammaRate(lines), 2, 64)
	//epsilon, _ := strconv.ParseInt(generateEpsilonRate(lines), 2, 64)
	//fmt.Println(gamma * epsilon)

	o2, _ := strconv.ParseInt(generateO2GeneratorRating(lines), 2, 64)
	co2, _ := strconv.ParseInt(generateCO2GeneratorRating(lines), 2, 64)
	fmt.Println(o2, co2)
	fmt.Println(o2 * co2)
}

func bitsToGrid(bitSlice []string) ([][]string, int) {
	var grid [][]string
	for _, binary := range bitSlice {
		bits := strings.Split(binary, "")
		grid = append(grid, bits)
	}
	return grid, len(grid[0])
}

func generateGammaRate(bitSlice []string) string {
	// find MOST frequent bit in each column and return binary number from those bits
	var gammaRate string
	grid, digitLength := bitsToGrid(bitSlice)

	var countZeros int
	var countOnes int
	currentPosition := 0
	for currentPosition < digitLength {
		for _, row := range grid {
			switch row[currentPosition] {
			case "0":
				countZeros++
				break
			case "1":
				countOnes++
			}
		}
		if countOnes >= countZeros {
			gammaRate += "1"
		} else {
			gammaRate += "0"
		}

		countZeros = 0
		countOnes = 0
		currentPosition++
	}
	fmt.Println(gammaRate)
	return gammaRate
}

func generateO2GeneratorRating(bitSlice []string) string {
	grid, digitLength := bitsToGrid(bitSlice)

	var zeros []int
	var ones []int
	indicesToCheck := make([]int, len(bitSlice))
	for i := range indicesToCheck {
		indicesToCheck[i] = i
	}
	currentPosition := 0
	for currentPosition < digitLength {
		for _, index := range indicesToCheck {
			bit := grid[index][currentPosition]
			if bit == "0" {
				zeros = append(zeros, index)
				continue
			}
			if bit == "1" {
				ones = append(ones, index)
			}
		}
		if len(ones) >= len(zeros) {
			indicesToCheck = ones
		} else {
			indicesToCheck = zeros
		}

		zeros = []int{}
		ones = []int{}
		currentPosition++
	}
	o2GeneratorRating := strings.Join(grid[indicesToCheck[0]], "")
	fmt.Println(o2GeneratorRating)
	return o2GeneratorRating
}
func generateCO2GeneratorRating(bitSlice []string) string {
	grid, digitLength := bitsToGrid(bitSlice)

	var zeros []int
	var ones []int
	indicesToCheck := make([]int, len(bitSlice))
	for i := range indicesToCheck {
		indicesToCheck[i] = i
	}
	currentPosition := 0
	fmt.Println(len(indicesToCheck))
	for currentPosition < digitLength {
		for _, index := range indicesToCheck {
			bit := grid[index][currentPosition]
			if bit == "0" {
				zeros = append(zeros, index)
				continue
			}
			if bit == "1" {
				ones = append(ones, index)
			}
		}
		if len(zeros) <= len(ones) {
			if len(zeros) > 0 {
				indicesToCheck = zeros
			}
		} else {
			if len(ones) > 0 {
				indicesToCheck = ones
			}
		}

		zeros = []int{}
		ones = []int{}
		currentPosition++
	}
	co2GeneratorRating := strings.Join(grid[indicesToCheck[0]], "")
	fmt.Println(co2GeneratorRating)
	return co2GeneratorRating
}

func generateEpsilonRate(bitSlice []string) string {
	// find LEAST frequent bit in each column and return binary number from those bits
	var epsilonRate string
	grid, digitLength := bitsToGrid(bitSlice)

	var countZeros int
	var countOnes int
	currentPosition := 0
	for currentPosition < digitLength {
		for _, row := range grid {
			switch row[currentPosition] {
			case "0":
				countZeros++
				break
			case "1":
				countOnes++
			}
		}
		if countOnes >= countZeros {
			epsilonRate += "0"
		} else {
			epsilonRate += "1"
		}

		countZeros = 0
		countOnes = 0
		currentPosition++
	}
	fmt.Println(epsilonRate)
	return epsilonRate
}
