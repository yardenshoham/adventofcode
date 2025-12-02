package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func isRepeating(asString string, i int) bool {
	if len(asString)%i != 0 {
		return false
	}
	firstPart := asString[:i]
	for j := i; j < len(asString); j += i {
		if asString[j:j+i] != firstPart {
			return false
		}
	}
	return true
}

func isInvalidID(n int) bool {
	asString := strconv.Itoa(n)
	for i := 1; i < len(asString); i++ {
		if isRepeating(asString, i) {
			return true
		}
	}
	return false
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
