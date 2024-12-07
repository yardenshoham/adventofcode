package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func splitInts(b string) []int {
	parts := strings.Fields(b)
	res := make([]int, 0, len(parts))
	for _, part := range parts {
		num, err := strconv.Atoi(part)
		if err != nil {
			panic(err)
		}
		res = append(res, num)
	}
	return res
}

func isPossible(target int, numbers []int) bool {
	if len(numbers) == 1 {
		return target == numbers[0]
	}
	rest := []int{}
	if len(numbers) > 2 {
		rest = append(rest, numbers[2:]...)
	}
	add := append([]int{numbers[0] + numbers[1]}, rest...)
	mul := append([]int{numbers[0] * numbers[1]}, rest...)
	catValue, err := strconv.Atoi(strconv.Itoa(numbers[0]) + strconv.Itoa(numbers[1]))
	if err != nil {
		panic(err)
	}
	cat := append([]int{catValue}, rest...)
	return isPossible(target, add) || isPossible(target, mul) || isPossible(target, cat)
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	sum := 0
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), ":")
		target, err := strconv.Atoi(parts[0])
		if err != nil {
			panic(err)
		}
		numbers := splitInts(parts[1])
		if isPossible(target, numbers) {
			sum += target
		}
	}
	fmt.Println(sum)
}
