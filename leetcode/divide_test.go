package leetcode

import "testing"

// go test divide.go divide_test.go
// go test -run leetcode/TestDivide

func TestDivide(t *testing.T) {
	testCases := []struct {
		dividend int
		divisor int
		expected int
	}{
		{
			5,
			5, 
			1,
		},
		{
			0,
			5, 
			0,
		},
		{
			-5,
			5, 
			-1,
		},
		{
			5, 
			-1,
			-5,
		},
		{
			15,
			-3,
			-5,
		},
		{
			16,
			3, 
			5,
		},
	}
	for i, testCase := range testCases {
		got := divide(testCase.dividend, testCase.divisor)
		if got != testCase.expected {
			t.Errorf("#%d got %d, expected %d ", i, got, testCase.expected)
		}
	}

}