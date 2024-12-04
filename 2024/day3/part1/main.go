package main

import (
	"fmt"
	"os"
	"regexp"
)

func main() {
	memory, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	r := regexp.MustCompile(`mul\(\d+,\d+\)`)
	matches := r.FindAllString(string(memory), -1)
	sum := 0
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
