package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	solver := NewPart2Solver()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		solver.Line(text)
	}
	if err := scanner.Err(); err != nil {
		if err != io.EOF {
			log.Fatal(err)
		}
	}

	fmt.Printf("result: %s\n", solver.End())
}
