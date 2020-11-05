package escapeanalysis

import (
	"math/rand"
	"sort"
	"testing"
	"time"
)

// go test -run none -bench InterfaceSort -benchtime 3s -benchmem -memprofile mem.out
// go tool pprof -alloc_space escape_analysis.test mem.out
// list Benchmark

func BenchmarkSliceSort(b *testing.B) {
	a := generateRandomInts(100)

	for i := 0; i < b.N; i++ {
		sliceSort(a)
	}
}

func BenchmarkInterfaceSort(b *testing.B) {
	a := generateRandomInts(100)

	for i := 0; i < b.N; i++ {
		interfaceSort(ints(a))
	}
}

func TestSliceSort(t *testing.T) {
	a := generateRandomInts(10)
	b := make([]int, len(a))
	copy(b, a)
	sort.Ints(a)
	sliceSort(b)

	for i, vb := range b {
		if a[i] != vb {
			t.Errorf("element at %d, expected %d, got %d\n", i, a[i], vb)
		}
	}
}

func generateRandomInts(len int) []int {
	a := make([]int, len)

	rand.Seed(time.Now().UnixNano())
	for i := range a {
		a[i] = rand.Intn(len)
	}

	return a
}
