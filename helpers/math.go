package helpers

import (
	"golang.org/x/exp/constraints"
)

// Abs returns the abolute value of integer/float numbers
func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

func Min[T constraints.Ordered](a, b T) T {
	if a < b {
		return a
	}
	return b
}
