package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"strconv"
)

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	mapSize := aoc.Point{Y: len(input), X: len(input[0])}
	heightMap := make(map[aoc.Point]int)
	for y, line := range input {
		// Parse line
		for x, char := range line {
			heightMap[aoc.Point{X: x, Y: y}], _ = strconv.Atoi(string(char))
		}
	}

	// Part 1
	var startingPoints []aoc.Point
	for k, v := range heightMap {
		if v == 0 {
			startingPoints = append(startingPoints, aoc.Point{X: k.X, Y: k.Y})
		}
	}

	totalScore := 0
	for _, startingPoint := range startingPoints {
		score, _ := calculateScore(startingPoint, heightMap, mapSize)
		totalScore += score
	}
	fmt.Println(totalScore)

	// Part 2
	totalRating := 0
	for _, startingPoint := range startingPoints {
		_, rating := calculateScore(startingPoint, heightMap, mapSize)
		totalRating += rating
	}

}

func calculateScore(startingPoint aoc.Point, heightMap map[aoc.Point]int, mapSize aoc.Point) (int, int) {
	queue := aoc.Queue[aoc.Point]{}
	queue.Enqueue(aoc.QueueItem[aoc.Point]{
		Value:  startingPoint,
		Traces: []aoc.Point{startingPoint},
	})

	// Solve BFS
	reachablePeaks := make(map[aoc.Point]bool)
	rating := 0
	for !queue.IsEmpty() {
		point := queue.Dequeue()

		// Check if point is 9 (target)
		if heightMap[point.Value] == 9 {
			reachablePeaks[point.Value] = true
			rating++
			continue
		}

		// Push children
		currentHeight := heightMap[point.Value]
		childDiffs := []aoc.Point{{Y: 1, X: 0}, {Y: -1, X: 0}, {Y: 0, X: 1}, {Y: 0, X: -1}}
		for _, childDiff := range childDiffs {
			newPoint := point.Value.Add(childDiff)

			if !newPoint.OutOfBounds(mapSize) {
				newHeight := heightMap[newPoint]
				if newHeight == currentHeight+1 {
					queue.Enqueue(aoc.QueueItem[aoc.Point]{Value: newPoint, Traces: append(point.Traces, newPoint)})
				}
			}
		}
	}

	return len(reachablePeaks), rating
}
