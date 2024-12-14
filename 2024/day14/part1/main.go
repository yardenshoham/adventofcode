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

	for i, r := range robots {
		robots[i].Position.X = mod(r.Position.X+100*r.Velocity.X, width)
		robots[i].Position.Y = mod(r.Position.Y+100*r.Velocity.Y, height)
	}

	sums := make([]int, 4)
	for _, r := range robots {
		if r.Position.X == width/2 || r.Position.Y == height/2 {
			continue
		}
		if r.Position.X < width/2 {
			if r.Position.Y < height/2 {
				sums[0]++
				continue
			}
			sums[1]++
			continue
		}
		if r.Position.Y < height/2 {
			sums[2]++
			continue
		}
		sums[3]++
	}
	fmt.Println(sums[0] * sums[1] * sums[2] * sums[3])
}
