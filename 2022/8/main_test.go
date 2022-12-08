package main

import (
	"testing"
)

func TestStartOfPacket(t *testing.T) {
	tests := []struct {
		input struct {
			grid [][]int
			x, y int
		}
		want int
	}{
		{
			input: struct {
				grid [][]int
				x, y int
			}{
				grid: [][]int{{3, 0, 3, 7, 3}, {2, 5, 5, 1, 2}, {6, 5, 3, 3, 2}, {3, 3, 5, 4, 9}, {3, 5, 3, 9, 0}},
				x:    2,
				y:    1,
			},
			want: 4,
		},
		{
			input: struct {
				grid [][]int
				x, y int
			}{
				grid: [][]int{{3, 0, 3, 7, 3}, {2, 5, 5, 1, 2}, {6, 5, 3, 3, 2}, {3, 3, 5, 4, 9}, {3, 5, 3, 9, 0}},
				x:    2,
				y:    3,
			},
			want: 8,
		},
	}

	for i, tc := range tests {
		got := scenicScore(tc.input.grid, tc.input.x, tc.input.y)
		if got != tc.want {
			t.Fatalf("[%d]: expected: %v, got: %v", i, tc.want, got)
		}
	}
}
