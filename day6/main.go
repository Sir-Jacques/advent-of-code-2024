package main

import (
	"fmt"

	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type Board struct {
	walls      [][]bool
	startGuard Guard
}

type Guard struct {
	pos aoc.Position
	dir int
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse board
	board := Board{walls: make([][]bool, len(input))}
	for y, line := range input {
		row := make([]bool, len(line))
		for x, char := range line {
			row[x] = char == '#' // Wall
			if char == '^' {
				board.startGuard = Guard{pos: aoc.Position{X: x, Y: y}, dir: 0}
			}
		}
		board.walls[y] = row
	}

	// Part 1, count traces
	var guard Guard
	traces, _ := guard.calculatePath(board)
	fmt.Println(len(traces))

	// Part 2, add obstruction somewhere on traces (other positions are never visited by guard)
	validObstructionCount := 0
	for pos := range traces {
		// Copy board and add new obstruction
		modifiedBoard := board
		modifiedBoard.walls = aoc.Copy2DSlice(board.walls)
		modifiedBoard.walls[pos.Y][pos.X] = true

		// Check if path is a loop
		_, loop := guard.calculatePath(modifiedBoard)
		if loop {
			validObstructionCount++
		}
	}
	fmt.Println(validObstructionCount)
}

func (g *Guard) isInBounds(board Board) bool {
	return g.pos.X >= 0 && g.pos.X < len(board.walls[0]) && g.pos.Y >= 0 && g.pos.Y < len(board.walls)
}

func isValidMove(board Board, pos aoc.Position, xDiff, yDiff int) bool {
	newX, newY := pos.X+xDiff, pos.Y+yDiff
	if newX < 0 || newX >= len(board.walls[0]) || newY < 0 || newY >= len(board.walls) {
		return true // Allowed to walk out of bounds
	}
	return !board.walls[newY][newX] // If there is a wall, it's not a valid move
}

func (g *Guard) calculatePath(board Board) (map[aoc.Position]bool, bool) {
	// Reset states
	g.pos, g.dir = board.startGuard.pos, board.startGuard.dir
	traces := make(map[aoc.Position]bool)
	guardStates := make(map[Guard]bool)

	// Move guard until loop or out of bounds
	for g.isInBounds(board) {
		// Move guard
		g.move(board)

		// Check if guard is now out of bounds
		if !g.isInBounds(board) {
			return traces, false
		}

		// Check if state was seen before (thus entering loop)
		if _, exists := guardStates[*g]; exists {
			return traces, true // We've already seen this state
		}

		// Store (new) state
		guardStates[*g] = true
		traces[g.pos] = true
	}

	return traces, false
}

func (g *Guard) move(board Board) {
	// Move guard if possible, switch direction otherwise
	if g.dir == 0 && isValidMove(board, g.pos, 0, -1) {
		g.pos.Y--
	} else if g.dir == 1 && isValidMove(board, g.pos, 1, 0) {
		g.pos.X++
	} else if g.dir == 2 && isValidMove(board, g.pos, 0, 1) {
		g.pos.Y++
	} else if g.dir == 3 && isValidMove(board, g.pos, -1, 0) {
		g.pos.X--
	} else {
		g.dir = (g.dir + 1) % 4
	}
}
