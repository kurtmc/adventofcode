package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	previous := -1
	increasesCount := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		i, err := strconv.Atoi(text)
		if err != nil {
			panic(err)
		}

		if previous != -1 {
			if i > previous {
				increasesCount++
			}
		}

		previous = i
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("count: %d\n", increasesCount)
}
