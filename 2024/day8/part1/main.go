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
	antennas := make(map[byte][]image.Point)
	grid := bytes.Fields(input)
	for y := range grid {
		for x := range grid[y] {
			if grid[y][x] != '.' {
				antennas[grid[y][x]] = append(antennas[grid[y][x]], image.Pt(x, y))
			}
		}
	}
	antiNodes := make(map[image.Point]struct{})
	for _, freq := range antennas {
		for i := range freq {
			for j := i + 1; j < len(freq); j++ {
				d := freq[i].Sub(freq[j])
				minAntiNode := freq[i].Add(d)
				maxAntiNode := freq[j].Sub(d)
				if 0 <= minAntiNode.X && minAntiNode.X < len(grid) && 0 <= minAntiNode.Y && minAntiNode.Y < len(grid) {
					antiNodes[minAntiNode] = struct{}{}
				}
				if 0 <= maxAntiNode.X && maxAntiNode.X < len(grid) && 0 <= maxAntiNode.Y && maxAntiNode.Y < len(grid) {
					antiNodes[maxAntiNode] = struct{}{}
				}
			}
		}
	}

	fmt.Println(len(antiNodes))
}