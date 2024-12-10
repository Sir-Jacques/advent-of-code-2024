package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"strconv"
)

type StackNode struct {
	Value  aoc.Point
	Traces []aoc.Point
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")
	mapSize := aoc.Point{Y: len(input), X: len(input[0])}

	// Parse input
	var startingPoints []aoc.Point
	heightMap := make(map[aoc.Point]int)
	for y, line := range input {
		for x, char := range line {
			point := aoc.Point{X: x, Y: y}
			heightMap[point], _ = strconv.Atoi(string(char)) // Store height
			if char == '0' {
				startingPoints = append(startingPoints, point) // Store starting points
			}
		}
	}

	// Part 1
	totalScore := 0
	for _, startingPoint := range startingPoints {
		score, _ := solveTrails(startingPoint, heightMap, mapSize)
		totalScore += score
	}
	fmt.Println(totalScore)

	// Part 2
	totalRating := 0
	for _, startingPoint := range startingPoints {
		_, rating := solveTrails(startingPoint, heightMap, mapSize)
		totalRating += rating
	}
	fmt.Println(totalRating)
}

func solveTrails(startingPoint aoc.Point, heightMap map[aoc.Point]int, mapSize aoc.Point) (int, int) {
	reachablePeaks := make(map[aoc.Point]bool)
	rating := 0

	// Solve DFS
	stack := aoc.NewStack[StackNode]()
	stack.Push(StackNode{Value: startingPoint, Traces: []aoc.Point{startingPoint}})
	for !stack.IsEmpty() {
		point := stack.Pop()

		// Check if point is 9 (target)
		if heightMap[point.Value] == 9 {
			reachablePeaks[point.Value] = true
			rating++
			continue
		}

		// Push children, only if inside bounds and has height+1
		currentHeight := heightMap[point.Value]
		for _, childDiff := range []aoc.Point{{Y: 1, X: 0}, {Y: -1, X: 0}, {Y: 0, X: 1}, {Y: 0, X: -1}} {
			newPoint := point.Value.Add(childDiff)
			if !newPoint.OutOfBounds(mapSize) && heightMap[newPoint] == currentHeight+1 {
				stack.Push(StackNode{Value: newPoint, Traces: append(point.Traces, newPoint)})
			}
		}
	}

	return len(reachablePeaks), rating
}
