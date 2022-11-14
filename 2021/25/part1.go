package main

import (
	"fmt"
	"reflect"
)

type Part1Solver struct {
	theMap [][]byte
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		theMap: make([][]byte, 0),
	}
}

func (s *Part1Solver) Line(l string) {
	s.theMap = append(s.theMap, []byte(l))
}

func (s *Part1Solver) End() string {

	count := 0
	for {
		n := step(s.theMap)
		count++
		if reflect.DeepEqual(n, s.theMap) {
			return fmt.Sprintf("%d", count)
		}
		s.theMap = n
	}
}

func printMap(m [][]byte) {
	r := "\n"
	for _, row := range m {
		r = r + fmt.Sprintf("%s\n", row)
	}

	fmt.Println(r)
}

func dup(m [][]byte) [][]byte {
	result := make([][]byte, len(m))

	i := 0
	for _, v := range m {
		result[i] = make([]byte, len(v))
		copy(result[i], v)
		i++
	}

	return result
}

func step(m [][]byte) [][]byte {
	X := len(m[0])
	Y := len(m)
	result := dup(m)

	// phase 1
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if m[y][x] == '>' && m[y][(x+1)%X] == '.' {
				result[y][x] = '.'
				result[y][(x+1)%X] = '>'
			}
		}
	}

	m = dup(result)

	// phase 2
	for y := 0; y < Y; y++ {
		for x := 0; x < X; x++ {
			if m[y][x] == 'v' && m[(y+1)%Y][x] == '.' {
				result[y][x] = '.'
				result[(y+1)%Y][x] = 'v'
			}
		}
	}

	return dup(result)
}
