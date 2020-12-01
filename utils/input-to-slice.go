package utils

import (
	"bufio"
	"os"
	"strconv"
)

// InputToSlice reads a file and returns a slice of ints (expects '\n' as delineator)
func InputToSlice(filePath string) ([]int, error) {
	var nums []int

	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return nil, err
		}
		nums = append(nums, value)
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}

	return nums, nil
}
