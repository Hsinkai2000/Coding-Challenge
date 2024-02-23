package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSum_to_n_a(t *testing.T) {
	// Test cases
	tests := []struct {
		n        int
		expected int
	}{
		{5, 15},
		{1, 1},
		{2, 3},
		{3, 6},
		{10, 55},
		{100, 5050},
		{0, 0},
	}

	// Run tests
	for _, test := range tests {
		t.Run("Should return Summation of n", func(t *testing.T) {
			got := Sum_to_n_a(test.n)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestSum_to_n_b(t *testing.T) {
	// Test cases
	tests := []struct {
		n        int
		expected int
	}{
		{5, 15},
		{1, 1},
		{2, 3},
		{3, 6},
		{10, 55},
		{100, 5050},
		{0, 0},
	}

	// Run tests
	for _, test := range tests {
		t.Run("Should return Summation of n", func(t *testing.T) {
			got := Sum_to_n_b(test.n)
			assert.Equal(t, test.expected, got)
		})
	}
}

func TestSum_to_n_c(t *testing.T) {
	// Test cases
	tests := []struct {
		n        int
		expected int
	}{
		{5, 15},
		{1, 1},
		{2, 3},
		{3, 6},
		{10, 55},
		{100, 5050},
		{0, 0},
	}

	// Run tests
	for _, test := range tests {
		t.Run("Should return Summation of n", func(t *testing.T) {
			got := Sum_to_n_c(test.n)
			assert.Equal(t, test.expected, got)
		})
	}
}
