package main

import (
	"bytes"
	"fmt"
	"os"
)

func possibilities(design []byte, towels map[byte][][]byte, designs map[string]int) int {
	d, ok := designs[string(design)]
	if ok {
		return d
	}
	if len(design) == 0 {
		return 1
	}
	sum := 0
	for _, towel := range towels[design[0]] {
		if bytes.HasPrefix(design, towel) {
			p := possibilities(design[len(towel):], towels, designs)
			sum += p
		}
	}
	designs[string(design)] = sum
	return sum
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
	towels := map[byte][][]byte{}
	for _, towel := range bytes.Split(parts[0], []byte{',', ' '}) {
		towels[towel[0]] = append(towels[towel[0]], towel)
	}
	designs := map[string]int{}
	sum := 0
	for _, design := range bytes.Fields(parts[1]) {
		sum += possibilities(design, towels, designs)
	}
	fmt.Println(sum)
}
