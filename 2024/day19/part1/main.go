package main

import (
	"bytes"
	"fmt"
	"os"
)

func possible(design []byte, towels map[byte][][]byte, designs map[string]bool) bool {
	d, ok := designs[string(design)]
	if ok {
		return d
	}
	if len(design) == 0 {
		return true
	}
	for _, towel := range towels[design[0]] {
		if bytes.HasPrefix(design, towel) {
			p := possible(design[len(towel):], towels, designs)
			designs[string(design)] = p
			if p {
				return true
			}
		}
	}
	return false
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
	designs := map[string]bool{}
	sum := 0
	for _, design := range bytes.Fields(parts[1]) {
		if possible(design, towels, designs) {
			sum++
		}
	}
	fmt.Println(sum)
}
