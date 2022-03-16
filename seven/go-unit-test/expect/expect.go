package expect

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func ValidSum(t *testing.T, got int, want int) {
	t.Helper()
	require.NotZerof(t, got, "Add() = %v", got)
	assert.Equal(t, want, got, "Add() = %v, want %v", got, want)
}
