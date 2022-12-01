package main

import (
	"fmt"
	"strconv"
)

type Part2Solver struct {
	leaderboard *Leaderboard
	previous    int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		leaderboard: NewLeaderboard(3),
	}
}

type Leaderboard struct {
	size  int
	items []int
}

func NewLeaderboard(size int) *Leaderboard {
	return &Leaderboard{
		size:  size,
		items: make([]int, 0),
	}
}

func (l *Leaderboard) CheckAndAdd(value int) {
	if len(l.items) < l.size {
		l.items = append(l.items, value)
		return
	}

	var min, minIndex int = 0, 0
	for i, v := range l.items {
		if min == 0 || v < min {
			min = v
			minIndex = i
		}
	}

	if value > min {
		l.items[minIndex] = value
	}
}

func (s *Part2Solver) Line(l string) {
	if l == "" {
		s.leaderboard.CheckAndAdd(s.previous)
		s.previous = 0
	} else {
		i, err := strconv.Atoi(l)
		if err != nil {
			panic(err)
		}
		s.previous = s.previous + i
	}
}

func (s *Part2Solver) End() string {
	// Need to run line with "" input to finalize the results
	s.Line("")

	result := 0
	for _, v := range s.leaderboard.items {
		result = result + v
	}
	return fmt.Sprintf("%d", result)
}
