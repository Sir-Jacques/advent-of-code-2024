package main

import (
	"fmt"
	"strconv"
	"strings"

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
			vals := strings.Split(line, "|")
			firstUpdate, _ := strconv.Atoi(vals[0])
			laterUpdate, _ := strconv.Atoi(vals[1])
			rules = append(rules, Rule{first: firstUpdate, later: laterUpdate})
		case 1:
			// Parse UpdateOrder
			updateList := strings.Split(line, ",")
			updateOrder := make([]int, len(updateList))
			for i, p := range updateList {
				updateOrder[i], _ = strconv.Atoi(p)
			}
			updateOrders = append(updateOrders, updateOrder)
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
	for i := 0; i < len(updateOrder); i++ {
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
