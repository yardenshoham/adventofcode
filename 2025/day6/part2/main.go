package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

func solve(grid map[image.Point]byte, opX, lenY int) int {
	op := grid[image.Pt(opX, lenY-1)]
	x := opX
	solution := 0
	for {
		var num int
		for y := 0; y < lenY-1; y++ {
			digitByte, ok := grid[image.Pt(x, y)]
			if !ok {
				return solution
			}
			if digitByte != ' ' {
				num = num*10 + int(digitByte-'0')
			}
		}
		if num == 0 {
			break
		}
		if solution == 0 {
			solution = num
		} else if op == '+' {
			solution += num
		} else {
			solution *= num
		}
		x++
	}
	return solution
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	grid := map[image.Point]byte{}
	y := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		for x, numOrOp := range scanner.Bytes() {
			grid[image.Pt(x, y)] = numOrOp
		}
		y++
	}
	x := 0
	sum := 0
	for {
		curr, ok := grid[image.Pt(x, y-1)]
		if !ok {
			break
		}
		if curr == '+' || curr == '*' {
			sum += solve(grid, x, y)
		}
		x++
	}
	fmt.Println(sum)
}
