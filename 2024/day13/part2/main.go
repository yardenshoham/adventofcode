package main

import (
	"bufio"
	"fmt"
	"image"
	"os"
)

type clawMachine struct {
	A     image.Point
	B     image.Point
	Prize image.Point
}

func minPrice(cm clawMachine) int {
	denominator := cm.A.X*cm.B.Y - cm.B.X*cm.A.Y
	if denominator == 0 {
		return 0
	}
	a := cm.Prize.X*cm.B.Y - cm.Prize.Y*cm.B.X
	b := cm.Prize.Y*cm.A.X - cm.Prize.X*cm.A.Y
	if a%denominator != 0 || b%denominator != 0 {
		return 0
	}
	if a/denominator < 0 || b/denominator < 0 {
		return 0
	}
	return 3*(a/denominator) + b/denominator
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
			clawMachines[len(clawMachines)-1].Prize = image.Pt(x+10000000000000, y+10000000000000)
		}
		i = (i + 1) % 4
	}
	sum := 0
	for _, cm := range clawMachines {
		sum += minPrice(cm)
	}
	fmt.Println(sum)
}
