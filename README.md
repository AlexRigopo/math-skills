
# Math Skills Project

This project is a simple Go-based application designed to perform statistical calculations on a list of numbers. The main program reads a set of numbers from a file and calculates various statistical measures such as the **average**, **median**, **variance**, and **standard deviation**. These calculations are implemented in separate Go files within the `calculations` subdirectory.

## Project Structure

```
math-skills/
│
├── main.go            # The main program file that handles reading the input and displaying results
└── calculations/       # Subdirectory containing functions for statistical calculations
    ├── average.go     # Contains the function for calculating the average
    ├── median.go      # Contains the function for calculating the median
    ├── variance.go    # Contains the function for calculating the variance
    └── standarddeviation.go  # Contains the function for calculating the standard deviation
```

## How to Run

### Prerequisites
Make sure you have Go installed on your machine. You can download Go from the official website: [https://golang.org/dl/](https://golang.org/dl/)

### Steps
1. Clone this repository or download the files to your local machine.
2. Ensure that the input data file (`data.txt`) is placed in the same directory as `main.go`. This file should contain a list of numbers, each on a new line.
3. Open a terminal in the project directory and run the following command:
   ```
   go run main.go
   ```

This will read the numbers from `data.txt`, calculate the average, median, variance, and standard deviation, and display the results in the terminal.

## Functions

- **Average**: Calculates the mean of the numbers.
- **Median**: Finds the middle value in the list of numbers.
- **Variance**: Measures the spread of the numbers from the mean.
- **Standard Deviation**: A measure of the amount of variation or dispersion of the numbers.

## Example Output

For an input file (`data.txt`) with the following numbers:
```
42
7
89
15
63
28
94
3
56
71
```

The output would be:
```
Numbers read from file: [42 7 89 15 63 28 94 3 56 71]
Average: 46
Median: 49
Variance: 983.8
Standard Deviation: 31
```

## License

This project is licensed under the MIT License - see the [LICENSE](LICENSE) file for details.
