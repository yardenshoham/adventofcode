package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	puzzle, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	board := bytes.Fields(puzzle)
	sum := 0
	for i := 0; i < len(board)-2; i++ {
		for j := 0; j < len(board[0])-2; j++ {
			if board[i+1][j+1] == 'A' && (board[i][j] == 'M' && board[i+2][j+2] == 'S' || board[i][j] == 'S' && board[i+2][j+2] == 'M') && (board[i+2][j] == 'M' && board[i][j+2] == 'S' || board[i+2][j] == 'S' && board[i][j+2] == 'M') {
				sum++
			}
		}
	}
	fmt.Println(sum)
}
