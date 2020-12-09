package utils

import (
	"io/ioutil"
	"log"
	"strconv"
	"strings"
)

// InputToSlice reads a file and returns a slice of strings
func InputToSlice(filePath string, delim string) []string {
	content, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Fatal(err)
	}
	list := strings.Split(string(content), delim)

	return list
}

// StringsToInts converts a slice of strings to a slice of ints
func StringsToInts(input []string) []int {

	var output []int
	for _, i := range input {
		value, err := strconv.Atoi(i)
		if err != nil {
			log.Fatal(err)
		}
		output = append(output, value)
	}

	return output
}
