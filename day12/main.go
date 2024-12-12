package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"slices"
)

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")
	bounds := aoc.Point{X: len(input[0]), Y: len(input)}
	garden, gardenSets := make(map[aoc.Point]rune), make(map[rune][]aoc.Point)
	for y, line := range input {
		for x, char := range line {
			garden[aoc.Point{X: x, Y: y}] = char
			gardenSets[char] = append(gardenSets[char], aoc.Point{X: x, Y: y})
		}
	}

	// Separate groups that do not connect
	var gardenGroup [][]aoc.Point
	for _, set := range gardenSets {
		for _, group := range separateGroups(set, bounds) {
			gardenGroup = append(gardenGroup, group)
		}
	}

	// Part 1
	price := 0
	for _, group := range gardenGroup {
		price += calculatePerimeter(group, bounds) * len(group)
	}
	fmt.Println(price)

	// Part 2
	price = 0
	for _, group := range gardenGroup {
		price += calculateCorners(group, bounds) * len(group)
	}
	fmt.Println(price)
}

// Split groups that do not connect into separate groups
func separateGroups(group []aoc.Point, mapSize aoc.Point) [][]aoc.Point {
	processed := make(map[aoc.Point]bool)
	var result [][]aoc.Point

	for len(group) > 0 {
		queue := aoc.NewQueue[aoc.Point]()
		queue.Enqueue(group[0])
		var currentGroup []aoc.Point

		// BFS
		for !queue.IsEmpty() {
			point := queue.Dequeue()
			if processed[point] {
				continue
			}
			processed[point] = true
			currentGroup = append(currentGroup, point)

			// Check if direct neighbors are part of this group
			for _, diff := range []aoc.Point{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}} {
				neighbor := point.Add(diff)
				if !neighbor.OutOfBounds(mapSize) && slices.Contains(group, neighbor) {
					queue.Enqueue(neighbor)
				}
			}
		}

		// Remove processed items from group
		filtered := make([]aoc.Point, 0)
		for _, point := range group {
			if !processed[point] {
				filtered = append(filtered, point)
			}
		}
		group = filtered

		// Extend result with found groups
		result = append(result, currentGroup)
	}
	return result
}

// Returns the number of grid sides that touch another group
func calculatePerimeter(points []aoc.Point, mapSize aoc.Point) int {
	// Check all direct neighbours
	perimeter := 0
	for _, point := range points {
		for _, diff := range []aoc.Point{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}} {
			neighbor := point.Add(diff)
			if hasFence(neighbor, points, mapSize) {
				perimeter++
			}
		}
	}
	return perimeter
}

func calculateCorners(points []aoc.Point, mapSize aoc.Point) int {
	convexCorners, concaveCorners := 0, 0
	for _, point := range points {
		fences := 0
		// Horizontal
		hor, ver := false, false
		for _, diff := range []aoc.Point{{X: 1, Y: 0}, {X: -1, Y: 0}, {X: 0, Y: 1}, {X: 0, Y: -1}} {
			if hasFence(point.Add(diff), points, mapSize) {
				fences++
				if diff.X != 0 {
					hor = true
				} else {
					ver = true
				}
			}
		}
		if fences == 4 { // Singleton group, 4 corners
			convexCorners += 4
			continue
		}
		if hor && ver { // Group with both horizontal and vertical fences must be a corner
			convexCorners += fences - 1
		}

		// Check all possible concave corners
		if checkConcaveCorner(point, 1, 1, points) {
			concaveCorners++
		}
		if checkConcaveCorner(point, -1, 1, points) {
			concaveCorners++
		}
		if checkConcaveCorner(point, 1, -1, points) {
			concaveCorners++
		}
		if checkConcaveCorner(point, -1, -1, points) {
			concaveCorners++
		}
	}
	return convexCorners + concaveCorners
}

// Check if the current point is a valid concave corner
func checkConcaveCorner(point aoc.Point, diffX int, diffY int, points []aoc.Point) bool {
	return slices.Contains(points, point.Add(aoc.Point{X: diffX, Y: 0})) && // Horizontal neighbor must be part of this group
		slices.Contains(points, point.Add(aoc.Point{X: 0, Y: diffY})) && // Vertical neighbor must be part of this group
		!slices.Contains(points, point.Add(aoc.Point{X: diffX, Y: diffY})) // Diagonal neighbor must NOT be part of this group
}

func hasFence(neighbor aoc.Point, currentGroup []aoc.Point, mapSize aoc.Point) bool {
	// Edge of map
	if neighbor.OutOfBounds(mapSize) {
		return true
	}

	// Return if neighbor is part of current group (if so is should be fenced)
	return !slices.Contains(currentGroup, neighbor)
}
