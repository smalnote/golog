package leetcode

import "testing"

func maxProfit(prices []int) int {
	if len(prices) <= 0 {
		return 0
	}
	r := 0
	pmin := prices[0]
	for i := 1; i < len(prices); i++ {
		if pmin > prices[i] {
			pmin = prices[i]
		} else {
			m := prices[i] - pmin
			if r < m {
				r = m
			}
		}
	}
	return r
}

func TestMaxProfit(t *testing.T) {
	testCases := []struct {
		prices []int
		maxProfit int
	}{
		{
			[]int{1, 2, 3, 4, 5},
			4,
		},
		{
			[]int{7, 1, 10, 4, 2, 3},
			9,
		},
		{
			[]int{7, 5, 3, 1},
			0,
		},
		{
			[]int{7, 1, 5, 3, 6, 4},
			5,
		},
	}
	for i, testCase := range testCases {
		got := maxProfit(testCase.prices)
		if got != testCase.maxProfit {
			t.Errorf("#%d got %d, expected %d ", i, got, testCase.maxProfit)
		}
	}

}