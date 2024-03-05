package maxvalue

import "math"

func max_value(nums []float64) float64 {

	if len(nums) == 0 {
		return 0
	}

	/*
	 *set maximum to Infinity
	 *to handle negative numbers
	 */
	max := math.Inf(-1)

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
