package main

import (
	"bytes"
	"fmt"
	"image"
	"os"
)

func inBounds(grid [][]byte, p image.Point) bool {
	return 0 <= p.X && p.X < len(grid) && 0 <= p.Y && p.Y < len(grid)
}

func fencePrice(region map[image.Point]struct{}) int {
	perimeter := 0
	for plot := range region {
		for _, step := range []image.Point{{0, -1}, {0, 1}, {1, 0}, {-1, 0}} {
			if _, ok := region[plot.Add(step)]; !ok {
				perimeter++
			}
		}
	}
	return len(region) * perimeter
}

func travel(grid [][]byte, point image.Point, region map[image.Point]struct{}, plot byte) {
	if !inBounds(grid, point) || grid[point.Y][point.X] != plot {
		return
	}
	region[point] = struct{}{}
	grid[point.Y][point.X] = 0
	for _, step := range []image.Point{{0, -1}, {0, 1}, {1, 0}, {-1, 0}} {
		travel(grid, point.Add(step), region, plot)
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	regions := []map[image.Point]struct{}{}
	grid := bytes.Fields(input)
	for y := range grid {
		for x, plot := range grid[y] {
			if plot != 0 {
				region := make(map[image.Point]struct{})
				travel(grid, image.Pt(x, y), region, plot)
				regions = append(regions, region)
			}
		}
	}
	sum := 0
	for _, region := range regions {
		sum += fencePrice(region)
	}
	fmt.Println(sum)
}
