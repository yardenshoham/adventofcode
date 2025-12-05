package main

import (
	"fmt"
	"image"
	"os"
	"slices"
	"strings"
)

func isOverlapping(r1, r2 image.Point) bool {
	sub := min(r1.Y, r2.Y) - max(r1.X, r2.X)
	return sub >= 0
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(input), "\n\n")
	ranges := []image.Point{}
	for _, r := range strings.Fields(parts[0]) {
		var from, to int
		_, err := fmt.Sscanf(r, "%d-%d", &from, &to)
		if err != nil {
			panic(err)
		}
		ranges = append(ranges, image.Pt(from, to))
	}
	slices.SortFunc(ranges, func(a, b image.Point) int {
		return a.X - b.X
	})
	distinctRanges := []image.Point{}
	to := ranges[0].Y
	from := ranges[0].X
	for _, r := range ranges {
		if isOverlapping(image.Pt(from, to), r) {
			from = min(from, r.X)
			to = max(to, r.Y)
			continue
		}
		distinctRanges = append(distinctRanges, image.Pt(from, to))
		from = r.X
		to = r.Y
	}
	distinctRanges = append(distinctRanges, image.Pt(from, to))
	fresh := 0
	for _, r := range distinctRanges {
		fresh += r.Y - r.X + 1
	}
	fmt.Println(fresh)
}
