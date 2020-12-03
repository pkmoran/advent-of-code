package utils

import (
	"io/ioutil"
	"strconv"
	"strings"
)

// InputToSlice reads a file and returns a slice of strings
func InputToSlice(filePath string, delim string) ([]string, error) {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		return nil, err
	}
	list := strings.Split(string(content), delim)

	return list, nil
}

// StringsToInts converts a slice of strings to a slice of ints
func StringsToInts(input []string) ([]int, error) {

	var output []int
	for _, i := range input {
		value, err := strconv.Atoi(i)
		if err != nil {
			return nil, err
		}
		output = append(output, value)
	}

	return output, nil
}
