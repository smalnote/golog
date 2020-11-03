package main

// 测试chan, slice类型的逃逸分析
// go build -gcflags " -m -m "
// 为避免main函数在包下冲突, 不命名为mian, 可临时修改为main编译运行
func escapeAnalysisChanSlice() {
	createSliceButNotReturn()
	createSliceAndReturn()
	c := createChan()
	v, ok := <-c
	if ok {
		println(v)
	}
}

//go:noinline
func createSliceButNotReturn() {
	s := make([]int, 10) // not escape to heap, in stack frame
	for i := range s {
		s[i] = i
	}
}

//go:noinline
func createSliceAndReturn() []int {
	s := make([]int, 10) // escape to heap
	for i := range s {
		s[i] = i
	}
	return s // escape
}

//go:noinline
func createChan() chan int {
	c := make(chan int, 10) // escape to heap
	c <- 0
	close(c)
	return c
}
