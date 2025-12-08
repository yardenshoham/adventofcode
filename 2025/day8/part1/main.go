package main

import (
	"bufio"
	"cmp"
	"fmt"
	"math"
	"os"
	"slices"
)

type point struct {
	X int
	Y int
	Z int
}

type distanceBetweenPoints struct {
	Distance float64
	P1       point
	P2       point
}

func distance(p1, p2 point) float64 {
	return math.Sqrt(math.Pow(float64(p1.X-p2.X), 2) + math.Pow(float64(p1.Y-p2.Y), 2) + math.Pow(float64(p1.Z-p2.Z), 2))
}

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	points := []point{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var p point
		_, err := fmt.Sscanf(scanner.Text(), "%d,%d,%d", &p.X, &p.Y, &p.Z)
		if err != nil {
			panic(err)
		}
		points = append(points, p)
	}
	distances := []distanceBetweenPoints{}
	for i, p := range points {
		for j := i + 1; j < len(points); j++ {
			distances = append(distances, distanceBetweenPoints{distance(p, points[j]), p, points[j]})
		}
	}
	slices.SortFunc(distances, func(d1, d2 distanceBetweenPoints) int { return cmp.Compare(d1.Distance, d2.Distance) })
	circuits := map[int]map[point]struct{}{}
	junctionBoxes := map[point]int{}
	nextCircuit := 0
	for i := range 1000 {
		circuitIndex1, ok1 := junctionBoxes[distances[i].P1]
		circuitIndex2, ok2 := junctionBoxes[distances[i].P2]
		if ok1 && ok2 && circuitIndex1 == circuitIndex2 {
			continue
		}
		if ok1 && !ok2 {
			junctionBoxes[distances[i].P2] = circuitIndex1
			circuits[circuitIndex1][distances[i].P2] = struct{}{}
			continue
		}
		if ok2 && !ok1 {
			junctionBoxes[distances[i].P1] = circuitIndex2
			circuits[circuitIndex2][distances[i].P1] = struct{}{}
			continue
		}
		if !ok1 && !ok2 {
			circuitIndex := nextCircuit
			nextCircuit++
			junctionBoxes[distances[i].P1] = circuitIndex
			junctionBoxes[distances[i].P2] = circuitIndex
			circuits[circuitIndex] = map[point]struct{}{
				distances[i].P1: {},
				distances[i].P2: {},
			}
			continue
		}
		for junctionBox := range circuits[circuitIndex2] {
			junctionBoxes[junctionBox] = circuitIndex1
			circuits[circuitIndex1][junctionBox] = struct{}{}
		}
		delete(circuits, circuitIndex2)
	}
	results := []int{}
	for _, points := range circuits {
		results = append(results, len(points))
	}
	slices.Sort(results)
	fmt.Println(results[len(results)-1] * results[len(results)-2] * results[len(results)-3])
}
