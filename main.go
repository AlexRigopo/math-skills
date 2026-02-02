package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"

	"MATH-SKILLS/calculations"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		value, err := strconv.Atoi(scanner.Text())
		if err != nil {
			fmt.Println("Invalid number:", scanner.Text())
			return
		}
		numbers = append(numbers, value)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading file:", err)
		return
	}

	fmt.Println("Numbers read from file:", numbers)

	// Now you can pass `numbers` to any function

	fmt.Println("Average: ", calculations.Average(numbers))
	fmt.Println("Median: ", calculations.Median(numbers))

}
