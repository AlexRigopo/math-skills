package calculations

func Average(numbers []int) int {
	if len(numbers) == 0 {
		return 0 // Avoid division by zero
	}

	sum := 0
	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}

	res := int(float64(sum)/float64(len(numbers)) + 0.5)

	return res
}
