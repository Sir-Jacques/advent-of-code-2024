package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type CrossWord [][]uint8

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	crossWord := newCrossWord(input)

	// Part 1
	xmasCount := 0
	for y := range len(crossWord) {
		for x := range len(crossWord[y]) {
			xmasCount += crossWord.matchSubstringAllDirs(x, y, "XMAS")
		}
	}

	// Part 2
	crossMasCount := 0
	for y := range len(crossWord) {
		for x := range len(crossWord[y]) {
			if crossWord.posIsCrossMas(x, y) {
				crossMasCount++
			}
		}
	}

	fmt.Println(xmasCount)
	fmt.Println(crossMasCount)
}

func (cw CrossWord) matchSubstringAllDirs(xPos int, yPos int, word string) int {
	count := 0
	for xDiff := range []int{-1, 0, 1} {
		for yDiff := range []int{-1, 0, 1} {
			if xDiff == 0 && yDiff == 0 {
				continue
			}
			if cw.matchSubstringsForPosition(xPos, yPos, xDiff, yDiff, word) {
				count++
			}
		}
	}
	return count
}

// Current char must the be A, looking for crossing MAS in diagonal directions
func (cw CrossWord) posIsCrossMas(xPos int, yPos int) bool {
	if xPos < 1 || yPos < 1 || xPos >= len(cw[0])-1 || yPos >= len(cw)-1 {
		return false // Search area out of bounds
	}
	if cw[yPos][xPos] != 'A' {
		return false // Not an A
	}

	// Return whether 2 diagonal "MAS" are found, also count backwards hits
	return (cw[yPos+1][xPos-1] == 'M' && cw[yPos-1][xPos+1] == 'S' || cw[yPos+1][xPos-1] == 'S' && cw[yPos-1][xPos+1] == 'M') &&
		(cw[yPos-1][xPos-1] == 'M' && cw[yPos+1][xPos+1] == 'S' || cw[yPos-1][xPos-1] == 'S' && cw[yPos+1][xPos+1] == 'M')
}

func (cw CrossWord) matchSubstringsForPosition(xPos int, yPos int, xDiff int, yDiff int, word string) bool {
	for i := range len(word) {
		x := xPos + i*xDiff
		y := yPos + i*yDiff
		if x < 0 || x >= len(cw) || y < 0 || y >= len(cw[0]) {
			return false // Search area out of bounds
		}
		if cw[y][x] != word[i] {
			return false // Mismatch at current character
		}
	}
	return true
}

func newCrossWord(horizontals []string) CrossWord {
	crossWord := make(CrossWord, len(horizontals))
	for i, word := range horizontals {
		crossWord[i] = []uint8(word)
	}
	return crossWord
}
