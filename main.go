package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("data.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()

	var numbers []int

	scanner := bufio.NewScanenr(file)
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
}
