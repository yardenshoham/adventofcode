package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var first, second *int
		for _, b := range scanner.Bytes() {
			if '0' <= b && b <= '9' {
				d := int(b - '0')
				if first == nil {
					first = &d
					continue
				}
				second = &d
			}
		}
		if second == nil {
			second = first
		}
		sum += (*first)*10 + (*second)
	}
	fmt.Println(sum)
}
