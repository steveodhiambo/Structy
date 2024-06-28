package main

import (
	"fmt"
	"math"
)

func max_value(nums []float64) float64 {
	if len(nums) == 0 {
		return 0
	}
	max := math.Inf(-1)

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}

func main() {
	numbers := []float64{1.2, 3.4, 2.5, 4.6, 0.7}
	maxNum := max_value(numbers)
	fmt.Printf("The maximum value is: %.2f\n", maxNum)
}
