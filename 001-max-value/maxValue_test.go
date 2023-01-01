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
}

func TestMax_value(t *testing.T) {
	for _, tc := range testCases {
		output := Max_value(tc.input)
		if output != tc.expectedOutput {
			t.Errorf("Test case %q failed: expected %v, got %v", tc.name, tc.expectedOutput, output)
		}
	}

}
