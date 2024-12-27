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
	current := "AAA"
	steps := 0
	for {
		for _, step := range parts[0] {
			if current == "ZZZ" {
				fmt.Println(steps)
				return
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
