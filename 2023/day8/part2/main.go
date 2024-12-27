package main

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

type node struct {
	Left  string
	Right string
}

func repetition(current string, nodes map[string]node, instructions []byte) int {
	steps := 0
	for {
		for _, step := range instructions {
			if current[2] == 'Z' {
				return steps
			}
			if step == 'L' {
				current = nodes[current].Left
			} else {
				current = nodes[current].Right
			}
			steps++
		}
	}
}

func gcd(a, b int) int {
	if b == 0 {
		return a
	}
	return gcd(b, a%b)
}

func lcm(a, b int) int {
	return (a * b) / gcd(a, b)
}

func lcmNumbers(numbers []int) int {
	if len(numbers) == 2 {
		return lcm(numbers[0], numbers[1])
	}
	return lcmNumbers(append([]int{lcm(numbers[0], numbers[1])}, numbers[2:]...))
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := bytes.Split(input, []byte{'\n', '\n'})
	if len(parts) != 2 {
		panic("bad input")
	}
	nodes := map[string]node{}
	scanner := bufio.NewScanner(bytes.NewBuffer(parts[1]))
	for scanner.Scan() {
		line := scanner.Text()
		nodes[line[:3]] = node{line[7:10], line[12:15]}
	}
	numbers := []int{}
	for n := range nodes {
		if n[2] == 'A' {
			numbers = append(numbers, repetition(n, nodes, parts[0]))
		}
	}
	fmt.Println(lcmNumbers(numbers))
}
