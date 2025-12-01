package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// https://stackoverflow.com/a/59299881
func mod(a, b int) int {
	return (a%b + b) % b
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	at := 50
	password := 0
	for scanner.Scan() {
		line := scanner.Text()
		sign := 1
		if line[0] == 'L' {
			sign = -1
		}
		num, err := strconv.Atoi(line[1:])
		if err != nil {
			panic(err)
		}
		for range num {
			at = mod(at+sign, 100)
			if at == 0 {
				password++
			}
		}

	}
	fmt.Println(password)
}
