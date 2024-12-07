package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
	"slices"
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
	obstructions := 0
	for y := range grid {
		for x := range grid[y] {
			if image.Pt(x, y) == guard {
				continue
			}
			tmpGuard := guard
			tmpStep := step
			visited := make(map[image.Point][]image.Point)
			for 0 <= tmpGuard.X && tmpGuard.X < len(grid[0]) && 0 <= tmpGuard.Y && tmpGuard.Y < len(grid) {
				var next image.Point
				for next = tmpGuard.Add(tmpStep); (0 <= next.X && next.X < len(grid) && 0 <= next.Y && next.Y < len(grid) && grid[next.Y][next.X] == '#') || image.Pt(x, y) == next; next = tmpGuard.Add(tmpStep) {
					tmpStep = direction[tmpStep]
				}
				if slices.Contains(visited[tmpGuard], tmpStep) {
					obstructions++
					break
				}
				visited[tmpGuard] = append(visited[tmpGuard], tmpStep)
				tmpGuard = next
			}
		}
	}
	fmt.Println(obstructions)
}
