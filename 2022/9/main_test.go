package main

import (
	"reflect"
	"testing"
)

func TestExpandGrid(t *testing.T) {
	tests := []struct {
		input struct {
			grid   [][]rune
			amount int
		}
		want [][]rune
	}{
		{
			input: struct {
				grid   [][]rune
				amount int
			}{
				grid:   [][]rune{{'.', '.'}, {'T', 'H'}},
				amount: 1,
			},
			want: [][]rune{{'.', '.', '.', '.'}, {'.', '.', '.', '.'}, {'.', 'T', 'H', '.'}, {'.', '.', '.', '.'}},
		},
	}

	for i, tc := range tests {
		got := expandGridHelper(tc.input.grid, tc.input.amount)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("[%d]: expected: %c, got: %c", i, tc.want, got)
		}
	}
}
