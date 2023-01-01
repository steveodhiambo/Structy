package isprime

import "testing"

type testCase struct {
	name           string
	input          int
	expectedOutput bool
}

var testCases = []testCase{
	{
		name:           "test_00",
		input:          2,
		expectedOutput: true,
	},
	{
		name:           "test_01",
		input:          3,
		expectedOutput: true,
	},
	{
		name:           "test_02",
		input:          4,
		expectedOutput: false,
	},
	{
		name:           "test_03",
		input:          5,
		expectedOutput: true,
	},
	{
		name:           "test_04",
		input:          6,
		expectedOutput: false,
	},
	{
		name:           "test_05",
		input:          2017,
		expectedOutput: true,
	},
	{
		name:           "test_06",
		input:          713,
		expectedOutput: false,
	},
}

func TestIsPrime(t *testing.T) {

	for _, tc := range testCases {
		output := is_prime(tc.input)

		if output != tc.expectedOutput {
			t.Errorf("Test case %q failed: expected %v, got %v", tc.name, tc.expectedOutput, output)
		}
	}

}
