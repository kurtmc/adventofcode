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
