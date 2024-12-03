package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	input := readInput(filepath.Join(filepath.Dir(filename), "input.txt"))

	// Part 1

	// Part 2

}

func readInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
