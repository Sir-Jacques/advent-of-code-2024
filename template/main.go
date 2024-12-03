package main

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"runtime"
	"strings"
)

func main() {
	// Read input
	_, filename, _, _ := runtime.Caller(0)
	moduleDir := filepath.Dir(filename)
	input := readInput(fmt.Sprintf("%s/input.txt", moduleDir))

}

func readInput(filename string) []string {
	content, err := ioutil.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}

	return strings.Split(string(content), "\n")
}
