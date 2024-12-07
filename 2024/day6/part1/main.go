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
	grid := bytes.Fields(input)
	var guard image.Point
	var step image.Point
	for y := range grid {
		x := bytes.IndexAny(grid[y], "<>^v")
		if x == -1 {
			continue
		}
		guard = image.Pt(x, y)
		switch grid[y][x] {
		case '^':
			step = up
		case '>':
			step = right
		case 'v':
			step = down
		case '<':
			step = left
		}
		break
	}

	direction := map[image.Point]image.Point{
		up:    right,
		right: down,
		down:  left,
		left:  up,
	}
	visited := make(map[image.Point]struct{})
	for 0 < guard.X && guard.X < len(grid[0])-1 && 0 < guard.Y && guard.Y < len(grid)-1 {
		visited[guard] = struct{}{}
		var next image.Point
		for next = guard.Add(step); grid[next.Y][next.X] == '#'; next = guard.Add(step) {
			step = direction[step]
		}
		guard = next
	}
	fmt.Println(len(visited) + 1)
}
