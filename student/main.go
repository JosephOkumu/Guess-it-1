package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 1 {
		fmt.Println("Usage: go run <main.go>")
		return
	}
	scanner := bufio.NewScanner(os.Stdin)
	var numbers []float64

	for scanner.Scan() {
		input := scanner.Text()
		num, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error parsing input: %v\n", err)
			continue
		}

		numbers = append(numbers, num)
		if len(numbers) > 1 {
			lower, upper := predictRange(numbers)
			fmt.Printf("%d %d\n", lower, upper)
		}		
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}

func predictRange(numbers []float64) (int, int) {
	start := len(numbers) -4
	if start < 0 {
		start = 0
	}
	slice := numbers[start:]
	mean := average(slice)
	varyance := variance(slice)
	stdDeviation := stdDev(varyance) 

	lower := mean - (1.8 * stdDeviation)
	upper := mean + (1.8 * stdDeviation)

	return int(math.Round(lower)), int(math.Round(upper))

}

func average(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func variance(numbers []float64) float64 {
	mean := average(numbers)
	varSum := 0.0
	for _, num := range numbers {
		varSum += math.Pow(num-mean, 2)
	}
	return varSum / float64(len(numbers))
}

func stdDev(variance float64) float64 {
	return math.Sqrt(variance)
}
