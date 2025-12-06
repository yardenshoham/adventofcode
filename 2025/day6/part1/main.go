package main

import (
	"fmt"
	"image"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	grid := map[image.Point]string{}
	var lenX, lenY int
	for y, line := range strings.Split(string(input), "\n") {
		for x, numOrOp := range strings.Fields(line) {
			grid[image.Pt(x, y)] = numOrOp
			lenX = x + 1
		}
		lenY = y + 1
	}
	sum := 0
	for i := range lenX {
		op := grid[image.Pt(i, lenY-1)]
		solution, err := strconv.Atoi(grid[image.Pt(i, 0)])
		if err != nil {
			panic(err)
		}
		for j := 1; j < lenY-1; j++ {
			num, err := strconv.Atoi(grid[image.Pt(i, j)])
			if err != nil {
				panic(err)
			}
			if op == "+" {
				solution += num
			} else {
				solution *= num
			}
		}
		sum += solution
	}
	fmt.Println(sum)
}
