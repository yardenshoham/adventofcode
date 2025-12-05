package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func isFresh(ranges map[image.Point]struct{}, ingredientID int) bool {
	for r := range ranges {
		if r.X <= ingredientID && ingredientID <= r.Y {
			return true
		}
	}
	return false
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := strings.Split(string(input), "\n\n")
	ranges := map[image.Point]struct{}{}
	for _, r := range strings.Fields(parts[0]) {
		var from, to int
		_, err := fmt.Sscanf(r, "%d-%d", &from, &to)
		if err != nil {
			panic(err)
		}
		ranges[image.Pt(from, to)] = struct{}{}
	}
	fresh := 0
	for _, ingredient := range strings.Fields(parts[1]) {
		ingredientID, err := strconv.Atoi(ingredient)
		if err != nil {
			panic(err)
		}
		if isFresh(ranges, ingredientID) {
			fresh++
		}
	}
	fmt.Println(fresh)
}
