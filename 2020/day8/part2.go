package main

import "fmt"

type changeTo struct {
	name  string
	index int
	acc   int
}

func part2(instructions insMap) {
	acc := 0

	// list of jmp or nop instructions that can be changed
	var change []changeTo
	changed := false
	i := 0
	for {
		if i == len(instructions) {
			break
		}

		if instructions[i].visited {
			changed = true
			changeIndex := len(change) - 1

			// use the last item in the change list to

			// update the current loop index
			i = change[changeIndex].index
			// update the name of the instruction at that location
			instructions[i].name = change[changeIndex].name
			// reset the acc to what it was during that loop
			acc = change[changeIndex].acc
			// remove last item from change list
			change = change[:changeIndex]
			// reset the visited count so we don't end up back here accidentally
			resetVisited(instructions)
			continue
		}

		instructions[i].visited = true

		switch instructions[i].name {
		case "acc":
			acc += instructions[i].value
			i = i + 1
		case "jmp":
			if !changed {
				change = append(change, changeTo{"nop", i, acc})
			}
			i = i + instructions[i].value
		case "nop":
			if !changed {
				change = append(change, changeTo{"jmp", i, acc})
			}
			i = i + 1
		}
	}

	fmt.Println(acc)
}

func resetVisited(instructions insMap) {
	for _, v := range instructions {
		v.visited = false
	}
}
