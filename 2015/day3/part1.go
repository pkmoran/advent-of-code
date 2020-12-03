package main

import (
	"advent-of-code/utils"
	"log"
)

const UP = '^'
const DOWN = 'V'
const RIGHT = '>'
const LEFT = '<'

func main() {
	moves, err := utils.InputToSlice("./input.txt", "")
	if err != nil {
		log.Fatal(err)
	}

}
