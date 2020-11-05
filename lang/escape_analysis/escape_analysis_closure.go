package escapeanalysis

// 闭包对逃逸分析的影响
// 闭包, 在匿名函数中引用局部变量
// go guild -gcflas " -m -m "
func escapeAnalysisClosure() {
	counter := createCounter()
	println(counter())
	println(counter())
	createCounterWithoutReturn()
}

type counter (func() int)

func createCounter() counter {
	count := 0 // escape to heap
	return func() int {
		count++ // reference
		return count
	}
}

func createCounterWithoutReturn() func() {
	count := 0 // escape
	return func() {
		count++ // reference
	}
}
