package main

import (
	"fmt"
	"math/rand"
	"strings"

	"github.com/eiannone/keyboard"
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
		for c := range len(board[r]) {
			row += fmt.Sprintf(" %3v |", board[r][c])
		}
		boardStr += generateBorder(len(row)) + "\n" + row + "\n"
	}
	boardStr += generateBorder(len(row)) + "\n"

	// clear screen
	fmt.Print("\033[H\033[2J")
	fmt.Println("Press ESC or q to quit")
	fmt.Print(boardStr)
}

func randomizeBoard(board Board) {
	for range 100 {
		r1, c1 := rand.Intn(len(board)), rand.Intn(len(board))
		r2, c2 := rand.Intn(len(board)), rand.Intn(len(board))

		board[r1][c1], board[r2][c2] = board[r2][c2], board[r1][c1]
	}
}

func getEmptyPos(board Board) (int, int) {
	for r := range len(board) {
		for c := range len(board[r]) {
			if board[r][c] == "" {
				return r, c
			}
		}
	}

	return -1, -1
}

func move(board Board, dir string) {
	er, ec := getEmptyPos(board)
	r, c := -1, -1

	switch dir {
	case "U":
		// grab element from below to swap with
		r, c = er+1, ec

		if r >= len(board) {
			return
		}
	case "D":
		// grab element from above to swap with
		r, c = er-1, ec
		if r < 0 {
			return
		}
	case "L":
		// grab element from the right to swap with
		r, c = er, ec+1
		if c >= len(board[r]) {
			return
		}
	case "R":
		// grab element from the left to swap with
		r, c = er, ec-1
		if c < 0 {
			return
		}
	}

	board[r][c], board[er][ec] = board[er][ec], board[r][c]
}

func isSolved(board Board, solvedBoard Board) bool {
	for r := range len(board) {
		for c := range len(board[r]) {
			if board[r][c] != solvedBoard[r][c] {
				return false
			}
		}
	}

	return true
}

func main() {
	size := 3
	board := createBoard(size)
	solvedBoard := createBoard(size)
	randomizeBoard(board)

	if err := keyboard.Open(); err != nil {
		panic(err)
	}
	defer func() {
		_ = keyboard.Close()
	}()

	for {

		// Game loop
		printBoard(board)

		char, key, err := keyboard.GetKey()
		if err != nil {
			panic(err)
		}
		// fmt.Printf("You pressed: rune %q, key %X\r\n", char, key)
		if key == keyboard.KeyEsc || char == 'q' {
			break
		}

		switch char {
		case 'w':
			move(board, "U")
		case 'a':
			move(board, "L")
		case 's':
			move(board, "D")
		case 'd':
			move(board, "R")
		}

		if isSolved(board, solvedBoard) {
			printBoard(board)
			fmt.Printf("You solved %v x %v board!\n", size, size)
			break
		}

	}
}
