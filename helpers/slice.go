package helpers

// CountElementInList returns the number of times the element is found in the slice
func CountElementInList[T comparable](slice []T, element T) int {
	count := 0
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}
