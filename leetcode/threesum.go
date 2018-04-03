package leetcode

import "sort"

// Given an array S of n integers, are there elements
// a, b, c in S such that a + b + c = 0? Find all unique
// triplets in the array which gives the sum of zero.
// NOTE: The solution set must not contain duplicate triplets.

func threeSum(nums []int) [][]int {
	s := make([][]int, len(nums))[0:0]
	m2 := make(map[int][][2]int, len(nums))
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			m2[v+nums[j]] = append(m2[v+nums[j]], [2]int{i, j})
		}
	}

	for i, v := range nums {
		if t2s, ok := m2[-v]; ok {
			for _, t2 := range t2s {
				if i == t2[0] || i == t2[1] {
					continue
				}
				e := []int{nums[i], nums[t2[0]], nums[t2[1]]}
				sort.Ints(e)
				existed := false
				for _, ee := range s {
					if e[0] == ee[0] && e[1] == ee[1] && e[2] == ee[2] {
						existed = true
					}
				}
				if !existed {
					s = append(s, e)
				}
			}
		}
	}
	return s
}
