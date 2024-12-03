package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	moduleDir := filepath.Dir(filename)
	input := readInput(fmt.Sprintf("%s/input.txt", moduleDir))

	// Part 1
	result1 := 0
	for _, line := range input {
		result1 += getMulSummation(line)
	}
	fmt.Println(result1)

	// Part 2
	result2 := 0
	for _, line := range input {
		extendedLine := line + "do()"
		// Bootstrap with expected states and filter out dont()->do()
		pattern := regexp.MustCompile(`don't\(\).*?do\(\)`)
		filteredLine := pattern.ReplaceAllString(extendedLine, "")
		result2 += getMulSummation(filteredLine)
	}
	fmt.Println(result2)
}

func getMulSummation(input string) int {
	result := 0
	// Compile the regex pattern
	pattern := `mul\(\d+,\d+\)`
	re := regexp.MustCompile(pattern)

	// Find all matches
	matches := re.FindAllString(input, -1)

	for _, match := range matches {
		int0, _ := strconv.Atoi(strings.Split(strings.Split(match, "(")[1], ",")[0])
		int1, _ := strconv.Atoi(strings.Split(strings.Split(match, "(")[1], ",")[1][:len(strings.Split(strings.Split(match, "(")[1], ",")[1])-1])
		result += int0 * int1
	}

	return result
}

func readInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(content), "\n")
}
