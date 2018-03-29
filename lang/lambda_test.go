package lang

import "testing"

func f(n int) func(int) int {
	return func(i int) int {
		n += i
		return n
	}
}

func TestFn(t *testing.T) {
	ff := f(10)
	t.Log(ff(1))
	t.Log(ff(2))
}
