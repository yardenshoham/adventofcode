package main

import (
	"bytes"
	"fmt"
	"os"
	"slices"
	"strconv"
)

func isCorrect(update []byte, pageOrderingRules map[string][][]byte) bool {
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
				return false
			}
		}
		currentPages[string(p)] = struct{}{}
	}
	return true
}

func dfs(v string, visited map[string]struct{}, stack *[]string, relevantPageOrderingRules map[string][][]byte) {
	if _, ok := visited[v]; ok {
		return
	}
	visited[v] = struct{}{}
	for _, u := range relevantPageOrderingRules[v] {
		dfs(string(u), visited, stack, relevantPageOrderingRules)
	}
	*stack = append(*stack, v)
}

func topologicalOrderMiddle(update []byte, pageOrderingRules map[string][][]byte) int {
	pages := bytes.Split(update, []byte{','})
	allPages := make(map[string]struct{}, len(pages))
	for _, p := range pages {
		allPages[string(p)] = struct{}{}
	}
	relevantPageOrderingRules := make(map[string][][]byte)
	ordered := make([]string, 0, len(relevantPageOrderingRules))
	for after, befores := range pageOrderingRules {
		if _, ok := allPages[after]; ok == false {
			continue
		}
		for _, b := range befores {
			if _, ok := allPages[string(b)]; ok == false {
				continue
			}
			relevantPageOrderingRules[after] = append(relevantPageOrderingRules[after], b)
		}
		ordered = append(ordered, after)
	}
	slices.SortFunc(ordered, func(a, b string) int { return len(relevantPageOrderingRules[a]) - len(relevantPageOrderingRules[b]) })
	stack := make([]string, 0, len(ordered))
	visited := make(map[string]struct{}, len(ordered))
	for _, v := range ordered {
		dfs(v, visited, &stack, relevantPageOrderingRules)
	}
	num, err := strconv.Atoi(string(stack[len(stack)/2]))
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
		if !isCorrect(update, pageOrderingRules) {
			sum += topologicalOrderMiddle(update, pageOrderingRules)
		}
	}
	fmt.Println(sum)
}
