package leetcode

import "testing"

func maxProfit4(t int, prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	dp := make([]int, len(prices))

	for k := 0; k < t; k++ {
		predp := 0
		for i := 1; i < len(prices); i++ {
			dp[i], predp = predp, dp[i] // swap for next iteration
			min := prices[0]
			for j := 1; j < i; j++ {
				tmpMin := prices[j] - dp[j-1]
				if min > tmpMin {
					min = tmpMin
				}
			}
			tmpMax := prices[i] - min
			if dp[i] < tmpMax {
				dp[i] = tmpMax
			}
		}
	}
	return dp[len(prices)-1]
}

func TestMaxProfit4(t *testing.T) {
	testCases := []struct {
		k         int
		prices    []int
		maxProfit int
	}{
		{
			2,
			[]int{3, 3, 5, 0, 0, 3, 1, 4},
			6,
		},
	}
	for i, testCase := range testCases {
		got := maxProfit4(testCase.k, testCase.prices)
		if got != testCase.maxProfit {
			t.Errorf("#%d got %d, expected %d ", i, got, testCase.maxProfit)
		}
	}

}
