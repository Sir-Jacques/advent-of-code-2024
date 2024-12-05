package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type WordGrid [][]uint8

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	wordGrid := newWordGrid(input)

	// Part 1
	xmasCount := 0
	for y := 0; y < len(wordGrid); y++ {
		for x := 0; x < len(wordGrid[y]); x++ {
			xmasCount += wordGrid.matchSubstringAllDirs(x, y, "XMAS")
		}
	}

	// Part 2
	crossMasCount := 0
	for y := 0; y < len(wordGrid); y++ {
		for x := 0; x < len(wordGrid[y]); x++ {
			if wordGrid.posIsCrossMas(x, y) {
				crossMasCount++
			}
		}
	}

	fmt.Println(xmasCount)
	fmt.Println(crossMasCount)
}

func (wordGrid WordGrid) matchSubstringAllDirs(xPos int, yPos int, word string) int {
	count := 0
	for xDiff := -1; xDiff <= 1; xDiff++ {
		for yDiff := -1; yDiff <= 1; yDiff++ {
			if xDiff == 0 && yDiff == 0 {
				continue
			}
			if wordGrid.matchSubstringsForPosition(xPos, yPos, xDiff, yDiff, word) {
				count++
			}
		}
	}
	return count
}

// Current char must the be A, looking for crossing MAS in diagonal directions
func (wordGrid WordGrid) posIsCrossMas(xPos int, yPos int) bool {
	if xPos < 1 || yPos < 1 || xPos >= len(wordGrid[0])-1 || yPos >= len(wordGrid)-1 {
		return false // Search area out of bounds
	}
	if wordGrid[yPos][xPos] != 'A' {
		return false // Not an A
	}

	// Return whether 2 diagonal "MAS" are found, also count backwards hits
	return (wordGrid[yPos+1][xPos-1] == 'M' && wordGrid[yPos-1][xPos+1] == 'S' || wordGrid[yPos+1][xPos-1] == 'S' && wordGrid[yPos-1][xPos+1] == 'M') &&
		(wordGrid[yPos-1][xPos-1] == 'M' && wordGrid[yPos+1][xPos+1] == 'S' || wordGrid[yPos-1][xPos-1] == 'S' && wordGrid[yPos+1][xPos+1] == 'M')
}

func (wordGrid WordGrid) matchSubstringsForPosition(xPos int, yPos int, xDiff int, yDiff int, word string) bool {
	for i := 0; i < len(word); i++ {
		x := xPos + i*xDiff
		y := yPos + i*yDiff
		if x < 0 || x >= len(wordGrid) || y < 0 || y >= len(wordGrid[0]) {
			return false // Search area out of bounds
		}
		if wordGrid[y][x] != word[i] {
			return false // Mismatch at current character
		}
	}
	return true
}

func newWordGrid(horizontals []string) WordGrid {
	wordGrid := make(WordGrid, len(horizontals))
	for i, word := range horizontals {
		wordGrid[i] = []uint8(word)
	}
	return wordGrid
}
