package leetcode

// 给定一个无序数组a, 目标值t, 求数组中与目标值t最接近的k个数
// 构造一个小顶堆, 比较大小的标准为离目值t的远近(差的绝对值)
import "container/heap"
import "testing"

type near struct {
	t int
	a []int
}

func (n near) Len() int { return len(n.a) }
func (n near) Less(i, j int) bool {
	lhs := n.a[i] - n.t
	if lhs < 0 {
		lhs = -lhs
	}
	rhs := n.a[j] - n.t
	if rhs < 0 {
		rhs = -rhs
	}
	return lhs < rhs
}
func (n near) Swap(i, j int) {
	n.a[i], n.a[j] = n.a[j], n.a[i]
}
func (n *near) Push(elem interface{}) {
	n.a = append(n.a, elem.(int))
}
func (n *near) Pop() interface{} {
	l := len(n.a) - 1
	elem := n.a[l]
	n.a = n.a[:l]
	return elem
}

func nearestKNumbers(a []int, t int, k int) []int {
	n := &near{
		t,
		a, 
	}
	heap.Init(n)
	res := make([]int, k)
	for i := 0; i < k; i++ {
		res[i] = heap.Pop(n).(int)
	}
	return res
}

func TestNearestKNumbers(t *testing.T) {
	testCases := []struct {
		t int
		a []int
		k int
		e []int
	} {
		{
			0,
			[]int { -10, -1, 0, 1, 2, 3, 4, 5, 20 },
			3,
			[]int { 0, -1, 1},
		},
		{
			20,
			[]int { -10, -1, 0, 1, 2, 3, 4, 5, 20 },
			2,
			[]int { 20, 5},
		},
		{
			-5,
			[]int { 10, 1009, 3920, -100, -200, 100, 5},
			3,
			[]int {5, 10, -100},
		},
	}
	for i, tc := range testCases {
		re := nearestKNumbers(tc.a, tc.t, tc.k)
		for j, v := range re {
			if v != tc.e[j] {
				t.Errorf("#%d, expected %v, got %v \n", i, tc.e, re)
				break
			}
		}
	}
}