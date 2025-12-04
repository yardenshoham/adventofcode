package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func isAccessible(grid map[image.Point]byte, x, y int) bool {
	if grid[image.Pt(x, y)] != '@' {
		return false
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
		return false
	}
	return true
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
	for {
		toRemove := map[image.Point]struct{}{}
		count := 0
		for xCheck := range y {
			for yCheck := range y {
				if isAccessible(grid, xCheck, yCheck) {
					count++
					toRemove[image.Pt(xCheck, yCheck)] = struct{}{}
				}
			}
		}
		for p := range toRemove {
			grid[image.Pt(p.X, p.Y)] = '.'
		}
		accessibleRolls += count
		if count == 0 {
			break
		}
	}
	fmt.Println(accessibleRolls)
}
