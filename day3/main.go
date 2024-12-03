package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	input := readInput(filepath.Join(filepath.Dir(filename), "input.txt"))

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

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
