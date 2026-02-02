package calculations

import (
	"sort"
)

func Median(numbers []int) float64 {

	if len(numbers) == 0 {
		return 0
	}

	sort.Ints(numbers)

	n := len(numbers)
	if n%2 != 0 {
		return float64(numbers[(n-1)/2])
	}

	return (float64(numbers[(n-1)/2]) + float64(numbers[n/2])) / 2
}
