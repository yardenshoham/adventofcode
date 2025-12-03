package main

import (
	"bufio"
	"fmt"
	"os"
)

func maxInSlice(s []byte, from, to int) (byte, int) {
	var m byte
	var ind int
	for i := from; i < to; i++ {
		if s[i] > m {
			m = s[i]
			ind = i
		}
	}
	return m, ind
}

func joltage(batteries []byte) int {
	maxBattery, maxBatteryIndex := maxInSlice(batteries, 0, len(batteries))
	maxBefore, _ := maxInSlice(batteries, 0, maxBatteryIndex)
	maxAfter, _ := maxInSlice(batteries, maxBatteryIndex+1, len(batteries))
	before := (maxBefore-'0')*10 + (maxBattery - '0')
	after := (maxBattery-'0')*10 + (maxAfter - '0')
	return int(max(before, after))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	joltageSum := 0
	for scanner.Scan() {
		joltageSum += joltage(scanner.Bytes())
	}
	fmt.Println(joltageSum)
}
