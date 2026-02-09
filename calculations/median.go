package calculations

import (
	"math"
	"sort"
)

func Median(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Avoid calculating median for an empty slice
	}

	sort.Ints(numbers) // Sort the numbers first

	res := 0

	n := len(numbers)
	if n%2 != 0 {
		res = numbers[(n-1)/2] // Return middle element for odd length
	} else {
		res = int(math.Round(float64(numbers[(n-1)/2]) + float64(numbers[n/2])/2.0))
	}

	// Return the average of the two middle elements for even length
	return res
}
