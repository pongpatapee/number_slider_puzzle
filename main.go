package main

import (
	"fmt"
	"strings"
)

type Board [][]string

func createBoard(size int) Board {
	board := make(Board, size)

	for i := range size {
		row := make([]string, size)
		for j := range size {
			if i == size-1 && j == size-1 {
				row[j] = ""
				continue
			}
			row[j] = fmt.Sprintf("%d", (j+1)+i*size)
		}
		board[i] = row
	}

	return board
}

func generateBorder(length int) string {
	if length < 2 {
		return strings.Repeat("-", length)
	}

	border := "+" + strings.Repeat("-", length-2) + "+"

	return border
}

/*
+-----------+
| 1 | 2 | 3 |
+-----------+
| 4 | 5 | 6 |
+-----------+
| 7 | 8 | 9 |
+-----------+
*/
func printBoard(board Board) {
	boardStr := ""
	row := ""

	for r := range len(board) {
		row = "|"
		for c := range len(board) {
			row += fmt.Sprintf(" %3v |", board[r][c])
		}
		boardStr += generateBorder(len(row)) + "\n" + row + "\n"
	}
	boardStr += generateBorder(len(row)) + "\n"

	fmt.Print(boardStr)
}

func main() {
	board := createBoard(5)
	printBoard(board)
}
