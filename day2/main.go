package main

import (
	"fmt"
	aoc "github.com/sir-jacques/advent-of-code-2024/helpers"
)

type Report []int

func main() {
	// Read input
	input := aoc.ReadInput("input.txt")

	// Parse reports
	var reports []Report
	for _, line := range input {
		reports = append(reports, aoc.ParseSeperatedInts(line, " "))
	}

	// Part 1
	count := 0
	for _, report := range reports {
		if isSafe(report) {
			count++
		}
	}
	fmt.Println(count)

	// Part 2
	dampenedCount := 0
	for _, report := range reports {
		if isSafeDampened(report) {
			dampenedCount++
		}
	}
	fmt.Println(dampenedCount)
}

func isSafeDampened(report Report) bool {
	for _, r := range generateDampenedReports(report) {
		if isSafe(r) {
			return true
		}
	}
	return false
}

func generateDampenedReports(report Report) []Report {
	var result []Report
	result = append(result, report)
	for i := range report {
		subresult := Report{}
		for j, val := range report {
			if j != i {
				subresult = append(subresult, val)
			}
		}
		result = append(result, subresult)
	}
	return result
}

func isSafe(report Report) bool {
	decreasing := report[0] > report[1]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Decreasing report
		if decreasing && (diff > -1 || diff < -3) {
			return false
		}

		// Increasing report
		if !decreasing && (diff < 1 || diff > 3) {
			return false
		}
	}
	return true
}
