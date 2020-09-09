package algorithm

import "testing"
import "sort"

func TestQsort(t *testing.T) {
	
	arrays := [][]int { 
		{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		{9, 8, 8, 1, 5, 4, 3, 2, 1, 0},
		{0, 8, 8, 8, 5, 5, 4, 2, 9, 0},
		{0, 8, 8, 8, 5, 5, 4, 2, 9, 0},
		{0, 8, 8, 2, 5, 5, 4, 2, 9, 0},
		{0, 8, 8, 8, 7, 6, 4, 2, 9, 0},
		{1, 1, 1},
		{0},
		{},
		{1, 2},
		{2, 1},
	}

	for _, a := range arrays {
		testArray(a)
		t.Log(a)
	}
}

func testArray(a []int) {
	qsort(a, 0, len(a) - 1)
	b := copyArray(a)
	sort.Ints(b)
	checkArray(a, b)
}

func copyArray(a []int) (copy []int) {
	copy = make([]int, len(a))
	for i := 0; i < len(a); i++ {
		copy[i] = a[i]
	}
	return copy
}

func checkArray(a, expected []int) {
	if len(a) != len(expected) {
		panic("array length not equal")
	}
	for i := 0; i < len(a); i++ {
		if expected[i] != a[i] {
			panic("array element inconsistent at index ")
		}
	}
}