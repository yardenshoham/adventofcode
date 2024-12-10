package main

import (
	"bytes"
	"fmt"
	"os"
)

func isCandidate(grid [][]byte, x0, y0, x1, y1 int) bool {
	return 0 <= x1 && x1 < len(grid) && 0 <= y1 && y1 < len(grid) && grid[y1][x1] == grid[y0][x0]+1
}

func rank(grid [][]byte, x, y int) int {
	if grid[y][x] == '9' {
		return 1
	}
	result := 0
	if isCandidate(grid, x, y, x+1, y) {
		result += rank(grid, x+1, y)
	}
	if isCandidate(grid, x, y, x-1, y) {
		result += rank(grid, x-1, y)
	}
	if isCandidate(grid, x, y, x, y+1) {
		result += rank(grid, x, y+1)
	}
	if isCandidate(grid, x, y, x, y-1) {
		result += rank(grid, x, y-1)
	}
	return result
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	trailheads := 0
	grid := bytes.Fields(input)
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] == '0' {
				trailheads += rank(grid, x, y)
			}
		}
	}
	fmt.Println(trailheads)
}
