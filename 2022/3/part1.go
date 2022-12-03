package main

import (
	"fmt"
	"unicode"
)

type Part1Solver struct {
	prioritySum int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{}
}

func (s *Part1Solver) Line(l string) {
	left := l[:len(l)/2]
	right := l[len(l)/2:]
	fmt.Println(left)
	fmt.Println(right)

	commonLetter := getCommonLetters(left, right)[0]
	fmt.Printf("common letters: %c, %v\n", commonLetter, getPriority(commonLetter))

	s.prioritySum = s.prioritySum + getPriority(commonLetter)
}

func (s *Part1Solver) End() string {
	return fmt.Sprintf("%d", s.prioritySum)
}

func getCommonLetters(a, b string) []rune {
	h := make(map[rune]int, 0)

	for _, v := range a {
		h[v] = 1
	}

	for _, v := range b {
		if _, ok := h[v]; ok {
			h[v] = 2
		}
	}

	result := make([]rune, 0)
	for r, v := range h {
		if v == 2 {
			result = append(result, r)
		}
	}

	return result
}

func getPriority(r rune) int {
	if IsLower(r) {
		return int(r) - 96
	} else {
		return int(r) - 38
	}
}

func IsUpper(r rune) bool {
	if !unicode.IsUpper(r) && unicode.IsLetter(r) {
		return false
	}
	return true
}

func IsLower(r rune) bool {
	if !unicode.IsLower(r) && unicode.IsLetter(r) {
		return false
	}
	return true
}
