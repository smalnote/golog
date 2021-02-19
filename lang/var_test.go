package lang

import "testing"

func TestVarChan(t *testing.T) {
	var a chan struct{}
	t.Logf("var a chan struct{} == nil ? %v", a == nil)
	var b []struct{}
	t.Logf("var b []struct{} == nil ? %v", b == nil)
	var c map[string]struct{}
	t.Logf("var c map[string]struct{} == nil ? %v", c == nil)

	t.Logf("len(c) == %v", len(c))
}

func TestMaxDuration(t *testing.T) {
	inf := 1<<63 - 1
	t.Logf("%+v", inf)
}
