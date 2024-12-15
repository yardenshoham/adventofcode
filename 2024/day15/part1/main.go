package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := bytes.Split(input, []byte{'\n', '\n'})
	if len(parts) != 2 {
		panic("bad input")
	}
	grid := map[image.Point]byte{}
	var robot image.Point
	for y, row := range bytes.Fields(parts[0]) {
		for x, cell := range row {
			if cell == '.' {
				continue
			}
			if cell == '@' {
				robot = image.Pt(x, y)
				continue
			}
			grid[image.Pt(x, y)] = cell
		}
	}
	steps := map[byte]image.Point{
		'^': {0, -1},
		'v': {0, 1},
		'>': {1, 0},
		'<': {-1, 0},
	}
	for _, move := range bytes.ReplaceAll(parts[1], []byte{'\n'}, []byte{}) {
		var current image.Point
		var firstBox image.Point
		for current = robot.Add(steps[move]); grid[current] == 'O'; current = current.Add(steps[move]) {
			if firstBox.Eq(image.Pt(0, 0)) {
				firstBox = current
			}
		}
		if grid[current] == '#' {
			continue
		}
		if !firstBox.Eq(image.Pt(0, 0)) {
			grid[current] = 'O'
			delete(grid, firstBox)
		}
		robot = robot.Add(steps[move])
	}
	sum := 0
	for point, cell := range grid {
		if cell == 'O' {
			sum += point.Y*100 + point.X
		}
	}
	fmt.Println(sum)
}
