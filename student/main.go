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
	var previousnumbers []float64
	for scanner.Scan() {
		number, err := strconv.ParseFloat(scanner.Text(), 64)
		if err != nil {
			fmt.Printf("Error: %v\n", err)
			continue
		}
		previousnumbers = append(previousnumbers, number)
		if len(previousnumbers) > 1 {
			lower, upper := guessit(previousnumbers)
			fmt.Printf("%d %d\n", lower, upper)
		} else if len(previousnumbers) == 0 {
			fmt.Printf("%v %v\n", math.NaN(), math.NaN())
		}
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error reading input: %v\n", err)
	}

}
func guessit(numbers []float64) (int, int) {
	var (
		mean     float64
		sd       float64
		sum      float64
		variance float64
	)

	n := float64(len(numbers))
	for _, v := range numbers {
		sum += v
	}
	mean = sum / n //mcalculated mean
	// calculate standard deviation
	for _, v := range numbers {
		variance += ((v - mean) * (v - mean))
	}
	sd = math.Sqrt(variance / n) //calculated standard deviation

	// Define prediction interval based on the mean and standard deviation
	lowerLimit := int(math.Round(mean - 2*sd))
	upperLimit := int(math.Round(mean + 2*sd))
	return lowerLimit, upperLimit
}
