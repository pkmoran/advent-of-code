package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("./input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var nums []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			log.Fatal(err)
		}
		nums = append(nums, value)
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

done:
	for _, i := range nums {
		for _, j := range nums {
			if i+j == 2020 {
				fmt.Printf("%d * %d = %d\n", i, j, i*j)
				break done
			}
		}
	}

	part2()
}
