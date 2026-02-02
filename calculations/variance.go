package calculations

func Variance(numbers []int) float64 {
	if len(numbers) == 0 {
		return 0 // Avoid division by zero
	}

	n := len(numbers)
	ave := Average(numbers) // Call average function

	sum := 0.0
	for i := 0; i < n; i++ {
		sum += float64((numbers[i] - ave) * (numbers[i] - ave)) // Squared difference
	}

	return sum / float64(n)
}
