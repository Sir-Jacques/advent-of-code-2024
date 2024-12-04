package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

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
		difference += aoc.Abs(slice2[k] - slice1[k])
	}
	fmt.Println(difference)

	// Part 2
	similarity := 0
	for _, v := range slice1 {
		similarity += v * aoc.CountOccurrencesInList(slice2, v)
	}
	fmt.Println(similarity)
}
