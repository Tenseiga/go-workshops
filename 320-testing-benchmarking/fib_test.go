package slow

import "testing"

func benchmarkFib(i int, b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(i)
	}
}

func BenchmarkFib01(b *testing.B) { benchmarkFib(1, b) }
func BenchmarkFib02(b *testing.B) { benchmarkFib(2, b) }
func BenchmarkFib03(b *testing.B) { benchmarkFib(3, b) }
func BenchmarkFib10(b *testing.B) { benchmarkFib(10, b) }
func BenchmarkFib20(b *testing.B) { benchmarkFib(20, b) }
func BenchmarkFib40(b *testing.B) { benchmarkFib(40, b) }
func BenchmarkFib42(b *testing.B) { benchmarkFib(42, b) }
