package main

import (
	"bufio"
	"fmt"
	"os"
)

func parseMatch(m []byte) int {
	if len(m) == 1 {
		return int(m[0] - '0')
	}
	switch string(m) {
	case "one":
		return 1
	case "two":
		return 2
	case "three":
		return 3
	case "four":
		return 4
	case "five":
		return 5
	case "six":
		return 6
	case "seven":
		return 7
	case "eight":
		return 8
	case "nine":
		return 9
	}
	return 0
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	sum := 0
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		var first, second int
		for i, r := range line {
			d := 0
			if '0' <= r && r <= '9' {
				d = int(r - '0')
			}
			if i < len(line)-2 {
				if line[i:i+3] == "one" {
					d = 1
				}
				if line[i:i+3] == "two" {
					d = 2
				}
				if line[i:i+3] == "six" {
					d = 6
				}
			}
			if i < len(line)-3 {
				if line[i:i+4] == "four" {
					d = 4
				}
				if line[i:i+4] == "five" {
					d = 5
				}
				if line[i:i+4] == "nine" {
					d = 9
				}
			}
			if i < len(line)-4 {
				if line[i:i+5] == "three" {
					d = 3
				}
				if line[i:i+5] == "seven" {
					d = 7
				}
				if line[i:i+5] == "eight" {
					d = 8
				}
			}
			if first == 0 {
				first = d
			}
			if d != 0 {
				second = d
			}
		}
		sum += first*10 + second
	}
	fmt.Println(sum)
}
