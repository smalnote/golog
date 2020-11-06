package escapeanalysis

// go build -gcflags "-m -m"
func escapeAnalysisGoFunc() {
	a := make([]int, 100)
	go func(a []int) { // func literal escapes to heap
		for i := range a {
			a[i] = i
		}
	}(a) // a escape to heap for its a parameter of escape func literal

	l := 100
	b := make([]int, l) // non-constant size slice escapes to heap
	b = b[:]

	c := make([]int, 100) // does not escape
	c = c[:]
}
