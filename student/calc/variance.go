package guessit

import "math"

// Calculate variance
func Variance(numbers []float64) float64 {
	mean := Average(numbers)
	varSum := 0.0
	for _, num := range numbers {
		varSum += math.Pow(num-mean, 2)
	}
	return varSum / float64(len(numbers))
}
