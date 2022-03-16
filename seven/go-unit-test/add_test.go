package unit

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.ozon.dev/qa/teachers/qa-route-256/go-unit-test/expect"
)

func TestAdd(t *testing.T) {

	// <arrange> - <assert>
	t.Run("should make sum for positive", func(t *testing.T) {
		t.Parallel()
		// arrange
		want := 3

		// action
		got := Add(1, 2)

		// assert
		assert.Equal(t, want, got, "Add() = %v, want %v", got, want)
	})

	t.Run("should make sum for negative numbers", func(t *testing.T) {
		t.Parallel()

		want := -4

		got := Add(-1, -3)

		assert.Equal(t, want, got, "Add() = %v, want %v", got, want)
	})

	t.Run("should make sum for zero", func(t *testing.T) {

		want := 1

		got := Add(1, 0)

		t.Logf("got = %d", got)
		assert.Equal(t, want, got, "Add() = %v, want %v", got, want)
		t.Run("inner test", func(t *testing.T) {
			t.Log("Passed")
		})
	})
}

func TestAdd1(t *testing.T) {
	tests := []struct {
		name string
		a    int
		b    int
		want int
	}{
		{
			name: "positive numbers",
			a:    1,
			b:    2,
			want: 3,
		},
		{
			name: "negative numbers",
			a:    -1,
			b:    -2,
			want: -3,
		},
	}
	for _, tt := range tests {
		tt := tt

		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()
			// action
			got := Add(tt.a, tt.b)

			// assert
			expect.ValidSum(t, got, tt.want)
		})
	}
}
