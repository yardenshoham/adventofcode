package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

var (
	up    = image.Pt(0, -1)
	right = image.Pt(1, 0)
	down  = image.Pt(0, 1)
	left  = image.Pt(-1, 0)
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	gridBytes := bytes.Fields(input)
	grid := map[image.Point]byte{}
	var sPoint image.Point
	for y := range gridBytes {
		for x, cell := range gridBytes[y] {
			grid[image.Pt(x, y)] = cell
			if cell == 'S' {
				sPoint = image.Pt(x, y)
			}
		}
	}
	step := down
	steps := 1
	for current := sPoint.Add(step); current != sPoint; current = current.Add(step) {
		steps++
		switch grid[current] {
		case 'F':
			if step == up {
				step = right
				continue
			}
			step = down
		case '7':
			if step == right {
				step = down
				continue
			}
			step = left
		case 'J':
			if step == down {
				step = left
				continue
			}
			step = up
		case 'L':
			if step == left {
				step = up
				continue
			}
			step = right
		}
	}
	fmt.Println(steps / 2)
}
