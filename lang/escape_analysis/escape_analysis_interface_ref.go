package escapeanalysis

// go build -gcflags " -m -m "
func escapeAnalysisInterfaceRef() {
	a := []int{9, 8, 7, 6, 5, -1, 9, 3, 1}
	interfaceSort(ints(a)) // escape to heap, for interface-converted
}

type sortable interface {
	Len() int
	Less(i, j int) bool
	Swap(i, j int)
}

//go:noinline
func interfaceSort(s sortable) {
	quickSort(s, 0, s.Len())
}

//go:noinline
func sliceSort(a []int) {
	sliceQuickSort(a, 0, len(a))
}

func quickSort(s sortable, start, end int) {
	if start >= end {
		return
	}
	i, j := start, start+1
	for ; j < end; j++ {
		if s.Less(j, start) {
			i++
			s.Swap(i, j)
		}
	}
	s.Swap(start, i)
	quickSort(s, start, i)
	quickSort(s, i+1, end)
}

func sliceQuickSort(a []int, start, end int) {
	if start >= end {
		return
	}
	i, j := start, start+1
	for ; j < end; j++ {
		if a[j] < a[start] {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[start], a[i] = a[i], a[start]
	sliceQuickSort(a, start, i)
	sliceQuickSort(a, i+1, end)
}

type ints []int

func (a ints) Len() int {
	return len(a)
}

func (a ints) Less(i, j int) bool {
	return a[i] < a[j]
}

func (a ints) Swap(i, j int) {
	a[i], a[j] = a[j], a[i]
}
