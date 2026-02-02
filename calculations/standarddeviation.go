package calculations

import "math"

func StandardDeviation(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Avoid division by zero
	}

	n := len(numbers)

	ave := Average(numbers) // Call average function

	sum := 0
	for i := 0; i < n; i++ {
		sum += (numbers[i] - ave) * (numbers[i] - ave) // Squared difference
	}

	// Calculate standard deviation and round the result
	sd := math.Sqrt(float64(sum) / float64(n))
	fsd := int(math.Round(sd)) // Round to nearest integer

	return fsd
}

