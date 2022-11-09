package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

// part 1
// func main() {
// 	file, err := os.Open("input-part-1.txt")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer file.Close()
//
// 	horizontal := 0
// 	depth := 0
//
// 	scanner := bufio.NewScanner(file)
// 	for scanner.Scan() {
// 		text := scanner.Text()
// 		if text == "" {
// 			break
// 		}
//
// 		parts := strings.Split(text, " ")
//
// 		action := parts[0]
// 		value, err := strconv.Atoi(parts[1])
// 		if err != nil {
// 			panic(err)
// 		}
//
// 		if action == "forward" {
// 			horizontal = horizontal + value
// 		} else if action == "down" {
// 			depth = depth + value
// 		} else if action == "up" {
// 			depth = depth - value
// 		} else {
// 			panic(fmt.Sprintf("unknown action: %s", action))
// 		}
//
// 	}
//
// 	if err := scanner.Err(); err != nil {
// 		log.Fatal(err)
// 	}
//
// 	fmt.Printf("result: %d\n", horizontal*depth)
// }

// part 2
func main() {
	file, err := os.Open("input-part-2.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	horizontal := 0
	depth := 0
	aim := 0
	step := 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}

		parts := strings.Split(text, " ")

		action := parts[0]
		value, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}

		if action == "forward" {
			horizontal = horizontal + value
			depth = depth + aim*value
		} else if action == "down" {
			aim = aim + value
		} else if action == "up" {
			aim = aim - value
		} else {
			panic(fmt.Sprintf("unknown action: %s", action))
		}

		step++

	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %d\n", horizontal*depth)
}
