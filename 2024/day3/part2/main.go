package main

import (
	"bytes"
	"fmt"
	"os"
	"regexp"
)

func main() {
	memory, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}

	whitespace := regexp.MustCompile(`\s+`)
	memory = whitespace.ReplaceAll(memory, []byte{})

	dontUntilDo := regexp.MustCompile(`(?U)don't\(\).*do\(\)`)
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)

	sum := 0

	memoryWithoutDontDos := dontUntilDo.ReplaceAll(memory, []byte{})
	memoryWithoutDonts, _, _ := bytes.Cut(memoryWithoutDontDos, []byte{'d', 'o', 'n', '\'', 't', '(', ')'})
	matches := r.FindAllString(string(memoryWithoutDonts), -1)
	for _, match := range matches {
		var a, b int
		_, err = fmt.Sscanf(match, "mul(%d,%d)", &a, &b)
		if err != nil {
			panic(err)
		}
		sum += a * b

	}

	fmt.Println(sum)
}
