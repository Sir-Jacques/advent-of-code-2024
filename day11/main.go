package main

import (
	"fmt"
	"strconv"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type StoneState struct {
	number, depthRemaining int
}

var memoizationMap map[StoneState]int

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")
	numbers := aoc.ParseSeperatedInts(input[0], " ")

	memoizationMap = make(map[StoneState]int)

	// Part 1
	fmt.Println(calculateTotalStones(numbers, 25))

	// Part 2
	fmt.Println(calculateTotalStones(numbers, 75))
}

func calculateTotalStones(numbers []int, depth int) int {
	totalSize := 0
	for _, n := range numbers {
		totalSize += blinkRecursive(n, depth)
	}
	return totalSize
}

// Returns the amount of stones 'n' has after 'depthRemaining' blinks
func blinkRecursive(n int, depthRemaining int) int {
	// Leaf node, count is 1
	if depthRemaining == 0 {
		return 1
	}

	// Perform memoization lookup
	stoneState := StoneState{number: n, depthRemaining: depthRemaining}
	lookup, exists := memoizationMap[stoneState]
	if exists {
		return lookup
	}

	// Rule 1; 0 becomes 1
	if n == 0 {
		return blinkRecursive(1, depthRemaining-1)
	}

	// Rule 2; even numbers get split into 2 (add their recursive leaf counts)
	if len(strconv.Itoa(n))%2 == 0 {
		int0, int1 := splitInt(n)
		memoizationMap[stoneState] = blinkRecursive(int0, depthRemaining-1) + blinkRecursive(int1, depthRemaining-1)
		return memoizationMap[stoneState]
	}

	// Rule 3; odd numbers get multiplied by 2024
	memoizationMap[stoneState] = blinkRecursive(n*2024, depthRemaining-1)
	return memoizationMap[stoneState]
}

// Split int into 2 ints of equal digits
func splitInt(n int) (int, int) {
	digits := strconv.Itoa(n)
	int0, _ := strconv.Atoi(digits[:len(digits)/2])
	int1, _ := strconv.Atoi(digits[len(digits)/2:])
	return int0, int1
}
