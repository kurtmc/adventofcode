package main

import (
	"fmt"
	"strconv"
)

type Part1Solver struct {
	x       int
	y       int
	riskMap [][]int
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		x:       0,
		y:       0,
		riskMap: make([][]int, 0),
	}
}

func (s *Part1Solver) Line(l string) {
	var row []int
	row = make([]int, 0)
	s.x = len(l)
	for _, v := range l {
		risk, err := strconv.Atoi(fmt.Sprintf("%c", v))
		if err != nil {
			panic(err)
		}

		row = append(row, risk)

	}
	s.riskMap = append(s.riskMap, row)

	s.y++
}

func (s *Part1Solver) End() string {
	fmt.Println("risk map:")
	fmt.Println(s.riskMap)
	printMatrix(s.riskMap)

	fmt.Printf("width: %d\n", s.x)
	fmt.Printf("height: %d\n", s.y)

	dist := make(map[Point]int, 0)
	prev := make(map[Point]Point, 0)
	Q := make([]Point, 0)

	for i := 0; i < s.x; i++ {
		for j := 0; j < s.y; j++ {
			v := Point{X: i, Y: j}
			dist[v] = 9999999
			// prev[v] := undefined
			Q = append(Q, Point{X: i, Y: j})
		}

	}

	// dist[source] ← 0
	dist[Point{X: 0, Y: 0}] = 0

	// while Q is not empty:
	for len(Q) != 0 {
		// u ← vertex in Q with min dist[u]
		u := minDistance(&Q, dist)

		n := neighborsAndInQ(u, Q, s.x, s.y)

		// for each neighbor v of u still in Q:
		for _, v := range n {
			// alt ← dist[u] + Graph.Edges(u, v)
			alt := dist[u] + s.riskMap[v.Y][v.X]
			// if alt < dist[v]:
			if (dist[v] == -1 && alt >= 0) || alt < dist[v] {
				// dist[v] ← alt
				dist[v] = alt
				prev[v] = u
				// prev[v] ← u
			}
		}

	}

	//  S ← empty sequence
	S := make([]Point, 0)
	//  u ← target
	u := Point{X: s.x - 1, Y: s.y - 1}
	//  if prev[u] is defined or u = source:          // Do something only if the vertex is reachable

	if _, ok := prev[u]; ok || (u.X == 0 && u.Y == 0) {
		//      while u is defined:                       // Construct the shortest path with a stack S
		for ok {
			//          insert u at the beginning of S        // Push the vertex onto the stack
			S = append([]Point{u}, S...)
			//          u ← prev[u]                           // Traverse from target to source
			u, ok = prev[u]
		}
	}

	printPath(s.riskMap, S)

	lowestTotalRisk := 0
	for _, ss := range S {
		if ss.X == 0 && ss.Y == 0 {
			continue
		}

		lowestTotalRisk = lowestTotalRisk + s.riskMap[ss.Y][ss.X]

	}

	return fmt.Sprintf("%d", lowestTotalRisk)
}

func isInList(u Point, Q []Point) bool {
	for _, v := range Q {
		if v.X == u.X && v.Y == u.Y {
			return true
		}
	}
	return false
}

func neighborsAndInQ(u Point, Q []Point, x, y int) []Point {
	result := make([]Point, 0)

	neighborCoordinates := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, coord := range neighborCoordinates {
		if u.X+coord[0] >= 0 && u.Y+coord[1] >= 0 && u.X+coord[0] < x && u.Y+coord[1] < y {
			v := Point{X: u.X + coord[0], Y: u.Y + coord[1]}
			if isInList(v, Q) {
				result = append(result, v)
			}
		}
	}

	return result
}

func printMatrix(m [][]int) {
	x := len(m[0])
	y := len(m)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			fmt.Printf("%v ", m[i][j])
		}
		fmt.Printf("\n")
	}
}

func printPath(riskMap [][]int, S []Point) {
	x := len(riskMap[0])
	y := len(riskMap)
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			isInPath := false
			for _, s := range S {
				if i == 0 && j == 0 {
					isInPath = true
				}
				if s.X == j && s.Y == i {
					isInPath = true
				}

			}
			if isInPath {
				fmt.Printf("\x1b[%dm%d\x1b[0m", 34, riskMap[i][j])
			} else {
				fmt.Printf("%v", riskMap[i][j])
			}
		}
		fmt.Printf("\n")
	}
}

// vertex in Q with min dist[u]
func minDistance(Q *[]Point, dist map[Point]int) Point {
	//fmt.Println("vertex in Q with min dist[u]")
	//fmt.Printf("Q: %v\n", *Q)

	min := -1 // -1 represents infinity
	index := 0
	for k, v := range *Q {
		d := dist[v]
		if d == -1 { // -1 represents infinity
			continue
		}
		if min == -1 {
			min = d
			index = k
		} else if min > d {
			min = d
			index = k
		}

	}

	u := (*Q)[index]
	*Q = append((*Q)[:index], (*Q)[index+1:]...)
	return u
}

type Point struct {
	X, Y int
}
