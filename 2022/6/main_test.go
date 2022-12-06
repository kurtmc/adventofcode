package main

import (
	"testing"
)

func TestStartOfPacket(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  5,
		},
		{
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  6,
		},
		{
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  10,
		},
		{
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  11,
		},
	}

	for i, tc := range tests {
		got := StartOfPacket(tc.input)
		if got != tc.want {
			t.Fatalf("[%d]: expected: %v, got: %v", i, tc.want, got)
		}
	}
}

func TestStartOfMessage(t *testing.T) {
	tests := []struct {
		input string
		want  int
	}{
		{
			input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
			want:  19,
		},
		{
			input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
			want:  23,
		},
		{
			input: "nppdvjthqldpwncqszvftbrmjlhg",
			want:  23,
		},
		{
			input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
			want:  29,
		},
		{
			input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
			want:  26,
		},
	}

	for i, tc := range tests {
		got := StartOfMessage(tc.input)
		if got != tc.want {
			t.Fatalf("[%d]: expected: %v, got: %v", i, tc.want, got)
		}
	}
}
