package helpers

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

// ReadInput reads the input file and returns a slice of strings
func ReadInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}

// ParseSeperatedInts parses an input string into a slice of ints
func ParseSeperatedInts(line string, separator string) []int {
	var result []int
	for _, val := range strings.Split(line, separator) {
		num, _ := strconv.Atoi(val)
		result = append(result, num)
	}
	return result
}

func Copy2DSlice[T any](slice [][]T) [][]T {
	result := make([][]T, len(slice))
	for i := range slice {
		result[i] = make([]T, len(slice[i]))
		copy(result[i], slice[i])
	}
	return result
}