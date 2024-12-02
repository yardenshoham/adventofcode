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

func dropIndex(nums []int, i int) []int {
	res := make([]int, 0, len(nums)-1)
	res = append(res, nums[:i]...)
	res = append(res, nums[i+1:]...)
	return res
}

func isReportSafe(levels []int) bool {
	decreasing := false
	if levels[0] > levels[1] {
		decreasing = true
	}
	for i := 0; i < len(levels)-1; i++ {
		diff := levels[i] - levels[i+1]
		if diff == 0 || (decreasing && (diff > 3 || diff < 1)) || (!decreasing && (diff < -3 || diff > -1)) {
			return false
		}
	}
	return true
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	reports := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		levels := splitInts(line)
		if isReportSafe(levels) {
			reports++
			continue
		}
		for i := range levels {
			if isReportSafe(dropIndex(levels, i)) {
				reports++
				break
			}
		}
	}

	fmt.Println(reports)
}
