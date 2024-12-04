package helpers

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func CountOccurrencesInList(slice []int, element int) int {
	count := 0
	for _, v := range slice {
		if v == element {
			count++
		}
	}
	return count
}
