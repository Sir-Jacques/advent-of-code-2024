package helpers

import (
	"golang.org/x/exp/constraints"
)

func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func CountOccurrencesInList[T comparable](slice []T, element T) int {
	count := 0
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}
