package main

import (
	"container/heap"
	"fmt"
)

type Part2Solver struct {
	x         int
	y         int
	heightmap [][]rune
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		x:         0,
		y:         0,
		heightmap: make([][]rune, 0),
	}
}

func (s *Part2Solver) Line(l string) {
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

func FindShortestPath(heightMap *HeightMap) int {
	fmt.Printf("heightMap: %v\n", heightMap)
	fmt.Printf("source point: %v\n", heightMap.SourcePoint)
	fmt.Printf("end point: %v\n", heightMap.EndPoint)

	dist := make(map[Point]int, 0)
	prev := make(map[Point]Point, 0)

	// dist[source] ← 0
	dist[*heightMap.SourcePoint] = 0

	Q := make(PriorityQueue, 0)

	var itemMap map[Point]*Item = make(map[Point]*Item, 0)

	for i := 0; i < len(heightMap.heightmap[0]); i++ {
		for j := 0; j < len(heightMap.heightmap); j++ {
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

	//  S ← empty sequence
	S := make([]Point, 0)
	//  u ← target
	u := *heightMap.EndPoint
	//  if prev[u] is defined or u = source:          // Do something only if the vertex is reachable

	if _, ok := prev[u]; ok || (u.X == heightMap.SourcePoint.X && u.Y == heightMap.SourcePoint.Y) {
		//      while u is defined:                       // Construct the shortest path with a stack S

		limit := 999
		count := 0
		for count < limit && ok {
			//          insert u at the beginning of S        // Push the vertex onto the stack
			S = append([]Point{u}, S...)
			//          u ← prev[u]                           // Traverse from target to source
			u, ok = prev[u]

			count++
		}
	}

	fmt.Printf("S: %v\n", S)
	fmt.Printf("prev: %v\n", prev)

	return len(S) - 1
}

func (s *Part2Solver) End() string {
	startingPositions := make([]Point, 0)
	for y := 0; y < len(s.heightmap); y++ {
		for x := 0; x < len(s.heightmap[0]); x++ {
			if s.heightmap[y][x] == 'S' || s.heightmap[y][x] == 'a' {
				startingPositions = append(startingPositions, Point{X: x, Y: y})
			}
		}
	}

	fmt.Printf("starting positions: %v\n", startingPositions)
	min := -1
	for _, startingPositions := range startingPositions {
		heightMap := NewHeightMap(s.heightmap)
		heightMap.SourcePoint = &startingPositions
		shortestPath := FindShortestPath(heightMap)
		fmt.Printf("shortest path: %d\n", shortestPath)

		if min == -1 || shortestPath < min {
			min = shortestPath
		}
	}
	return fmt.Sprintf("%d", min)
}
