package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func copyMap(m map[string]int) map[string]int {
	r := make(map[string]int, len(m))
	for k, v := range m {
		r[k] = v
	}
	return r
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	count := map[string]int{}
	for _, stone := range strings.Fields(string(input)) {
		count[stone] = 1
	}
	for range 75 {
		newCount := map[string]int{}
		for stone, amount := range count {
			if stone == "0" {
				newCount["1"] += amount
				continue
			}
			if len(stone)%2 == 0 {
				asInt, err := strconv.Atoi(stone[len(stone)/2:])
				if err != nil {
					panic(err)
				}
				newCount[stone[:len(stone)/2]] += amount
				newCount[strconv.Itoa(asInt)] += amount
				continue
			}
			asInt, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			newCount[strconv.Itoa(asInt*2024)] += amount
		}
		count = newCount
	}
	sum := 0
	for _, c := range count {
		sum += c
	}
	fmt.Println(sum)
}
