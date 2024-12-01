package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

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

	frequency := make(map[int]int, len(right))
	for _, id := range right {
		frequency[id] += 1
	}

	similarityScore := 0

	for _, id := range left {
		similarityScore += id * frequency[id]
	}

	fmt.Println(similarityScore)
}
