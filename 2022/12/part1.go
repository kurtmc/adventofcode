package main

import (
	"container/heap"
	"fmt"
)

type Part1Solver struct {
	x         int
	y         int
	heightmap [][]rune
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		x:         0,
		y:         0,
		heightmap: make([][]rune, 0),
	}
}

type HeightMap struct {
	SourcePoint *Point
	EndPoint    *Point
	heightmap   [][]rune
}

func NewHeightMap(heightMap [][]rune) *HeightMap {
	result := &HeightMap{
		heightmap: heightMap,
	}

	for y := 0; y < len(heightMap); y++ {
		for x := 0; x < len(heightMap[0]); x++ {
			if heightMap[y][x] == 'S' {
				result.SourcePoint = &Point{X: x, Y: y}
			}
			if heightMap[y][x] == 'E' {
				result.EndPoint = &Point{X: x, Y: y}
			}
		}
	}

	return result
}

func (h *HeightMap) GetHeight(p Point) int {
	height := h.heightmap[p.Y][p.X]

	if height == 'S' {
		return 0
	}
	if height == 'E' {
		return int('z' - 'a')
	}

	return int(height - 'a')
}

func (h *HeightMap) Neighbors(p Point) []Point {
	fmt.Printf("get neighbors for: %v\n", p)
	maxX := len(h.heightmap[0])
	maxY := len(h.heightmap)
	result := make([]Point, 0)

	neighborCoordinates := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}

	pointHeight := h.GetHeight(p)
	fmt.Printf("point height: %d\n", pointHeight)

	for _, coord := range neighborCoordinates {
		if p.X+coord[0] >= 0 && p.Y+coord[1] >= 0 && p.X+coord[0] < maxX && p.Y+coord[1] < maxY {
			v := Point{X: p.X + coord[0], Y: p.Y + coord[1]}
			nextPointHeight := h.GetHeight(v)
			fmt.Printf("next point height: %d\n", nextPointHeight)
			if nextPointHeight-1 <= pointHeight {
				result = append(result, v)
			}
		}
	}

	fmt.Printf("result: %v\n", result)

	return result
}

func (s *Part1Solver) Line(l string) {
	var row []rune
	row = make([]rune, 0)
	s.x = len(l)
	for _, v := range l {
		height := v
		row = append(row, height)

	}
	s.heightmap = append(s.heightmap, row)
	s.y++
}

func (s *Part1Solver) End() string {
	fmt.Printf("width: %d\n", s.x)
	fmt.Printf("height: %d\n", s.y)

	heightMap := NewHeightMap(s.heightmap)

	fmt.Printf("heightMap: %v\n", heightMap)
	fmt.Printf("source point: %v\n", heightMap.SourcePoint)
	fmt.Printf("end point: %v\n", heightMap.EndPoint)

	dist := make(map[Point]int, 0)
	prev := make(map[Point]Point, 0)

	// dist[source] ← 0
	dist[*heightMap.SourcePoint] = 0

	Q := make(PriorityQueue, 0)

	var itemMap map[Point]*Item = make(map[Point]*Item, 0)

	for i := 0; i < s.x; i++ {
		for j := 0; j < s.y; j++ {
			v := Point{X: i, Y: j}
			// if v ≠ source
			if !(v.X == heightMap.SourcePoint.X && v.Y == heightMap.SourcePoint.Y) {
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

		n := heightMap.Neighbors(u)

		// for each neighbor v of u:
		for _, v := range n {
			// alt ← dist[u] + Graph.Edges(u, v)
			alt := dist[u] + 1
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
	u := *heightMap.EndPoint
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

	fmt.Printf("S: %v\n", S)
	fmt.Printf("prev: %v\n", prev)

	printPath(s.heightmap, S)

	return fmt.Sprintf("%d", len(S)-1)
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

type Point struct {
	X, Y int
}

func printPath(heightMap [][]rune, S []Point) {
	x := len(heightMap[0])
	y := len(heightMap)
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
				fmt.Printf("\x1b[%dm%c\x1b[0m", 34, heightMap[i][j])
			} else {
				fmt.Printf("%c", heightMap[i][j])
			}
		}
		fmt.Printf("\n")
	}
}
