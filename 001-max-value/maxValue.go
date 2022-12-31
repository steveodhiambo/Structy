package maxvalue

import "math"

func Max_value(nums []float64) float64 {
	//set maximum to Infity
	max := math.Inf(-1)

	for _, num := range nums {
		if num > max {
			max = num
		}
	}

	return max
}
