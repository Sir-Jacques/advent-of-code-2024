package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type Problem struct {
	target  int
	numbers []int
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse input
	var problems = make([]Problem, len(input))
	for i, line := range input {
		parts := strings.Split(line, ": ")
		problems[i].target, _ = strconv.Atoi(parts[0])
		problems[i].numbers = aoc.ParseSeperatedInts(parts[1], " ")
	}

	// Part 1
	sum1 := 0
	for _, problem := range problems {
		solved := problem.solveDFS(false)
		if solved {
			sum1 += problem.target
		}
	}
	fmt.Println(sum1)

	// Part 2
	sum2 := 0
	for _, problem := range problems {
		solved := problem.solveDFS(true)
		if solved {
			sum2 += problem.target
		}
	}
	fmt.Println(sum2)
}

func (p *Problem) solveDFS(includeConcatOperation bool) bool {
	// Empty input
	if len(p.numbers) == 0 {
		return false
	}

	// Single number must match target to solve
	if len(p.numbers) == 1 {
		return p.numbers[0] == p.target
	}

	// Push initial node
	stack := aoc.NewStack[*StackNode]()
	stack.Push(&StackNode{Accumulator: p.numbers[0], RemainingNumbers: p.numbers[1:]})

	// DFS
	for !stack.IsEmpty() {
		// Pop item and check if we're finished
		item := stack.Pop()
		if len(item.RemainingNumbers) == 0 {
			if item.Accumulator == p.target {
				return true
			}
			continue
		}

		// Push child nodes
		stack.Push(item.GetChild(func(a, b int) int { return a + b })) // Addition
		stack.Push(item.GetChild(func(a, b int) int { return a * b })) // Multiplication
		if includeConcatOperation {
			stack.Push(item.GetChild(concatInts)) // Concatenation
		}
	}

	return false
}

func concatInts(int0 int, int1 int) int {
	factor := 1
	for int1 >= factor {
		factor *= 10
	}
	return int0*factor + int1
}

type StackNode struct {
	Accumulator      int
	RemainingNumbers []int
}

func (sn *StackNode) GetChild(operation func(int, int) int) *StackNode {
	return &StackNode{
		Accumulator:      operation(sn.Accumulator, sn.RemainingNumbers[0]),
		RemainingNumbers: sn.RemainingNumbers[1:],
	}
}
