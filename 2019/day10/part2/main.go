package main

import (
	"bufio"
	"cmp"
	"fmt"
	"image"
	"math"
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

func canSee(points map[line][]image.Point, station image.Point) map[image.Point]struct{} {
	asteroidMap := map[image.Point]struct{}{}
	for _, asteroids := range points {
		for i, asteroid := range asteroids {
			if asteroid != station {
				continue
			}
			if i > 0 {
				asteroidMap[asteroids[i-1]] = struct{}{}
			}
			if i < len(asteroids)-1 {
				asteroidMap[asteroids[i+1]] = struct{}{}
			}
		}
	}
	return asteroidMap
}

type pointWithRad struct {
	p image.Point
	r float64
}

func firstToDestroy(pwrad []pointWithRad) int {
	for i, a := range pwrad {
		if a.r >= -math.Pi/2 {
			return i
		}
	}
	return 0
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
	var station image.Point
	var mostAsteroids int
	for s, a := range counts {
		if a > mostAsteroids {
			station = s
			mostAsteroids = a
		}
	}
	var destroyed int
	for {
		destroyed = destroyAsteroids(n, grid, station, destroyed)
		if destroyed == 200 {
			break
		}
	}
}

func destroyAsteroids(n int, grid map[image.Point]byte, station image.Point, destroyed int) int {
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
	asteroidMap := canSee(points, station)

	pwrad := []pointWithRad{}
	for a := range asteroidMap {
		rad := math.Atan2(float64(a.Y-station.Y), float64(a.X-station.X))
		pwrad = append(pwrad, pointWithRad{p: a, r: rad})
	}
	slices.SortFunc(pwrad, func(a, b pointWithRad) int {
		return cmp.Compare(a.r, b.r)
	})

	ftd := firstToDestroy(pwrad)
	for i := ftd; i < len(pwrad)+ftd; i++ {
		point := pwrad[i%len(pwrad)].p
		grid[point] = '.'
		destroyed++
		if destroyed == 200 {
			fmt.Println(100*point.X + point.Y)
			return destroyed
		}
	}
	return destroyed
}
