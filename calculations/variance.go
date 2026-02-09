package calculations

import "math"

func Variance(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Avoid division by zero
	}

	n := len(numbers)
	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	ave := float64(sum) / float64(n) // Calculate average in float64

	sum2 := 0.0
	for i := 0; i < n; i++ {
		sum2 += (float64(numbers[i]) - ave) * (float64(numbers[i]) - ave) // Squared difference
	}

	res := int(math.Round(sum2 / float64(n)))

	return res
}
