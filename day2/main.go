package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	moduleDir := filepath.Dir(filename)
	input := readInput(fmt.Sprintf("%s/input.txt", moduleDir))

	reports := [][]int{}

	for _, r := range input {
		// Parse space seperated ints into int slice
		reportStrings := strings.Split(r, " ")

		report := []int{}
		for _, v := range reportStrings {
			val, _ := strconv.Atoi(v)
			report = append(report, val)
		}
		reports = append(reports, report)
	}

	// Part 1
	count := 0
	for _, i := range reports {
		if isSafe(i) {
			count++
		}
	}
	fmt.Println(count)

	// Part 2
	dampenedCount := 0
	for _, i := range reports {
		if isSafeDampened(i) {
			dampenedCount++
		}
	}
	fmt.Println(dampenedCount)
}

func isSafeDampened(report []int) bool {
	dampenedReports := generateDampenedReports(report)

	for _, r := range dampenedReports {
		if isSafe(r) {
			return true
		}
	}

	return false
}

func generateDampenedReports(report []int) [][]int {
	result := [][]int{}

	// Keep input in output
	result = append(result, report)

	// Append output with input minus one element (all commutations)
	for i := 0; i < len(report); i++ {
		subresult := []int{}
		for j := 0; j < len(report); j++ {
			if j != i {
				subresult = append(subresult, report[j])
			}
		}
		result = append(result, subresult)
	}

	return result
}

func isSafe(report []int) bool {
	decreasing := false
	if report[0] > report[1] {
		decreasing = true
	}

	// Check if all are decreasing by 1 or 2 compared to the previous element
	for i := 1; i < len(report); i++ {
		if decreasing && report[i]-report[i-1] != -1 && report[i]-report[i-1] != -2 && report[i]-report[i-1] != -3 {
			return false
		}

		if !decreasing && report[i]-report[i-1] != 1 && report[i]-report[i-1] != 2 && report[i]-report[i-1] != 3 {
			return false
		}
	}

	return true
}

func readInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(content), "\n")
}
