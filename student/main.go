package main

import (
	"bufio"
	"fmt"
	"guessit/calc"
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
			lower, upper := guessit.PredictRange(numbers)
			fmt.Printf("%d %d\n", lower, upper)
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "Error reading input: %v\n", err)
	}
}
