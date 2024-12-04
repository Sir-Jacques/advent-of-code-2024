package helpers

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(filename string) []string {
	content, err := os.ReadFile(filename)
	if err != nil {
		fmt.Println(err)
	}
	return strings.Split(string(content), "\n")
}
