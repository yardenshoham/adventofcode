package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	file, err := os.Open("input.txt")
	check(err)
	defer file.Close()

	sum := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		if len(line) == 0 {
			panic("empty line")
		}
		if line[0] != '+' && line[0] != '-' {
			panic("unknown op: " + line)
		}
		num, err := strconv.Atoi(line)
		check(err)

		sum += num
	}

	fmt.Println(sum)
}
