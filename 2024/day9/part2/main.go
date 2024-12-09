package main

import (
	"fmt"
	"os"
	"slices"
)

func swap(s []int, from1, from2, n int) {
	for p := 0; p < n; p++ {
		s[from1+p], s[from2+p] = s[from2+p], s[from1+p]
	}
}

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		panic(err)
	}
	disk := []int{}
	files := map[int]int{}
	isFile := true
	id := 1
	var lastTaken int
	for _, fileOrFreeSpace := range input {
		amount := int(fileOrFreeSpace - '0')
		if isFile {
			disk = append(disk, slices.Repeat([]int{id}, amount)...)
			files[id] = amount
			id++
			if amount != 0 {
				lastTaken = len(disk) - 1
			}
		} else {
			disk = append(disk, slices.Repeat([]int{0}, amount)...)
		}
		isFile = !isFile
	}
	for range files {
		id := disk[lastTaken]
		for i := 0; i < lastTaken; i++ {
			if disk[i] == 0 {
				free := 0
				for j := i; j < len(disk); j++ {
					if disk[j] == 0 {
						free++
					} else {
						break
					}
				}
				if free >= files[id] {
					swap(disk, i, lastTaken-files[id]+1, files[id])
					i += free - 1
					break
				}
			}
		}
		lastTaken -= files[id]
		for disk[lastTaken] == 0 {
			lastTaken--
		}
	}
	checksum := 0
	for i, id := range disk {
		if id == 0 {
			continue
		}
		checksum += i * (id - 1)
	}
	fmt.Println(checksum)
}
