package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	stones := strings.Fields(string(input))
	for range 25 {
		tmpStones := []string{}
		for _, stone := range stones {
			if stone == "0" {
				tmpStones = append(tmpStones, "1")
				continue
			}
			if len(stone)%2 == 0 {
				asInt, err := strconv.Atoi(stone[len(stone)/2:])
				if err != nil {
					panic(err)
				}
				tmpStones = append(tmpStones, stone[:len(stone)/2], strconv.Itoa(asInt))
				continue
			}
			asInt, err := strconv.Atoi(stone)
			if err != nil {
				panic(err)
			}
			tmpStones = append(tmpStones, strconv.Itoa(asInt*2024))
		}
		stones = tmpStones
	}
	fmt.Println(len(stones))
}
