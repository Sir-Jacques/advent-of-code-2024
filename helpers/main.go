package helpers

import (
	"fmt"
	"os"
	"strings"
)

// ReadInput reads the input file and returns a slice of strings
func ReadInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
