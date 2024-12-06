package main

import (
	"fmt"
	"regexp"
	"strings"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")
	fullInput := strings.Join(input, "")

	// Part 1
	result1 := getMulSummation(fullInput)
	fmt.Println(result1)

	// Part 2
	filteredInput := regexp.MustCompile(`don't\(\).*?do\(\)`).ReplaceAllString(fullInput+"do()", "")
	result2 := getMulSummation(filteredInput)
	fmt.Println(result2)
}

// getMulSummation parses and returns result of mul() instruction in input string
func getMulSummation(input string) int {
	result := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	for _, match := range re.FindAllString(input, -1) {
		mulNums := aoc.ParseSeperatedInts(match[4:len(match)-1], ",")
		result += mulNums[0] * mulNums[1]
	}
	return result
}
