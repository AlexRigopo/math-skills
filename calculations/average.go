package calculations

func Average(numbers []int) float64 {

	if len(numbers) == 0 {
		return 0
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return float64(sum) / float64(len(numbers))
}
