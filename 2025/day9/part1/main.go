package main

import (
	"bufio"
	"cmp"
	"fmt"
	"image"
	"os"
	"slices"
)

func area(r image.Rectangle) int {
	return (r.Dx() + 1) * (r.Dy() + 1)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	points := []image.Point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p image.Point
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d", &p.X, &p.Y)
		if err != nil {
			panic(err)
		}
		points = append(points, p)
	}
	rectangles := []image.Rectangle{}
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			rectangles = append(rectangles, image.Rect(p.X, p.Y, points[j].X, points[j].Y))
		}
	}
	result := slices.MaxFunc(rectangles, func(a, b image.Rectangle) int { return cmp.Compare(area(a), area(b)) })
	fmt.Println(area(result))
}
