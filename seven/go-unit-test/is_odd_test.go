package unit

import "testing"

func BenchmarkIsOdd(b *testing.B) {
	b.Run("1", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsOdd1(10001)
		}
	})
	b.Run("2", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			IsOdd2(10001)
		}
	})
}
