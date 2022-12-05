package main

import (
	"testing"
)

func TestStack(t *testing.T) {
	//input := [][]int{
	//	{1, 1},
	//	{1, 1},
	//}
	//expected := [][]int{
	//	{1, 1, 2, 2},
	//	{1, 1, 2, 2},
	//	{2, 2, 3, 3},
	//	{2, 2, 3, 3},
	//}

	//actual := extend(input, 2)

	//if len(expected) != len(actual) {
	//	t.Fatalf("expected len(actual) = %d but got %d", len(expected), len(actual))
	//}
	//if len(expected[0]) != len(actual[0]) {
	//	t.Fatalf("expected len(actual[0]) = %d but got %d", len(expected[0]), len(actual[0]))
	//}

	//tests := []struct {
	//	input      [][]int
	//	multiplier int
	//	want       [][]int
	//}{
	//	{
	//		input: [][]int{
	//			{1, 1},
	//			{1, 1},
	//		},
	//		multiplier: 2,
	//		want: [][]int{
	//			{1, 1, 2, 2},
	//			{1, 1, 2, 2},
	//			{2, 2, 3, 3},
	//			{2, 2, 3, 3},
	//		},
	//	},
	//	{
	//		input: [][]int{
	//			{8},
	//		},
	//		multiplier: 5,
	//		want: [][]int{
	//			{8, 9, 1, 2, 3},
	//			{9, 1, 2, 3, 4},
	//			{1, 2, 3, 4, 5},
	//			{2, 3, 4, 5, 6},
	//			{3, 4, 5, 6, 7},
	//		},
	//	},
	//}

	//for _, tc := range tests {
	//	got := extend(tc.input, tc.multiplier)
	//	if !reflect.DeepEqual(tc.want, got) {
	//		t.Fatalf("expected: %v, got: %v", tc.want, got)
	//	}
	//}

	S := NewStack()
	S.stack = []byte("ABC")
	AssertStackHasElements(t, S, "ABC")

	result := S.Pop()
	if result != "C"[0] {
		t.Fatalf("expected .Pop() to return C but got %c", result)
	}
	result = S.Pop()
	if result != "B"[0] {
		t.Fatalf("expected .Pop() to return B but got %c", result)
	}
	result = S.Pop()
	if result != "A"[0] {
		t.Fatalf("expected .Pop() to return A but got %c", result)
	}

	S = NewStack()
	S.stack = []byte("ABC")

	popNResult := S.PopN(1)
	if len(popNResult) != 1 || popNResult[0] != "C"[0] {
		t.Fatalf("expected .PopN() to return [C] but got %c", popNResult)
	}
	popNResult = S.PopN(1)
	if len(popNResult) != 1 || popNResult[0] != "B"[0] {
		t.Fatalf("expected .PopN() to return [B] but got %c", popNResult)
	}
	popNResult = S.PopN(1)
	if len(popNResult) != 1 || popNResult[0] != "A"[0] {
		t.Fatalf("expected .PopN() to return [A] but got %c", popNResult)
	}

	S = NewStack()
	S.stack = []byte("ABC")
	AssertStackHasElements(t, S, "ABC")

	popNResult = S.PopN(3)
	AssertSlicesEqual(t, popNResult, []byte("ABC"))
	AssertStackHasElements(t, S, "")

	S = NewStack()
	S.stack = []byte("ABC")
	AssertStackHasElements(t, S, "ABC")

	S.PushN([]byte("DEF"))
	AssertStackHasElements(t, S, "ABCDEF")

	S.PushN([]byte("DEF"))
	AssertStackHasElements(t, S, "ABCDEFDEF")
	S.PopN(2)
	AssertStackHasElements(t, S, "ABCDEFD")
	S.PopN(3)
	AssertStackHasElements(t, S, "ABCD")

	a := S.PopN(1)
	AssertStackHasElements(t, S, "ABC")
	S.PushN(a)
	AssertStackHasElements(t, S, "ABCD")

	S.PushN([]byte("EFGHIJKL"))
	AssertStackHasElements(t, S, "ABCDEFGHIJKL")
	a = S.PopN(12)
	AssertStackHasElements(t, S, "")
	S.PushN(a)
	AssertStackHasElements(t, S, "ABCDEFGHIJKL")

	A := NewStack()
	A.stack = []byte("WGBJTV")
	AssertStackHasElements(t, A, "WGBJTV")

	B := NewStack()
	B.stack = []byte("LGJMDNV")
	AssertStackHasElements(t, B, "LGJMDNV")

	a = A.PopN(3)
	AssertStackHasElements(t, A, "WGB")
	AssertSlicesEqual(t, a, []byte("JTV"))

	B.PushN(a)
	AssertStackHasElements(t, B, "LGJMDNVJTV")
}

func AssertStackHasElements(t *testing.T, s *Stack, expected string) {
	if len(s.stack) != len(expected) {
		t.Fatalf("expected stack to contain '%s' but it contains '%s'", expected, s.stack)
	}
	for i, v := range []byte(expected) {
		if s.stack[i] != v {
			t.Fatalf("expected s.stack[%d] == %c but got %c, %s %s\n", i, v, s.stack[i], s.stack, expected)
		}
	}
}

func AssertSlicesEqual(t *testing.T, actual, expected []byte) {
	if len(actual) != len(expected) {
		t.Fatalf("expected len(actual) == %d but got %d", len(expected), len(actual))
	}
	for i, v := range []byte(expected) {
		if actual[i] != v {
			t.Fatalf("expected %s but got %s", expected, actual)
		}
	}
}
