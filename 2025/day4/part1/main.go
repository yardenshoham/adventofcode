package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func accessible(grid map[image.Point]byte, x, y int) int {
	if grid[image.Pt(x, y)] != '@' {
		return 0
	}
	adjacent := 0
	for _, xCheck := range []int{-1, 0, 1} {
		for _, yCheck := range []int{-1, 0, 1} {
			if xCheck == 0 && yCheck == 0 {
				continue
			}
			if grid[image.Pt(x+xCheck, y+yCheck)] == '@' {
				adjacent++
			}
		}
	}
	if adjacent >= 4 {
		return 0
	}
	return 1
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	grid := map[image.Point]byte{}
	scanner := bufio.NewScanner(file)
	y := 0
	for scanner.Scan() {
		for x, b := range scanner.Bytes() {
			grid[image.Pt(x, y)] = b
		}
		y++
	}
	accessibleRolls := 0
	for xCheck := range y {
		for yCheck := range y {
			accessibleRolls += accessible(grid, xCheck, yCheck)
		}
	}
	fmt.Println(accessibleRolls)
}
