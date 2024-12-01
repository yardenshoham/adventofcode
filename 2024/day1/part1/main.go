package main

import (
	"bufio"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func calculateDistance(x, y int) int {
	diff := x - y
	if diff < 0 {
		diff = -diff
	}
	return diff
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	left := []int{}
	right := []int{}

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			panic("empty line")
		}

		parts := strings.Split(line, "   ")
		if len(parts) != 2 {
			panic("must have exactly 2 parts: " + string(line))
		}

		leftPart, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		rightPart, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		left = append(left, leftPart)
		right = append(right, rightPart)
	}

	slices.Sort(left)
	slices.Sort(right)

	distance := 0

	for i := range left {
		distance += calculateDistance(left[i], right[i])
	}

	fmt.Println(distance)
}
