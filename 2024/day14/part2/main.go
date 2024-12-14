package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

const (
	width  = 101
	height = 103
)

type robot struct {
	Position image.Point
	Velocity image.Point
}

// https://stackoverflow.com/a/59299881
func mod(a, b int) int {
	return (a%b + b) % b
}

func printRobots(robots []robot) {
	grid := make([][]int, height)
	for y := range grid {
		row := make([]int, width)
		for x := range row {
			for _, r := range robots {
				if r.Position.Eq(image.Pt(x, y)) {
					row[x]++
				}
			}
			if row[x] == 0 {
				fmt.Print(".")
				continue
			}
			fmt.Print("*")
		}
		fmt.Println()
	}
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	robots := []robot{}
	for scanner.Scan() {
		var px, py, vx, vy int
		fmt.Sscanf(scanner.Text(), "p=%d,%d v=%d,%d", &px, &py, &vx, &vy)
		robots = append(robots, robot{image.Pt(px, py), image.Pt(vx, vy)})
	}
	for i := range 10000 {
		for i, r := range robots {
			robots[i].Position.X = mod(r.Position.X+r.Velocity.X, width)
			robots[i].Position.Y = mod(r.Position.Y+r.Velocity.Y, height)
		}
		fmt.Println(i + 1)
		printRobots(robots)
	}
	// just run it and output to a file, then Ctrl+F "*******************************" and
	// look for the first number above it
}
