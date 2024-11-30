package main

import (
	"bufio"
	"bytes"
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
	fileBytes, err := os.ReadFile("input.txt")
	check(err)

	freq := map[int]struct{}{0: {}}
	sum := 0

	for {
		scanner := bufio.NewScanner(bytes.NewReader(fileBytes))
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
			_, ok := freq[sum]
			if ok {
				goto done
			}
			freq[sum] = struct{}{}
		}
	}

done:
	fmt.Println(sum)
}
