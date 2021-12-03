package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

type location struct {
	aim int
	depth int
	position int
}

func (loc *location) move(val int) {
	loc.position+= val
}
func (loc *location) up(val int) {
	loc.depth-= val
}
func (loc *location) down(val int) {
	loc.depth+= val
}
func (loc *location) move2(val int) {
	loc.position+= val
	loc.depth+= loc.aim*val
}
func (loc *location) up2(val int) {
	loc.aim-= val
}
func (loc *location) down2(val int) {
	loc.aim+= val
}

func main() {
	var loc location
	lines := utils.InputToSlice("./input.txt", "\n")

	for _, l := range lines {
		split := strings.Split(l, " ")
		direction := split[0]
		value, _ := strconv.Atoi(split[1])

		//switch direction {
		//case "forward":
		//	loc.move(value)
		//	break
		//case "up":
		//	loc.up(value)
		//	break
		//case "down":
		//	loc.down(value)
		//	break
		//}
		switch direction {
		case "forward":
			loc.move2(value)
			break
		case "up":
			loc.up2(value)
			break
		case "down":
			loc.down2(value)
			break
		}
	}
		fmt.Println(loc.depth*loc.position)

	// multiply depth and pos
}