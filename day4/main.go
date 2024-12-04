package main

import (
	"fmt"
	"os"
	"strings"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type CrossWord struct {
	grid [][]uint8
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	crossWord := newCrossWord(input)

	// Part 1
	xmasCount := 0
	for y := 0; y < len(crossWord.grid); y++ {
		for x := 0; x < len(crossWord.grid[y]); x++ {
			xmasCount += crossWord.matchSubstringAllDirs(x, y, "XMAS")
		}
	}

	// Part 2
	crossMasCount := 0
	for y := 0; y < len(crossWord.grid); y++ {
		for x := 0; x < len(crossWord.grid[y]); x++ {
			if crossWord.posIsCrossMas(x, y) {
				crossMasCount++
			}
		}
	}

	fmt.Println(xmasCount)
	fmt.Println(crossMasCount)
}

func (cw *CrossWord) matchSubstringAllDirs(xPos int, yPos int, word string) int {
	count := 0
	for xDiff := -1; xDiff <= 1; xDiff++ {
		for yDiff := -1; yDiff <= 1; yDiff++ {
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
func (cw *CrossWord) posIsCrossMas(xPos int, yPos int) bool {
	if xPos < 1 || yPos < 1 || xPos >= len(cw.grid[0])-1 || yPos >= len(cw.grid)-1 {
		return false // Search area out of bounds
	}
	if cw.grid[yPos][xPos] != 'A' {
		return false // Not an A
	}

	// Return whether 2 diagonal "MAS" are found, also count backwards hits
	return (cw.grid[yPos+1][xPos-1] == 'M' && cw.grid[yPos-1][xPos+1] == 'S' || cw.grid[yPos+1][xPos-1] == 'S' && cw.grid[yPos-1][xPos+1] == 'M') &&
		(cw.grid[yPos-1][xPos-1] == 'M' && cw.grid[yPos+1][xPos+1] == 'S' || cw.grid[yPos-1][xPos-1] == 'S' && cw.grid[yPos+1][xPos+1] == 'M')
}

func (cw *CrossWord) matchSubstringsForPosition(xPos int, yPos int, xDiff int, yDiff int, word string) bool {
	for i := 0; i < len(word); i++ {
		x := xPos + i*xDiff
		y := yPos + i*yDiff
		if x < 0 || x >= len(cw.grid) || y < 0 || y >= len(cw.grid[0]) {
			return false // Search area out of bounds
		}
		if cw.grid[y][x] != word[i] {
			return false // Mismatch at current character
		}
	}
	return true
}

func newCrossWord(horizontals []string) CrossWord {
	grid := make([][]uint8, len(horizontals))
	for i, word := range horizontals {
		grid[i] = []uint8(word)
	}
	return CrossWord{grid: grid}
}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
