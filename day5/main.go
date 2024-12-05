package main

import (
	"fmt"
	"strconv"
	"strings"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

// A rule describes which page must come before the other
type Rule struct {
	before string
	after  string
}

type Book struct {
	pages        []string
	indexedPages map[string]int
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse input
	var rules []Rule
	var books []Book
	sector := 0
	for _, line := range input {
		// Next sector
		if line == "" {
			sector++
			continue
		}

		switch sector {
		case 0:
			// Parse rule
			vals := strings.Split(line, "|")
			rules = append(rules, Rule{before: vals[0], after: vals[1]})
		case 1:
			pages := strings.Split(line, ",")
			book := Book{pages: pages}
			book.indexPages()
			books = append(books, book)
		default:
			fmt.Println("Invalid sector")
		}
	}

	// Part 1
	sumOfCorrectMiddlePages := 0
	for _, book := range books {
		if book.followsAllRules(rules) {
			middlePageInt, _ := strconv.Atoi(book.getMiddlePage())
			sumOfCorrectMiddlePages += middlePageInt
		}
	}

	// Part 2
	sumOfFixedMiddlePages := 0
	for _, book := range books {
		if !book.followsAllRules(rules) {
			book.fixPages(rules)
			middlePageInt, _ := strconv.Atoi(book.getMiddlePage())
			sumOfFixedMiddlePages += middlePageInt
		}
	}

	fmt.Println(sumOfCorrectMiddlePages)
	fmt.Println(sumOfFixedMiddlePages)
}

func (b *Book) fixPages(rules []Rule) {
	for i := 0; i < len(b.pages); i++ {
		for j := i + 1; j < len(b.pages); j++ {
			// Check if pages violate any rule
			followsRules := true
			for _, rule := range rules {
				if b.pages[i] == rule.after && b.pages[j] == rule.before {
					followsRules = false
				}
			}

			// Flip wrong pages around
			if !followsRules {
				buffer := b.pages[i]
				b.pages[i] = b.pages[j]
				b.pages[j] = buffer
			}
		}
	}
	b.indexPages()
}

func (b *Book) indexPages() {
	b.indexedPages = make(map[string]int)
	for i, page := range b.pages {
		b.indexedPages[page] = i
	}
}

// Part 1
func (b *Book) getMiddlePage() string {
	middleIndex := len(b.pages) / 2
	return b.pages[middleIndex]
}

func (b *Book) followsAllRules(rules []Rule) bool {
	for _, rule := range rules {
		if !b.followsRule(rule) {
			return false
		}
	}
	return true
}

func (b *Book) followsRule(rule Rule) bool {
	beforeIndex, beforeExists := b.indexedPages[rule.before]
	afterIndex, afterExists := b.indexedPages[rule.after]
	if !beforeExists || !afterExists {
		return true
	}
	return beforeIndex < afterIndex
}
