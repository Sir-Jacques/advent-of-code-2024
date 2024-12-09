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

type Node struct {
	accumulator      int
	remainingNumbers []int
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
	stack := aoc.NewStack[*Node]()
	stack.Push(&Node{accumulator: p.numbers[0], remainingNumbers: p.numbers[1:]})

	// DFS
	for !stack.IsEmpty() {
		// Pop item and check if we're finished
		item := stack.Pop()
		if len(item.remainingNumbers) == 0 {
			if item.accumulator == p.target {
				return true
			}
			continue
		}

		// Push child nodes
		stack.Push(item.getChild(func(a, b int) int { return a + b })) // Addition
		stack.Push(item.getChild(func(a, b int) int { return a * b })) // Multiplication
		if includeConcatOperation {
			stack.Push(item.getChild(concatInts)) // Concatenation
		}
	}

	return false
}

func (qi *Node) getChild(operation func(int, int) int) *Node {
	return &Node{
		accumulator:      operation(qi.accumulator, qi.remainingNumbers[0]),
		remainingNumbers: qi.remainingNumbers[1:]}
}

func concatInts(int0 int, int1 int) int {
	factor := 1
	for int1 >= factor {
		factor *= 10
	}
	return int0*factor + int1
}
