package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

func isCandidate(grid [][]byte, x0, y0, x1, y1 int) bool {
	return 0 <= x1 && x1 < len(grid) && 0 <= y1 && y1 < len(grid) && grid[y1][x1] == grid[y0][x0]+1
}

func rank(grid [][]byte, nines map[image.Point]struct{}, x, y int) {
	if grid[y][x] == '9' {
		nines[image.Pt(x, y)] = struct{}{}
	}
	if isCandidate(grid, x, y, x+1, y) {
		rank(grid, nines, x+1, y)
	}
	if isCandidate(grid, x, y, x-1, y) {
		rank(grid, nines, x-1, y)
	}
	if isCandidate(grid, x, y, x, y+1) {
		rank(grid, nines, x, y+1)
	}
	if isCandidate(grid, x, y, x, y-1) {
		rank(grid, nines, x, y-1)
	}
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
				nines := make(map[image.Point]struct{})
				rank(grid, nines, x, y)
				trailheads += len(nines)
			}
		}
	}
	fmt.Println(trailheads)
}
