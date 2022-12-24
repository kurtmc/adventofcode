package main

import (
	"reflect"
	"testing"
)

func TestConvert(t *testing.T) {
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

func TestCompare(t *testing.T) {
	tests := []struct {
		input struct {
			a, b []interface{}
		}
		want int
	}{
		{
			input: struct {
				a, b []interface{}
			}{
				a: []interface{}{float64(1), float64(1), float64(3), float64(1), float64(1)},
				b: []interface{}{float64(1), float64(1), float64(3), float64(1), float64(1)},
			},
			want: 0,
		},
		{
			input: struct {
				a, b []interface{}
			}{
				a: []interface{}{float64(1), float64(1), float64(3), float64(1), float64(1)},
				b: []interface{}{float64(1), float64(1), float64(5), float64(1), float64(1)},
			},
			want: -1,
		},
		{
			input: struct {
				a, b []interface{}
			}{
				a: []interface{}{[]interface{}{float64(1)}, []interface{}{float64(2), float64(3), float64(4)}},
				b: []interface{}{[]interface{}{float64(1)}, float64(4)},
			},
			want: -1,
		},
	}

	for i, tc := range tests {
		got := compare(tc.input.a, tc.input.b)
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("[%d]: expected: %v, got: %v", i, tc.want, got)
		}
	}
}

func TestCompare2(t *testing.T) {
	tests := []struct {
		input struct {
			a, b string
		}
		want int
	}{
		{
			input: struct {
				a, b string
			}{
				a: "[1,1,3,1,1]",
				b: "[1,1,5,1,1]",
			},
			want: -1,
		},
		{
			input: struct {
				a, b string
			}{
				a: "[[1],[2,3,4]]",
				b: "[[1],4]",
			},
			want: -1,
		},
		{
			input: struct {
				a, b string
			}{
				a: "[9]",
				b: "[[8,7,6]]",
			},
			want: 1,
		},
		{
			input: struct {
				a, b string
			}{
				a: "[[4,4],4,4]",
				b: "[[4,4],4,4,4]",
			},
			want: -1,
		},
		{
			input: struct {
				a, b string
			}{
				a: "[7,7,7,7]",
				b: "[7,7,7]",
			},
			want: 1,
		},
		{
			input: struct {
				a, b string
			}{
				a: "[]",
				b: "[3]",
			},
			want: -1,
		},
		{
			input: struct {
				a, b string
			}{
				a: "[[[]]]",
				b: "[[]]",
			},
			want: 1,
		},
		{
			input: struct {
				a, b string
			}{
				a: "[1,[2,[3,[4,[5,6,7]]]],8,9]",
				b: "[1,[2,[3,[4,[5,6,0]]]],8,9]",
			},
			want: 1,
		},
	}

	for i, tc := range tests {
		got := compare(convert(tc.input.a), convert(tc.input.b))
		if !reflect.DeepEqual(tc.want, got) {
			t.Fatalf("[%d]: expected: %v, got: %v", i, tc.want, got)
		}
	}
}
