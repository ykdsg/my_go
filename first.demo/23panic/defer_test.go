package main

import "testing"

//命令行执行
//go test -bench . defer_test.go
func sum(max int) int {
	total := 0
	for i := 0; i < max; i++ {
		total += i
	}

	return total
}

func fooWithDefer() {
	defer func() {
		sum(10)
	}()
}

func fooWithoutDefer() {
	sum(10)
}

func BenchmarkFooWithDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithDefer()
	}
}

func BenchmarkFooWithoutDefer(b *testing.B) {
	for i := 0; i < b.N; i++ {
		fooWithoutDefer()
	}

}
