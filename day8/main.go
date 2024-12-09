package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
	"math"
)

// type Antenna aoc.Point
type AntennaSet []aoc.Point
type AntennaCombination struct {
	first, second aoc.Point
}
type Map struct {
	antennaSets []AntennaSet
	size        aoc.Point
}

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	var _map Map
	_map.size.X, _map.size.Y = len(input), len(input[0])

	antennas := make(map[rune]AntennaSet)
	for y, line := range input {
		for x, char := range line {
			if char == '.' {
				continue // Null character
			}
			antennas[char] = append(antennas[char], aoc.Point{X: x, Y: y})
		}
	}
	for _, a := range antennas {
		_map.antennaSets = append(_map.antennaSets, a)
	}

	// Part 1
	antiNodes1 := _map.getAntiNodes(false)
	fmt.Println(len(antiNodes1))

	// Part 2
	antiNodes2 := _map.getAntiNodes(true)
	fmt.Println(len(antiNodes2))
}

// getAntiNodes returns list of distinct antinodes (points) in map
func (m *Map) getAntiNodes(allHarmonies bool) []aoc.Point {
	// Calculate unique antinodes using map
	allAntiNodes := make(map[aoc.Point]bool)
	for _, antennaSet := range m.antennaSets {
		// For each combination of 2 inside this antennaset, calculate antinodes
		combinations := antennaSet.getAllCombinations()
		for _, combination := range combinations {
			antiNodes := combination.getAntiNodes(m.size, allHarmonies)
			for _, antiNode := range antiNodes {
				allAntiNodes[antiNode] = true
			}
		}
	}

	// Convert to list, return
	var result []aoc.Point
	for pos, _ := range allAntiNodes {
		result = append(result, pos)
	}
	return result
}

// getAllCombinations returns list of all combinations of 2 antennas in antennaset
func (as AntennaSet) getAllCombinations() (result []AntennaCombination) {
	for i := 0; i < len(as)-1; i++ {
		for j := i + 1; j < len(as); j++ {
			result = append(result, AntennaCombination{first: as[i], second: as[j]})
		}
	}
	return result
}

// getAntiNodes returns list of distinct antinodes (points) in antennacomibination
func (ac *AntennaCombination) getAntiNodes(mapSize aoc.Point, allHarmonies bool) (result []aoc.Point) {
	diff := ac.first.Sub(ac.second)

	firstHit := ac.first.Add(diff)
	if !firstHit.OutOfBounds(mapSize) {
		result = append(result, firstHit)
	}

	secondHit := ac.second.Sub(diff)
	if !secondHit.OutOfBounds(mapSize) {
		result = append(result, secondHit)
	}

	if allHarmonies {
		// Loop over all possible harmonies, stop when both hits are out-of-bounds
		for harmony := range math.MaxInt {
			firstHit = ac.first.Add(diff.ScalarMult(harmony))
			if !firstHit.OutOfBounds(mapSize) {
				result = append(result, firstHit)
			}

			secondHit = ac.second.Sub(diff.ScalarMult(harmony))
			if !secondHit.OutOfBounds(mapSize) {
				result = append(result, secondHit)
			}

			if firstHit.OutOfBounds(mapSize) && secondHit.OutOfBounds(mapSize) {
				break // Early abort
			}
		}
	}

	return result
}
