package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

type CrossWord struct {
	horizontals []string
	verticals   []string
	diagonalsLR []string
	diagonalsRL []string
}

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	input := readInput(filepath.Join(filepath.Dir(filename), "input.txt"))

	crossWord := CrossWord{horizontals: input}
	crossWord.generateVerticals()
	crossWord.generateDiagonalsLR()
	crossWord.generateDiagonalsRL()

	fmt.Println("debug")

	// Parse input into crossword

	// Part 1
	xmasCount := 0
	for _, word := range crossWord.horizontals {
		xmasCount += countSubstring(word, "XMAS")
		xmasCount += countSubstring(word, "SAMX")
	}
	for _, word := range crossWord.verticals {
		xmasCount += countSubstring(word, "XMAS")
		xmasCount += countSubstring(word, "SAMX")
	}
	for _, word := range crossWord.diagonalsRL {
		xmasCount += countSubstring(word, "XMAS")
		xmasCount += countSubstring(word, "SAMX")
	}
	for _, word := range crossWord.diagonalsLR {
		xmasCount += countSubstring(word, "XMAS")
		xmasCount += countSubstring(word, "SAMX")
	}
	// Part 2

	fmt.Println(xmasCount)
}

func countSubstring(input string, substring string) int {
	fmt.Printf("string: %s\n matches: %d\n\n", input, strings.Count(input, substring))
	return strings.Count(input, substring)
}

func reverseString(s string) string {
	// Convert string to a slice of runes
	runes := []rune(s)

	// Reverse the slice of runes
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}

	// Convert the slice of runes back to a string
	return string(runes)
}

func (cw *CrossWord) generateVerticals() {
	for i := range cw.horizontals {
		var vertical string
		for j := range cw.horizontals {
			vertical += string(cw.horizontals[j][i])
		}
		cw.verticals = append(cw.verticals, vertical)
	}
}

func (cw *CrossWord) generateDiagonalsLR() {
	x := 0
	y := 0

	// Bottom left
	for x < len(cw.horizontals) {
		var diagonal string
		for i := x; i < len(cw.horizontals); i++ {
			diagonal += string(cw.horizontals[i][y])
			y++
		}
		cw.diagonalsLR = append(cw.diagonalsLR, diagonal)
		x++
		y = 0
	}

	// Top right
	x = 0
	y = 1 // Exclude duplicate entry of first diagonal
	for y < len(cw.horizontals[0]) {
		var diagonal string
		for i := y; i < len(cw.horizontals[0]); i++ {
			diagonal += string(cw.horizontals[x][i])
			x++
		}
		cw.diagonalsLR = append(cw.diagonalsLR, diagonal)
		x = 0
		y++
	}
}

func (cw *CrossWord) generateDiagonalsRL() {
	// Top left
	for x := len(cw.horizontals) - 1; x >= 0; x-- {
		var diagonal string
		for i := x; i >= 0; i-- {
			diagonal += string(cw.horizontals[i][x-i])
		}
		cw.diagonalsRL = append(cw.diagonalsRL, diagonal)
	}

	// Bottom right
	for y := 1; y < len(cw.horizontals[0]); y++ {
		var diagonal string
		for i := y; i < len(cw.horizontals[0]); i++ {
			diagonal += string(cw.horizontals[len(cw.horizontals[0])-1-i][i])
		}
		cw.diagonalsRL = append(cw.diagonalsRL, diagonal)
	}
}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
