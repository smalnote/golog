package lang

import "testing"

func foo(s ...string) []string {
	s[1] = "golang"
	return s
}

func TestArrayParameterPassing(t *testing.T) {
	s := []string{"hello", "world"}
	foo(s...)
	t.Logf("s[1] = %s", s[1])
}
