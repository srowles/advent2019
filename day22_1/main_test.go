package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestDeal(t *testing.T) {
	tests := map[string]struct {
		in       []int
		expected []int
	}{
		"empty": {},
		"small": {
			in:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := deal(test.in)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestCut(t *testing.T) {
	tests := map[string]struct {
		in       []int
		expected []int
		cut      int
	}{
		"empty": {},
		"positive": {
			in:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2},
			cut:      3,
		},
		"negative": {
			in:       []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected: []int{6, 7, 8, 9, 0, 1, 2, 3, 4, 5},
			cut:      -4,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := cut(test.in, test.cut)
			require.Equal(t, test.expected, actual)
		})
	}
}

func TestDealIncrement(t *testing.T) {
	tests := map[string]struct {
		in        []int
		expected  []int
		increment int
	}{
		"ex1": {
			in:        []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9},
			expected:  []int{0, 7, 4, 1, 8, 5, 2, 9, 6, 3},
			increment: 3,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			actual := dealIncrement(test.in, test.increment)
			require.Equal(t, test.expected, actual)
		})
	}

}
