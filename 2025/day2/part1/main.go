package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isInvalidID(n int) bool {
	asString := strconv.Itoa(n)
	if len(asString)%2 == 1 {
		return false
	}
	return asString[:len(asString)/2] == asString[len(asString)/2:]
}

func main() {
	line, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	invalidIDSum := 0
	for _, r := range strings.Split(string(line), ",") {
		parts := strings.Split(r, "-")
		from, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		to, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		for i := from; i < to+1; i++ {
			if isInvalidID(i) {
				invalidIDSum += i
			}
		}
	}
	fmt.Println(invalidIDSum)
}
