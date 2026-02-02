package calculations

func Average(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Avoid division by zero
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	return sum / len(numbers)
}
