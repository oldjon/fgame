package main

import (
	"testing"
)

func Benchmark_CopyObj(b *testing.B) {
	lista := make([]Seg, 1000)
	listb := make([]Seg, 1000)
	for i := 0; i < 10000000; i++ {
		copy(lista, listb)
	}
}

func Benchmark_CopyPtr(b *testing.B) {
	lista := make([]*Seg, 1000)
	for i, _ := range lista {
		lista[i] = &Seg{}
	}
	listb := make([]*Seg, 1000)
	for i, _ := range listb {
		listb[i] = &Seg{}
	}
	for i := 0; i < 10000000; i++ {
		copy(lista, listb)
	}
}
