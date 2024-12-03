package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	input := readInput(filepath.Join(filepath.Dir(filename), "input.txt"))

	// Parse reports
	var reports [][]int
	for _, line := range input {
		nums := strings.Fields(line)
		var report []int
		for _, num := range nums {
			val, _ := strconv.Atoi(num)
			report = append(report, val)
		}
		reports = append(reports, report)
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

func isSafeDampened(report []int) bool {
	for _, r := range generateDampenedReports(report) {
		if isSafe(r) {
			return true
		}
	}
	return false
}

func generateDampenedReports(report []int) [][]int {
	var result [][]int
	result = append(result, report)
	for i := range report {
		subresult := []int{}
		for j, val := range report {
			if j != i {
				subresult = append(subresult, val)
			}
		}
		result = append(result, subresult)
	}
	return result
}

func isSafe(report []int) bool {
	decreasing := report[0] > report[1]
	for i := 1; i < len(report); i++ {
		diff := report[i] - report[i-1]

		// Decreasing report
		if decreasing && diff != -1 && diff != -2 && diff != -3 {
			return false
		}

		// Increasing report
		if !decreasing && diff != 1 && diff != 2 && diff != 3 {
			return false
		}
	}
	return true
}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
