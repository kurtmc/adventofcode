package main

import (
	"reflect"
	"testing"
)

func TestExtend(t *testing.T) {
	input := [][]int{
		{1, 1},
		{1, 1},
	}
	expected := [][]int{
		{1, 1, 2, 2},
		{1, 1, 2, 2},
		{2, 2, 3, 3},
		{2, 2, 3, 3},
	}

	actual := extend(input, 2)

	if len(expected) != len(actual) {
		t.Fatalf("expected len(actual) = %d but got %d", len(expected), len(actual))
	}
	if len(expected[0]) != len(actual[0]) {
		t.Fatalf("expected len(actual[0]) = %d but got %d", len(expected[0]), len(actual[0]))
	}

	tests := []struct {
		input      [][]int
		multiplier int
		want       [][]int
	}{
		{
			input: [][]int{
				{1, 1},
				{1, 1},
			},
			multiplier: 2,
			want: [][]int{
				{1, 1, 2, 2},
				{1, 1, 2, 2},
				{2, 2, 3, 3},
				{2, 2, 3, 3},
			},
		},
		{
			input: [][]int{
				{8},
			},
			multiplier: 5,
			want: [][]int{
				{8, 9, 1, 2, 3},
				{9, 1, 2, 3, 4},
				{1, 2, 3, 4, 5},
				{2, 3, 4, 5, 6},
				{3, 4, 5, 6, 7},
			},
		},
	}

	for _, tc := range tests {
		got := extend(tc.input, tc.multiplier)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("expected: %v, got: %v", tc.want, got)
		}
	}
}
