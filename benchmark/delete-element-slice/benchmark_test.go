package delete_elment_slice

import (
	"testing"
)

func Benchmark1(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := []string{"A", "B", "C", "D", "E"}
		i := 2
		a[i] = a[len(a)-1]
		a = a[:len(a)-1]
	}
}

func Benchmark2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := []string{"A", "B", "C", "D", "E"}
		i := 2
		a = append(a[:i], a[i+1:]...)
	}
}

func Benchmark3(b *testing.B) {
	for n := 0; n < b.N; n++ {
		a := []string{"A", "B", "C", "D", "E"}
		i := 2
		copy(a[i:], a[i+1:])
		a = a[:len(a)-1]
	}
}
