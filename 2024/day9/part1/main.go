package main

import (
	"fmt"
	"os"
	"slices"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	disk := []int{}
	isFile := true
	id := 1
	var nextFree int
	var lastTaken int
	for _, fileOrFreeSpace := range input {
		amount := int(fileOrFreeSpace - '0')
		if isFile {
			disk = append(disk, slices.Repeat([]int{id}, amount)...)
			id++
			if amount != 0 {
				lastTaken = len(disk) - 1
			}
		} else {
			disk = append(disk, slices.Repeat([]int{0}, amount)...)
			if nextFree == 0 && amount != 0 {
				nextFree = len(disk) - amount
			}
		}
		isFile = !isFile
	}
	for lastTaken != nextFree-1 {
		disk[nextFree], disk[lastTaken] = disk[lastTaken], disk[nextFree]
		for disk[nextFree] != 0 {
			nextFree++
		}
		for disk[lastTaken] == 0 {
			lastTaken--
		}
	}
	checksum := 0
	for i, id := range disk {
		if id == 0 {
			break
		}
		checksum += i * (id - 1)
	}
	fmt.Println(checksum)
}
