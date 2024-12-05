package main

import (
	"bytes"
	"fmt"
	"os"
	"strconv"
)

func middle(update []byte, pageOrderingRules map[string][][]byte) int {
	pages := bytes.Split(update, []byte{','})
	allPages := make(map[string]struct{}, len(pages))
	for _, p := range pages {
		allPages[string(p)] = struct{}{}
	}
	currentPages := make(map[string]struct{}, len(pages))
	for _, p := range pages {
		rules := pageOrderingRules[string(p)]
		for _, r := range rules {
			_, hasThePage := allPages[string(r)]
			_, readAlready := currentPages[string(r)]
			if hasThePage && !readAlready {
				return 0
			}
		}
		currentPages[string(p)] = struct{}{}
	}
	num, err := strconv.Atoi(string(pages[len(pages)/2]))
	if err != nil {
		panic(err)
	}
	return num
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	parts := bytes.Split(input, []byte{'\n', '\n'})
	if len(parts) != 2 {
		panic("bad input: " + string(input))
	}
	pageOrderingRules := make(map[string][][]byte)
	for _, por := range bytes.Fields(parts[0]) {
		beforeAfter := bytes.Split(por, []byte{'|'})
		after := string(beforeAfter[1])
		pageOrderingRules[after] = append(pageOrderingRules[after], beforeAfter[0])
	}

	sum := 0
	for _, update := range bytes.Fields(parts[1]) {
		sum += middle(update, pageOrderingRules)
	}
	fmt.Println(sum)
}
