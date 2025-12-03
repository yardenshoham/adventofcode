package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type memoInputs struct {
	n int
	i int
}

func joltage(batteries []byte, n, i int, memo map[memoInputs]int) int {
	if n == 0 {
		return 0
	}
	if i == len(batteries) {
		return math.MinInt
	}
	if num, ok := memo[memoInputs{n, i}]; ok {
		return num
	}
	with := int(batteries[i]-'0')*int(math.Pow10(n-1)) + joltage(batteries, n-1, i+1, memo)
	without := joltage(batteries, n, i+1, memo)
	res := max(with, without)
	memo[memoInputs{n, i}] = res
	return res
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	var joltageSum int
	for scanner.Scan() {
		joltageSum += joltage(scanner.Bytes(), 12, 0, map[memoInputs]int{})
	}
	fmt.Println(joltageSum)
}
