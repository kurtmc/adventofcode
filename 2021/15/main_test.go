package main

import "testing"

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
}
