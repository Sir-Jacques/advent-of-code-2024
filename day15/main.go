package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"slices"
)

type Board struct {
	boxes    map[aoc.Point]int
	walls    map[aoc.Point]bool
	robot    Robot
	size     aoc.Point
	boxWidth int
}

type Robot struct {
	position aoc.Point
}

type Box struct {
	id        int
	positions []aoc.Point
}

var (
	UP    = aoc.Point{X: 0, Y: -1}
	DOWN  = aoc.Point{X: 0, Y: 1}
	LEFT  = aoc.Point{X: -1, Y: 0}
	RIGHT = aoc.Point{X: 1, Y: 0}
)

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Part 1
	board, instructions := parseBoard(input, 1)
	for _, instruction := range instructions {
		board = moveRobot(board, instruction)
	}
	fmt.Println(getGPSSum(board))

	// Part 2
	board, instructions = parseBoard(input, 2)
	for _, instruction := range instructions {
		board = moveRobot(board, instruction)
	}
	fmt.Println(getGPSSum(board))
}

func parseBoard(input []string, boxWidth int) (Board, []aoc.Point) {
	board := Board{boxWidth: boxWidth, boxes: make(map[aoc.Point]int), walls: make(map[aoc.Point]bool)}
	var instructions []aoc.Point
	segment, boxIndex := 0, 1
	for y, line := range input {
		if line == "" {
			segment++
			board.size = aoc.Point{X: len(input[0]) * boxWidth, Y: y}
			continue
		}

		switch segment {
		case 0:
			for x, char := range line {
				var boxPositions []aoc.Point
				for i := range boxWidth {
					boxPositions = append(boxPositions, aoc.Point{X: x*boxWidth + i, Y: y})
				}
				pos := aoc.Point{X: x * boxWidth, Y: y}
				if char == '#' {
					for _, position := range boxPositions {
						board.walls[position] = true
					}
				}
				if char == 'O' {
					for _, position := range boxPositions {
						board.boxes[position] = boxIndex
					}
					boxIndex++
				}
				if char == '@' {
					board.robot = Robot{position: pos}
				}
			}
		case 1:
			for _, char := range line {
				switch char {
				case '^':
					instructions = append(instructions, UP)
				case 'v':
					instructions = append(instructions, DOWN)
				case '<':
					instructions = append(instructions, LEFT)
				case '>':
					instructions = append(instructions, RIGHT)
				}
			}
		}
	}

	return board, instructions
}

func printBoard(b Board) {
	for y := 0; y < b.size.Y; y++ {
		for x := 0; x < b.size.X; x++ {
			pos := aoc.Point{X: x, Y: y}
			if b.robot.position == pos {
				fmt.Print("@")
			} else if b.boxes[pos] != 0 {
				if b.boxWidth == 1 {
					fmt.Print("O")
				} else if b.boxes[pos] == b.boxes[pos.Add(aoc.Point{X: 1, Y: 0})] {
					fmt.Print("[")
				} else {
					fmt.Print("]")
				}
			} else if b.walls[pos] {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

func getGPSSum(b Board) int {
	sum := 0
	for key, boxIndex := range b.boxes {
		if boxIndex != 0 && b.boxes[key.Add(aoc.Point{X: -1, Y: 0})] != boxIndex {
			sum += 100*key.Y + key.X
		}
	}
	return sum
}

func getBoxPositions(b Board, boxId int) []aoc.Point {
	var boxPositions []aoc.Point
	for position, bId := range b.boxes {
		if bId == boxId {
			boxPositions = append(boxPositions, position)
		}
	}
	return boxPositions
}

func moveRobot(b Board, direction aoc.Point) Board {
	newRobotPos := b.robot.position.Add(direction)

	// Robot moves into wall/oob
	if newRobotPos.OutOfBounds(b.size) || b.walls[newRobotPos] {
		return b
	}

	// Robot moves into box
	boxId, collision := b.boxes[newRobotPos]
	if collision && boxId != 0 {
		dependentBoxes := getDependentBoxes(b, newRobotPos, direction)
		if canMoveDependentBoxes(b, dependentBoxes, direction) {
			// Set all old states to 0
			for _, box := range dependentBoxes {
				for _, pos := range box.positions {
					b.boxes[pos] = 0
				}
			}
			// Re-populate all states with boxIds on new positions
			for _, box := range dependentBoxes {
				for _, pos := range box.positions {
					newPos := pos.Add(direction)
					b.boxes[newPos] = box.id
				}
			}
			b.robot.position = newRobotPos // Move robot too
			return b
		}
	} else {
		// Robot is free to move
		b.robot.position = newRobotPos
	}

	return b
}

func getDependentBoxes(board Board, position aoc.Point, direction aoc.Point) []Box {
	dependentBoxes := make(map[aoc.Point]bool)
	currentBoxId := board.boxes[position]
	currentBoxPositions := getBoxPositions(board, currentBoxId)

	for _, boxPos := range currentBoxPositions {
		newBoxPos := boxPos.Add(direction)
		if !newBoxPos.OutOfBounds(board.size) && !board.walls[newBoxPos] && board.boxes[newBoxPos] != 0 {
			dependentBoxes[newBoxPos] = true
		}
	}

	allDependentBoxes := []Box{{id: currentBoxId, positions: currentBoxPositions}}
	for boxPos := range dependentBoxes {
		if !slices.Contains(currentBoxPositions, boxPos) {
			allDependentBoxes = append(allDependentBoxes, getDependentBoxes(board, boxPos, direction)...)
		}
	}
	return allDependentBoxes
}

func canMoveDependentBoxes(board Board, dependentBoxes []Box, direction aoc.Point) bool {
	for _, box := range dependentBoxes {
		for _, pos := range box.positions {
			newBoxPos := pos.Add(direction)
			if newBoxPos.OutOfBounds(board.size) || board.walls[newBoxPos] {
				return false
			}
		}
	}
	return true
}
