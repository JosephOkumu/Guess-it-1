package guessit

import "math"

// Calculate the range
func PredictRange(numbers []float64) (int, int) {
	start := len(numbers) - 4
	if start < 0 {
		start = 0
	}
	slice := numbers[start:]
	mean := Average(slice)
	varyance := Variance(slice)
	stdDeviation := StdDev(varyance)

	lower := mean - (1.8 * stdDeviation)
	upper := mean + (1.8 * stdDeviation)

	return int(math.Round(lower)), int(math.Round(upper))

}
