package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	input := readInput(filepath.Join(filepath.Dir(filename), "input.txt"))

	var slice1 []int
	var slice2 []int
	for _, i := range input {
		words := strings.Split(i, "   ")
		num1, _ := strconv.Atoi(words[0])
		num2, _ := strconv.Atoi(words[1])
		slice1 = append(slice1, num1)
		slice2 = append(slice2, num2)
	}

	sort.Ints(slice1)
	sort.Ints(slice2)

	// Part 1
	difference := 0
	for k, _ := range slice1 {
		difference += abs(slice2[k] - slice1[k])
	}
	fmt.Println(difference)

	// Part 2
	similarity := 0
	for _, v := range slice1 {
		similarity += v * countOccurrences(slice2, v)
	}
	fmt.Println(similarity)
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func countOccurrences(slice []int, element int) int {
	count := 0
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
