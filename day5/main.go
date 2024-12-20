package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type Rule struct {
	first int
	later int
}

type UpdateOrder []int

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse input
	var rules []Rule
	var updateOrders []UpdateOrder
	inputSector := 0
	for _, line := range input {
		if line == "" {
			inputSector++
			continue
		}

		switch inputSector {
		case 0:
			// Parse Rule
			vals := aoc.ParseSeperatedInts(line, "|")
			rules = append(rules, Rule{first: vals[0], later: vals[1]})
		case 1:
			// Parse UpdateOrder
			updateOrders = append(updateOrders, aoc.ParseSeperatedInts(line, ","))
		}
	}

	// Part 1
	sumOfCorrectMiddlePages := 0
	for _, updateOrder := range updateOrders {
		if updateOrder.followsAllRules(rules) {
			sumOfCorrectMiddlePages += updateOrder.getMiddlePage()
		}
	}
	fmt.Println(sumOfCorrectMiddlePages)

	// Part 2
	sumOfFixedMiddlePages := 0
	for _, updateOrder := range updateOrders {
		if !updateOrder.followsAllRules(rules) {
			updateOrder.fixUpdateOrder(rules)
			sumOfFixedMiddlePages += updateOrder.getMiddlePage()
		}
	}
	fmt.Println(sumOfFixedMiddlePages)
}

func (updateOrder UpdateOrder) fixUpdateOrder(rules []Rule) {
	// Loop over all pairs of updates
	for i := range len(updateOrder) {
		for j := i + 1; j < len(updateOrder); j++ {
			// Check if current order of selected 2 updates violate any rule, flip values if so (bubble sort)
			for _, rule := range rules {
				if updateOrder[i] == rule.later && updateOrder[j] == rule.first {
					updateOrder[i], updateOrder[j] = updateOrder[j], updateOrder[i]
				}
			}
		}
	}
}

func (updateOrder UpdateOrder) getMiddlePage() int {
	return updateOrder[len(updateOrder)/2]
}

func (updateOrder UpdateOrder) followsAllRules(rules []Rule) bool {
	// Store update indices
	indexedUpdates := make(map[int]int)
	for i, update := range updateOrder {
		indexedUpdates[update] = i
	}

	// Check if calculated indices follow all rules
	for _, rule := range rules {
		beforeIndex, beforeExists := indexedUpdates[rule.first]
		afterIndex, afterExists := indexedUpdates[rule.later]
		if beforeExists && afterExists && beforeIndex > afterIndex {
			return false
		}
	}
	return true
}
