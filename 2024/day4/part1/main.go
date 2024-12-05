package main

import (
	"bytes"
	"fmt"
	"os"
)

var XMAS = []byte{'X', 'M', 'A', 'S'}

func transpose(board [][]byte) [][]byte {
	xl := len(board[0])
	yl := len(board)
	result := make([][]byte, xl)
	for i := range result {
		result[i] = make([]byte, yl)
	}
	for i := 0; i < xl; i++ {
		for j := 0; j < yl; j++ {
			result[i][j] = board[j][i]
		}
	}
	return result
}

func countHorizontal(board [][]byte) int {
	sum := 0
	for _, line := range board {
		for i := 0; i < len(line)-3; i++ {
			window := line[i : i+4]
			reversedWindow := []byte{window[3], window[2], window[1], window[0]}
			if bytes.Equal(window, XMAS) || bytes.Equal(reversedWindow, XMAS) {
				sum++
			}
		}
	}
	return sum
}

func countVertical(board [][]byte) int {
	transposed := transpose(board)
	return countHorizontal(transposed)
}

func countDiagonal(board [][]byte) int {
	sum := 0
	for i := 0; i < len(board)-3; i++ {
		for j := 0; j < len(board[0])-3; j++ {
			diagonal := []byte{board[i][j], board[i+1][j+1], board[i+2][j+2], board[i+3][j+3]}
			reversedDiagonal := []byte{diagonal[3], diagonal[2], diagonal[1], diagonal[0]}
			if bytes.Equal(diagonal, XMAS) || bytes.Equal(reversedDiagonal, XMAS) {
				sum++
			}
		}
		for j := 3; j < len(board[0]); j++ {
			diagonal := []byte{board[i][j], board[i+1][j-1], board[i+2][j-2], board[i+3][j-3]}
			reversedDiagonal := []byte{diagonal[3], diagonal[2], diagonal[1], diagonal[0]}
			if bytes.Equal(diagonal, XMAS) || bytes.Equal(reversedDiagonal, XMAS) {
				sum++
			}
		}
	}
	return sum
}

func main() {
	puzzle, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	board := bytes.Fields(puzzle)
	fmt.Println(countHorizontal(board) + countVertical(board) + countDiagonal(board))
}
