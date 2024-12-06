package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type Board [][]bool

type Position struct {
	x int
	y int
}

type Guard struct {
	pos Position
	dir int
}

var board Board
var startGuard Guard

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse board
	board := make(Board, len(input))
	for y, line := range input {
		var row []bool
		for x, char := range line {
			row = append(row, char == '#')
			if char == '^' {
				startGuard = Guard{pos: Position{x: x, y: y}, dir: 0}
			}
		}
		board[y] = row
	}

	// Move while guard is inside the board
	guard := startGuard
	traces, _ := guard.calculatePath(board)
	fmt.Println(len(traces))

	// Part 2
	validObstructionCount := 0
	for y := range board {
		for x := range board[y] {
			// If position is already a wall or guard never visits here, skip
			if board[y][x] || !traces[Position{x: x, y: y}] {
				continue
			}

			// Copy and modify board, reset guard
			modifiedBoard := Board(aoc.Copy2DSlice(board))
			modifiedBoard[y][x] = true

			// Move while guard is inside the board
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
	newX := pos.x + xDiff
	newY := pos.y + yDiff
	if newX < 0 || newX >= len(board[0]) || newY < 0 || newY >= len(board) {
		return true // Out of bounds is okay
	}
	return !board[newY][newX] // If there is a wall, it's not a valid move
}

func (g *Guard) calculatePath(board Board) (traces map[Position]bool, loop bool) {
	g.pos = startGuard.pos
	g.dir = startGuard.dir

	traces = make(map[Position]bool)
	guardStates := make(map[Guard]bool)
	for g.isInBounds(board) {

		debug := g.move(board, guardStates, traces)
		if debug {
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

	// Check if state was seen before (entering loop)
	_, exists := guardStates[*g]
	if exists {
		return true // We've already seen this state
	}

	// Store state
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
