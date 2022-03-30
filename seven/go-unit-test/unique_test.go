package unit

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func BenchmarkUnique(b *testing.B) {
	sliceTest := []int{1, 2, 3, 4, 5, 6, 6, 7, 8, 8, 9, 0, 0, 0, 5, 5, 5, 3, 0, 54, 435435, 234235, 23552354}
	b.Run("1 variant", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			unique1(sliceTest)
		}
	})

	b.Run("2 variant", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			unique2(sliceTest)
		}
	})

	b.Run("3 variant", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			unique3(sliceTest)
		}
	})
}

func Test_unique2(t *testing.T) {
	tests := []struct {
		name     string
		elements []int
		want     []int
	}{
		{
			name:     "Simple test",
			elements: []int{1, 2, 3, 4, 4},
			want:     []int{1, 2, 3, 4},
		},
		{
			name:     "Not unique test",
			elements: []int{1, 2, 3, 4, 5},
			want:     []int{1, 2, 3, 4, 5},
		},
		{
			name:     "Null test",
			elements: []int{0},
			want:     []int{0},
		},
		{
			name:     "Null same test",
			elements: []int{0, 0, 0, 0, 0},
			want:     []int{0},
		},
		{
			name:     "Negative numbers test",
			elements: []int{-1, -2, -3, -4, -4},
			want:     []int{-1, -2, -3, -4},
		},
		{
			name:     "All numbers are same",
			elements: []int{5, 5, 5, 5, 5},
			want:     []int{5},
		},
		{
			name:     "Two same numbers",
			elements: []int{1, 2, 1, 3, 4, 4},
			want:     []int{1, 2, 3, 4},
		},
		{
			name:     "nil test",
			elements: nil,
			want:     nil,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			assert.Equalf(t, tt.want, unique2(tt.elements), "Func() = %v, want %v", tt.elements, tt.want)
		})
	}
}
