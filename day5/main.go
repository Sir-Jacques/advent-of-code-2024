package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type Rule struct {
	before int
	after  int
}

type Book struct {
	pages []int
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse input
	var rules []Rule
	var books []Book
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
			beforePage, _ := strconv.Atoi(vals[0])
			afterPage, _ := strconv.Atoi(vals[1])
			rules = append(rules, Rule{before: beforePage, after: afterPage})
		case 1:
			// Parse Book
			stringPages := strings.Split(line, ",")
			pages := make([]int, len(stringPages))
			for i, p := range stringPages {
				pages[i], _ = strconv.Atoi(p)
			}
			books = append(books, Book{pages: pages})
		}
	}

	// Part 1
	sumOfCorrectMiddlePages := 0
	for _, book := range books {
		if book.followsAllRules(rules) {
			sumOfCorrectMiddlePages += book.getMiddlePage()
		}
	}
	fmt.Println(sumOfCorrectMiddlePages)

	// Part 2
	sumOfFixedMiddlePages := 0
	for _, book := range books {
		if !book.followsAllRules(rules) {
			book.fixPages(rules)
			sumOfFixedMiddlePages += book.getMiddlePage()
		}
	}
	fmt.Println(sumOfFixedMiddlePages)
}

func (b *Book) fixPages(rules []Rule) {
	for i := 0; i < len(b.pages); i++ {
		for j := i + 1; j < len(b.pages); j++ {
			// Check if pages violate any rule, flip them if so
			for _, rule := range rules {
				if b.pages[i] == rule.after && b.pages[j] == rule.before {
					b.pages[i], b.pages[j] = b.pages[j], b.pages[i]
				}
			}
		}
	}
}

func (b *Book) getMiddlePage() int {
	return b.pages[len(b.pages)/2]
}

func (b *Book) followsAllRules(rules []Rule) bool {
	indexedPages := make(map[int]int)
	for i, page := range b.pages {
		indexedPages[page] = i
	}
	for _, rule := range rules {
		beforeIndex, beforeExists := indexedPages[rule.before]
		afterIndex, afterExists := indexedPages[rule.after]
		if beforeExists && afterExists && beforeIndex > afterIndex {
			return false
		}
	}
	return true
}
