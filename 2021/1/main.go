package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

// part 1
// func main() {
// 	file, err := os.Open("input.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	previous := -1
// 	increasesCount := 0
//
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		if text == "" {
// 			break
// 		}
//
// 		i, err := strconv.Atoi(text)
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		if previous != -1 {
// 			if i > previous {
// 				increasesCount++
// 			}
// 		}
//
// 		previous = i
// 	}
//
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Printf("count: %d\n", increasesCount)
// }

// part 2
func main() {
	file, err := os.Open("input-part-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	slidingWindowSize := 3
	sw := NewSlidingWindow(slidingWindowSize)

	increasesCount := 0

	scanner := bufio.NewScanner(file)
	index := 0
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		i, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		previousSwSum := sw.Sum()

		sw.Put(i)

		swSum := sw.Sum()

		if index >= slidingWindowSize && swSum > previousSwSum {
			increasesCount++
		}

		index++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("count: %d\n", increasesCount)
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
