package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	sequence := []float64{100}

	// Read from stdin
	for scanner.Scan() {
		input := scanner.Text()

		// Convert input to float64
		fl, err := strconv.ParseFloat(input, 64)
		if err != nil {
			fmt.Println("Error converting input to float:", err)
			return
		}
		// Append to sequence

		sequence = append(sequence, fl)

		// Call the guessNextNumberRange function
		min, max := guessNextNumberRange(sequence)
		fmt.Printf("%f %f\n", min, max)
		// fmt.Fprintf(data.txt, strconv.Itoa(int(min)), strconv.Itoa(int(max)))
	}
}

func guessNextNumberRange(sequence []float64) (min, max float64) {
	if len(sequence) < 2 {
		return 0, 0 // Not enough data to make a guess
	}

	// Calculate the average difference
	totalDiff := 0.0
	for i := 1; i < len(sequence); i++ {
		totalDiff += sequence[i] - sequence[i-1]
	}
	avgDiff := totalDiff / float64((len(sequence) - 1))

	// Calculate the standard deviation of differences
	varianceSum := 0.0
	for i := 1; i < len(sequence); i++ {
		diff := sequence[i] - sequence[i-1]
		varianceSum += (diff - avgDiff) * (diff - avgDiff)
	}
	stdDev := math.Sqrt(varianceSum / float64(len(sequence)-1))

	// Predict the range
	lastNum := sequence[len(sequence)-1]
	min = lastNum - (stdDev * 0.8)
	max = lastNum + (stdDev * 0.8)

	return min, max
}
