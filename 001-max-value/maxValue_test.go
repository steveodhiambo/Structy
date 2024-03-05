package maxvalue

import "testing"

type testCase struct {
	name           string
	input          []float64
	expectedOutput float64
}

var testCases = []testCase{
	{
		name:           "test_00",
		input:          []float64{4, 7, 2, 8, 10, 9},
		expectedOutput: 10,
	},
	{
		name:           "test_01",
		input:          []float64{10, 5, 40, 40.3},
		expectedOutput: 40.3,
	},
	{
		name:           "test_02",
		input:          []float64{-5, -2, -1, -11},
		expectedOutput: -1,
	},
	{
		name:           "test_03",
		input:          []float64{42},
		expectedOutput: 42,
	},
	{
		name:           "test_04",
		input:          []float64{1000, 8},
		expectedOutput: 1000,
	},
	{
		name:           "test_05",
		input:          []float64{1000, 8, 9000},
		expectedOutput: 9000,
	},
	{
		name:           "test_06",
		input:          []float64{2, 5, 1, 1, 4},
		expectedOutput: 5,
	},
	// EDGE CASES
	{
		name:           "empty_slice",
		input:          []float64{},
		expectedOutput: 0,
	},
	{
		name:           "max_float64",
		input:          []float64{-1, 0, 1.7976931348623157e+308},
		expectedOutput: 1.7976931348623157e+308,
	},
	{
		name:           "identical_elements",
		input:          []float64{3.14, 3.14, 3.14, 3.14},
		expectedOutput: 3.14,
	},
	{
		name:           "mixed_negative_positive",
		input:          []float64{-100, 0, 50, -50, 100},
		expectedOutput: 100,
	},
	{
		name:           "very_small_differences",
		input:          []float64{1.0001, 1.0002, 1.0003},
		expectedOutput: 1.0003,
	},
}

func TestMax_value(t *testing.T) {

	for _, tc := range testCases {
		output := max_value(tc.input)

		if output != tc.expectedOutput {
			t.Errorf("Test case %q failed: expected %v, got %v", tc.name, tc.expectedOutput, output)
		}
	}

}
