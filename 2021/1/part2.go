package main

import (
	"fmt"
	"strconv"
)

type Part2Solver struct {
	sw             *SlidingWindow
	increasesCount int
	index          int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		sw:             NewSlidingWindow(3),
		increasesCount: 0,
		index:          0,
	}
}

func (s *Part2Solver) Line(l string) {
	i, err := strconv.Atoi(l)
	if err != nil {
		panic(err)
	}

	previousSwSum := s.sw.Sum()

	s.sw.Put(i)

	swSum := s.sw.Sum()

	if s.index >= s.sw.size && swSum > previousSwSum {
		s.increasesCount++
	}

	s.index++
}

func (s *Part2Solver) End() string {
	return fmt.Sprintf("%d", s.increasesCount)
}

type SlidingWindow struct {
	size     int
	index    int
	elements []int
}

func NewSlidingWindow(size int) *SlidingWindow {
	return &SlidingWindow{
		size:     size,
		index:    0,
		elements: []int{},
	}
}

func (s *SlidingWindow) Put(x int) {
	if s.index < s.size {
		s.elements = append(s.elements, x)
	} else {
		s.elements[s.index%s.size] = x
	}
	s.index++
}

func (s *SlidingWindow) Sum() int {
	sum := 0
	for _, v := range s.elements {
		sum = sum + v
	}
	return sum
}
