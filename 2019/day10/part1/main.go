package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
	"slices"
	"strconv"
)

type line struct {
	m string
	b string
}

func getLine(x1, y1, x2, y2 int, grid map[image.Point]byte) (line, bool) {
	if grid[image.Pt(x1, y1)] != '#' || grid[image.Pt(x2, y2)] != '#' {
		return line{}, false
	}
	if x1 == x2 {
		return line{m: "inf", b: strconv.Itoa(x1)}, true
	}
	m := float32(y2-y1) / float32(x2-x1)
	b := float32(y1) - m*float32(x1)
	M := fmt.Sprintf("%.4f", m)
	B := fmt.Sprintf("%.4f", b)
	if M == "-0.0000" {
		M = "0.0000"
	}
	if B == "-0.0000" {
		B = "0.0000"
	}
	return line{m: M, b: B}, true
}

func updatePoints(x1, y1, x2, y2 int, l line, points map[line][]image.Point) {
	for _, p := range []image.Point{image.Pt(x1, y1), image.Pt(x2, y2)} {
		if !slices.Contains(points[l], p) {
			points[l] = append(points[l], p)
		}
	}
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
		line := scanner.Bytes()
		for x, b := range line {
			grid[image.Pt(x, y)] = b
		}
		y++
	}
	n := y + 1
	points := map[line][]image.Point{}
	for y1 := range n {
		for x1 := range n {
			for y2 := range n {
				for x2 := range n {
					l, found := getLine(x1, y1, x2, y2, grid)
					if !found {
						continue
					}
					updatePoints(x1, y1, x2, y2, l, points)
				}
			}
		}
	}
	counts := map[image.Point]int{}
	for _, asteroids := range points {
		for i, asteroid := range asteroids {
			count := 0
			if i > 0 {
				count++
			}
			if i < len(asteroids)-1 {
				count++
			}
			counts[asteroid] += count
		}
	}
	var mostAsteroids int
	for _, a := range counts {
		mostAsteroids = max(mostAsteroids, a)
	}
	fmt.Println(mostAsteroids)
}
