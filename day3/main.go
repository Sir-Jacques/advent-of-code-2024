package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Part 1
	result1 := 0
	for _, line := range input {
		result1 += getMulSummation(line)
	}
	fmt.Println(result1)

	// Part 2
	result2 := 0
	for _, line := range input {
		filteredLine := regexp.MustCompile(`don't\(\).*?do\(\)`).ReplaceAllString(line+"do()", "")
		result2 += getMulSummation(filteredLine)
	}
	fmt.Println(result2)
}

func getMulSummation(input string) int {
	result := 0
	re := regexp.MustCompile(`mul\(\d+,\d+\)`)
	for _, match := range re.FindAllString(input, -1) {
		nums := strings.Split(match[4:len(match)-1], ",")
		int0, _ := strconv.Atoi(nums[0])
		int1, _ := strconv.Atoi(nums[1])
		result += int0 * int1
	}

	return result
}
