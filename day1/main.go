package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"sort"
	"strconv"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	moduleDir := filepath.Dir(filename)
	input := readInput(fmt.Sprintf("%s/input.txt", moduleDir))

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

	solution := 0
	similarity := 0
	for k, v := range slice1 {
		solution += abs(slice2[k] - slice1[k])
		similarity += v * countOccurrences(slice2, v)
	}

	fmt.Println(solution)   // Part 1
	fmt.Println(similarity) // Part 2
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
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(content), "\n")
}
