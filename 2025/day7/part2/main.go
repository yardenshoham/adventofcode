package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

func beam(grid map[image.Point]byte, memo map[image.Point]int, x, y int) int {
	curr, ok := grid[image.Pt(x, y)]
	if !ok {
		return 0
	}
	if m, ok := memo[image.Pt(x, y)]; ok {
		return m
	}
	if curr == '^' {
		res := 1 + beam(grid, memo, x-1, y+1) + beam(grid, memo, x+1, y+1)
		memo[image.Pt(x, y)] = res
		return res
	}
	res := beam(grid, memo, x, y+1)
	memo[image.Pt(x, y)] = res
	return res
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := map[image.Point]byte{}

	var sX, sY int
	for y, line := range bytes.Fields(input) {
		for x, b := range line {
			grid[image.Pt(x, y)] = b
			if b == 'S' {
				sX = x
				sY = y
			}
		}
	}
	memo := map[image.Point]int{}
	fmt.Println(beam(grid, memo, sX, sY) + 1)
}
