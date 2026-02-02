package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"MATH-SKILLS/calculations" // Importing calculations package
)

func main() {
	// Open the file containing the data
	file, err := os.Open("data.txt")
	if err != nil {
		// Error handling if the file cannot be opened
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close() // Ensure file is closed when done

	var numbers []int // Slice to store the numbers read from the file

	// Read each line from the file
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Convert the string value to an integer
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			// Error handling for invalid numbers in the file
			fmt.Println("Invalid number:", scanner.Text())
			return
		}
		numbers = append(numbers, value) // Add the number to the slice
	}

	// Check if there were any errors while reading the file
	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	// Call functions from the calculations package to calculate statistics
	fmt.Println("Average:", calculations.Average(numbers))
	fmt.Println("Median:", calculations.Median(numbers))
	fmt.Println("Variance:", calculations.Variance(numbers))
	fmt.Println("Standard Deviation:", calculations.StandardDeviation(numbers))
}
