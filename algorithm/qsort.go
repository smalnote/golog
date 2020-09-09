package algorithm

func qsort(a []int, start, end int) {
	if start >= end {
		return
	}
	pvoit := a[start]
	var i, j int
	for i, j = start, start + 1; j <= end; j++ {
		if a[j] <= pvoit {
			i++
			a[i], a[j] = a[j], a[i]
		}
	}
	a[start], a[i] = a[i], a[start]
	qsort(a, start, i - 1)
	qsort(a, i + 1, end)
}