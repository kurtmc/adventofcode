package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

var x, y int = 0, 0

// part 1
func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var riskMap [][]int = make([][]int, 0)

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		var row []int
		row = make([]int, 0)
		x = len(text)
		for _, v := range text {
			risk, err := strconv.Atoi(fmt.Sprintf("%c", v))
			if err != nil {
				panic(err)
			}

			row = append(row, risk)

		}
		riskMap = append(riskMap, row)

		y++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("risk map:")
	fmt.Println(riskMap)
	printMatrix(riskMap)

	fmt.Printf("width: %d\n", x)
	fmt.Printf("height: %d\n", y)

	dist := make(map[Point]int, 0)
	prev := make(map[Point]Point, 0)
	Q := make([]Point, 0)

	for i := 0; i < x; i++ {
		for j := 0; j < y; j++ {
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
		fmt.Println("u := minDistance(&Q, dist)")
		fmt.Printf("u := %v\n", u)
		//Q = append(Q[:minIndex], Q[minIndex+1:]...)

		n := neighbors(u, Q)

		// for each neighbor v of u still in Q:
		for _, v := range n {
			// alt ← dist[u] + Graph.Edges(u, v)
			alt := dist[u] + riskMap[v.Y][v.X]
			fmt.Printf("alt: %d\n", alt)
			fmt.Printf("dist[%v]: %d\n", v, dist[v])
			// if alt < dist[v]:
			if (dist[v] == -1 && alt >= 0) || alt < dist[v] {
				// dist[v] ← alt
				dist[v] = alt
				prev[v] = u
				fmt.Printf("prev[%v] = %v\n", v, u)
				// prev[v] ← u
			}
		}

	}

	//  S ← empty sequence
	S := make([]Point, 0)
	//  u ← target
	u := Point{X: x - 1, Y: y - 1}
	//  if prev[u] is defined or u = source:          // Do something only if the vertex is reachable

	fmt.Printf("prev: %v\n", prev)

	fmt.Printf("u: %v\n", u)
	fmt.Printf("prev[%v]: %v\n", u, prev[u])
	if _, ok := prev[u]; ok || (u.X == 0 && u.Y == 0) {
		fmt.Println("####################")
		fmt.Println("####################")
		fmt.Println("####################")
		//      while u is defined:                       // Construct the shortest path with a stack S
		for ok {
			//          insert u at the beginning of S        // Push the vertex onto the stack
			S = append([]Point{u}, S...)
			//          u ← prev[u]                           // Traverse from target to source
			u, ok = prev[u]
		}
	}

	//fmt.Println(dist)
	//fmt.Println(len(prev))
	fmt.Printf("S: %v\n", S)

	printPath(riskMap, S)

	lowestTotalRisk := 0
	for _, s := range S {
		if s.X == 0 && s.Y == 0 {
			continue
		}

		lowestTotalRisk = lowestTotalRisk + riskMap[s.Y][s.X]

	}

	fmt.Printf("lowest total risk: %d\n", lowestTotalRisk)

}

func create3dSlice(x, y int, initialValue []int) [][][]int {
	result := make([][][]int, y)
	for i := range result {
		result[i] = make([][]int, x)

		for j := range result[i] {
			result[i][j] = initialValue
		}
	}

	return result
}

func isInList(u Point, Q []Point) bool {
	for _, v := range Q {
		if v.X == u.X && v.Y == u.Y {
			return true
		}
	}
	return false
}

func neighbors(u Point, Q []Point) []Point {
	fmt.Printf("find neighbors for u (%v)\n", u)

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

	fmt.Printf("neighbors: %v\n", result)

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
	fmt.Println("vertex in Q with min dist[u]")
	fmt.Printf("Q: %v\n", *Q)

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
	fmt.Printf("Q[minIndex]: %v\n", (*Q)[index])
	*Q = append((*Q)[:index], (*Q)[index+1:]...)
	return u
}

type Point struct {
	X, Y int
}
