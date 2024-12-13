package main

import (
	"bufio"
	"fmt"
	"image"
	"math"
	"os"
)

type clawMachine struct {
	A     image.Point
	B     image.Point
	Prize image.Point
}

func minPrice(cm clawMachine) int {
	result := math.MaxInt64
	for A := range 101 {
		for B := range 101 {
			if cm.A.Mul(A).Add(cm.B.Mul(B)).Eq(cm.Prize) {
				result = min(result, 3*A+B)
			}
		}
	}
	if result == math.MaxInt64 {
		result = 0
	}
	return result
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	clawMachines := []clawMachine{}
	i := 0
	for scanner.Scan() {
		var x, y int
		line := scanner.Text()
		switch i % 4 {
		case 0:
			fmt.Sscanf(line, "Button A: X+%d, Y+%d", &x, &y)
			clawMachines = append(clawMachines, clawMachine{A: image.Pt(x, y)})
		case 1:
			fmt.Sscanf(line, "Button B: X+%d, Y+%d", &x, &y)
			clawMachines[len(clawMachines)-1].B = image.Pt(x, y)
		case 2:
			fmt.Sscanf(line, "Prize: X=%d, Y=%d", &x, &y)
			clawMachines[len(clawMachines)-1].Prize = image.Pt(x, y)
		}
		i = (i + 1) % 4
	}
	sum := 0
	for _, cm := range clawMachines {
		sum += minPrice(cm)
	}
	fmt.Println(sum)
}
