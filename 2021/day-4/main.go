package main

import (
	"advent-of-code/utils"
	"fmt"
	"strconv"
	"strings"
)

const winner = "XXXXX"

func main() {
	lines := utils.InputToSlice("input.txt", "\n\n")
	//lines := utils.InputToSlice("test.txt", "\n\n")
	numbersToCall := strings.Split(lines[0], ",")
	grids := lines[1:]

	boards := make([][]string, len(grids))
	for i, grid := range grids {
		gridRows := strings.Split(grid, "\n")
		boards[i] = make([]string, len(gridRows))
		for j, row := range gridRows {
			boards[i][j] = strings.Replace(strings.TrimSpace(row), "  ", " ", -1)
		}
	}

	part1(boards, numbersToCall)
	part2(boards, numbersToCall)
}

func part1(boards [][]string, numbersToCall []string) {
	for i, calledNumber := range numbersToCall {
		numbersCalled := i + 1
		for _, board := range boards {
			updatedBoard := updateBoard(board, calledNumber)
			if numbersCalled >= len(board) {
				bingo, score := checkForBingo(updatedBoard, calledNumber)
				if bingo {
					fmt.Println(score)
					return
				}
			}
		}
	}
}

func part2(boards [][]string, numbersToCall []string) {
	lastBingoIndex := 0
	lastCalledNumberThatWon := ""
	winningBoards := make(map[int]bool)
	for i, calledNumber := range numbersToCall {
		numbersCalled := i + 1
		for j, board := range boards {
			_, alreadyWon := winningBoards[j]
			if alreadyWon {
				continue
			}
			updatedBoard := updateBoard(board, calledNumber)
			if numbersCalled >= len(board) {
				bingo, _ := checkForBingo(updatedBoard, calledNumber)
				if bingo {
					winningBoards[j] = true
					lastBingoIndex = j
					lastCalledNumberThatWon = calledNumber
				}
			}
		}
	}
	_, score := checkForBingo(boards[lastBingoIndex], lastCalledNumberThatWon)
	fmt.Println(score)
}

func updateBoard(board []string, calledNumber string) []string {
	for i, row := range board {
		tmp := strings.Fields(row)
		for j, num := range tmp {
			if num == calledNumber {
				tmp[j] = "X"
			}
		}
		board[i] = strings.Join(tmp, " ")
	}

	return board
}

func checkForBingo(board []string, calledNumber string) (bool, int) {
	if checkRows(board) || checkColumns(board) || checkDiagonal(board) {
		calledInt, _ := strconv.Atoi(calledNumber)
		return true, getScore(board, calledInt)
	}

	return false, 0
}

func checkRows(board []string) bool {
	for _, row := range board {
		if strings.Compare(strings.Join(strings.Fields(row), ""), winner) == 0 {
			return true
		}
	}
	return false
}

func checkColumns(board []string) bool {
	col := 0
	colLen := len(strings.Fields(board[0]))
	colVals := ""
	for col < colLen {
		for _, row := range board {
			colVals += strings.Fields(row)[col]
		}
		if strings.Compare(colVals, winner) == 0 {
			return true
		}
		colVals = ""
		col++
	}
	return false
}

func checkDiagonal(board []string) bool {
	row := 0
	col := 0
	colLen := len(strings.Fields(board[0]))
	vals := ""
	for col < colLen {
		vals += strings.Fields(board[row])[col]
		row++
		col++
	}

	if strings.Compare(vals, winner) == 0 {
		return true
	}

	row = 0
	col = colLen - 1

	for col >= 0 {
		vals += strings.Fields(board[row])[col]
		row++
		col--
	}

	if strings.Compare(vals, winner) == 0 {
		return true
	}
	return false
}

func getScore(board []string, calledNumber int) int {
	score := 0
	for _, row := range board {
		for _, square := range strings.Fields(row) {
			if square != "X" {
				val, _ := strconv.Atoi(square)
				score += val
			}
		}
	}
	return score * calledNumber
}
