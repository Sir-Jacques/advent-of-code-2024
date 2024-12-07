package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type Board [][]bool

type Position struct {
	x, y int
}

type Guard struct {
	pos Position
	dir int
}

var startGuard Guard

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse board
	board := make(Board, len(input))
	for y, line := range input {
		row := make([]bool, len(line))
		for x, char := range line {
			row[x] = char == '#' // Wall
			if char == '^' {
				startGuard = Guard{pos: Position{x: x, y: y}, dir: 0}
			}
		}
		board[y] = row
	}

	// Part 1
	guard := startGuard
	traces, _ := guard.calculatePath(board)
	fmt.Println(len(traces))

	// Part 2
	validObstructionCount := 0
	for y, row := range board {
		for x, cell := range row {
			// If position is already a wall or guard never visits here, skip
			if cell || !traces[Position{x: x, y: y}] {
				continue
			}

			// Copy board and add new obstruction
			modifiedBoard := Board(aoc.Copy2DSlice(board))
			modifiedBoard[y][x] = true

			// Check if path is a loop
			_, loop := guard.calculatePath(modifiedBoard)
			if loop {
				validObstructionCount++
			}
		}

	}
	fmt.Println(validObstructionCount)
}

func (g *Guard) isInBounds(board Board) bool {
	return g.pos.x >= 0 && g.pos.x < len(board[0]) && g.pos.y >= 0 && g.pos.y < len(board)
}

func isValidMove(board Board, pos Position, xDiff, yDiff int) bool {
	newX, newY := pos.x+xDiff, pos.y+yDiff
	if newX < 0 || newX >= len(board[0]) || newY < 0 || newY >= len(board) {
		return true // Out of bounds is okay
	}
	return !board[newY][newX] // If there is a wall, it's not a valid move
}

func (g *Guard) calculatePath(board Board) (traces map[Position]bool, loop bool) {
	// Reset guard
	g.pos, g.dir = startGuard.pos, startGuard.dir

	// Move guard until loop or out of bounds
	traces = make(map[Position]bool)
	guardStates := make(map[Guard]bool)
	for g.isInBounds(board) {
		if g.move(board, guardStates, traces) {
			return traces, true
		}
	}

	return traces, false
}

func (g *Guard) move(board Board, guardStates map[Guard]bool, traces map[Position]bool) (loop bool) {
	// Check if out of bounds
	if !g.isInBounds(board) {
		return false
	}

	// Check if state was seen before (thus entering loop)
	if _, exists := guardStates[*g]; exists {
		return true // We've already seen this state
	}

	// Store (new) state
	guardStates[*g] = true
	traces[g.pos] = true

	// Move guard if possible, switch direction otherwise
	if g.dir == 0 && isValidMove(board, g.pos, 0, -1) {
		g.pos.y--
	} else if g.dir == 1 && isValidMove(board, g.pos, 1, 0) {
		g.pos.x++
	} else if g.dir == 2 && isValidMove(board, g.pos, 0, 1) {
		g.pos.y++
	} else if g.dir == 3 && isValidMove(board, g.pos, -1, 0) {
		g.pos.x--
	} else {
		g.dir = (g.dir + 1) % 4
	}

	return false
}
