package main

import (
	"container/heap"
	"fmt"
	"strconv"
)

type Part2Solver struct {
	x       int
	y       int
	riskMap [][]int
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		x:       0,
		y:       0,
		riskMap: make([][]int, 0),
	}
}

func (s *Part2Solver) Line(l string) {
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

func (s *Part2Solver) End() string {
	newRiskMap := extend(s.riskMap, 5)
	s.riskMap = newRiskMap
	s.x = s.x * 5
	s.y = s.y * 5

	fmt.Printf("width: %d\n", s.x)
	fmt.Printf("height: %d\n", s.y)

	dist := make(map[Point]int, 0)
	prev := make(map[Point]Point, 0)

	// dist[source] ← 0
	dist[Point{X: 0, Y: 0}] = 0

	Q := make(PriorityQueue, 0)

	var itemMap map[Point]*Item = make(map[Point]*Item, 0)

	for i := 0; i < s.x; i++ {
		for j := 0; j < s.y; j++ {
			v := Point{X: i, Y: j}
			// if v ≠ source
			if !(v.X == 0 && v.Y == 0) {
				dist[v] = -1
			}
			item := &Item{
				value:    v,
				priority: dist[v],
			}
			itemMap[v] = item
			heap.Push(&Q, item)
		}

	}

	// while Q is not empty:
	for Q.Len() != 0 {
		// u ← Q.extract_min()
		item := heap.Pop(&Q).(*Item)
		u := item.value

		n := neighbors(u, s.x, s.y)

		// for each neighbor v of u:
		for _, v := range n {
			// alt ← dist[u] + Graph.Edges(u, v)
			alt := dist[u] + s.riskMap[v.Y][v.X]
			// if alt < dist[v]:
			if (dist[v] == -1 && alt >= 0) || alt < dist[v] {
				// dist[v] ← alt
				dist[v] = alt
				prev[v] = u
				Q.update(itemMap[v], itemMap[v].value, alt)

				// prev[v] ← u
			}
		}

	}

	fmt.Println("##################")

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

func neighbors(u Point, x, y int) []Point {
	result := make([]Point, 0)

	neighborCoordinates := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	for _, coord := range neighborCoordinates {
		if u.X+coord[0] >= 0 && u.Y+coord[1] >= 0 && u.X+coord[0] < x && u.Y+coord[1] < y {
			v := Point{X: u.X + coord[0], Y: u.Y + coord[1]}
			result = append(result, v)
		}
	}

	return result
}

func extend(matrix [][]int, multiplier int) [][]int {
	y := len(matrix)
	x := len(matrix[0])

	// create result matrix of the new size
	var result [][]int = make([][]int, y*multiplier)
	for n := 0; n < multiplier*y; n++ {
		result[n] = make([]int, x*multiplier)
	}

	// put input matrix into top left of result matrix
	for i := 0; i < y; i++ {
		for j := 0; j < x; j++ {
			result[i][j] = matrix[i][j]
		}
	}

	for m := 0; m < multiplier; m++ {
		for n := 0; n < multiplier; n++ {
			// skip 0, 0
			if m == 0 && n == 0 {
				continue
			}

			for i := 0; i < y; i++ {
				for j := 0; j < x; j++ {
					newRisk := 0
					if j+(n*x)-x < 0 {
						newRisk = (result[i+(m-1)*y][j+(n*x)] + 1) % 10
					} else {
						newRisk = (result[i+(m*y)][j+(n-1)*x] + 1) % 10
					}
					if newRisk == 0 {
						newRisk = 1
					}
					result[i+(m*y)][j+(n*x)] = newRisk
				}
			}
		}
	}

	return result
}

type Item struct {
	value    Point
	priority int
	index    int
}

type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	if pq[i].priority == -1 && pq[j].priority == -1 {
		return true
	} else if pq[i].priority == -1 {
		return false
	} else if pq[j].priority == -1 {
		return true
	}
	// we want lowest priority, so we use less and -1 represents infinity
	return pq[i].priority < pq[j].priority
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
	pq[i].index = i
	pq[j].index = j
}

func (pq *PriorityQueue) Push(x any) {
	n := len(*pq)
	item := x.(*Item)
	item.index = n
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*pq = old[0 : n-1]
	return item
}

func (pq *PriorityQueue) update(item *Item, value Point, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(pq, item.index)
}
