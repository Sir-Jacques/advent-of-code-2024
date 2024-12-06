package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"sort"
)

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	var list0 []int
	var list1 []int
	for _, i := range input {
		vals := aoc.ParseSeperatedInts(i, "   ")
		list0 = append(list0, vals[0])
		list1 = append(list1, vals[1])
	}

	sort.Ints(list0)
	sort.Ints(list1)

	// Part 1
	difference := 0
	for k := range list0 {
		difference += aoc.Abs(list0[k] - list1[k])
	}
	fmt.Println(difference)

	// Part 2
	similarity := 0
	for _, v := range list0 {
		similarity += v * aoc.CountElementInList(list1, v)
	}
	fmt.Println(similarity)
}
