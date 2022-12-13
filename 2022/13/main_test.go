package main

import (
	"reflect"
	"testing"
)

func TestExpandGrid(t *testing.T) {
	tests := []struct {
		input struct {
			packet string
		}
		want []interface{}
	}{
		{
			input: struct {
				packet string
			}{
				packet: "[1,1,3,1,1]",
			},
			want: []interface{}{float64(1), float64(1), float64(3), float64(1), float64(1)},
		},
		{
			input: struct {
				packet string
			}{
				packet: "[1,[2,[3,[4,[5,6,7]]]],8,9]",
			},
			want: []interface{}{float64(1), []interface{}{float64(2), []interface{}{float64(3), []interface{}{float64(4), []interface{}{float64(5), float64(6), float64(7)}}}}, float64(8), float64(9)},
		},
	}

	for i, tc := range tests {
		got := convert(tc.input.packet)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("[%d]: expected: %v, got: %v", i, tc.want, got)
		}
	}
}
