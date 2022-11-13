package main

import (
	"errors"
	"fmt"
	"os"
)

//go:generate go run gen.go

var mainGo string = `package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	solver := NewPart1Solver()

	file, err := os.Open("input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		text := scanner.Text()
		if text == "" {
			break
		}
		solver.Line(text)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("result: %s\n", solver.End())
}`
var partGo string = `package main

type Part%[1]dSolver struct {
}

func NewPart%[1]dSolver() *Part%[1]dSolver {
	return &Part%[1]dSolver{}
}

func (s *Part%[1]dSolver) Line(l string) {
}

func (s *Part%[1]dSolver) End() string {
	return ""
}`

var years []string = []string{"2021"}

func main() {
	fmt.Println("go generate!")

	for _, year := range years {
		for i := 1; i <= 25; i++ {

			dir := fmt.Sprintf("%s/%d", year, i)
			os.MkdirAll(dir, os.ModePerm)

			// go.mod
			path := dir + "/go.mod"
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				goMod := []byte(fmt.Sprintf(`module github.com/kurtmc/adventofcode/%s/%d

go 1.19
`, year, i))
				err := os.WriteFile(path, goMod, 0644)
				if err != nil {
					panic(err)
				}
			}
			// main.go
			path = dir + "/main.go"
			if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
				err := os.WriteFile(path, []byte(mainGo), 0644)
				if err != nil {
					panic(err)
				}
			}

			for i := 1; i <= 2; i++ {
				path = fmt.Sprintf("%s/part%d.go", dir, i)
				if _, err := os.Stat(path); errors.Is(err, os.ErrNotExist) {
					err := os.WriteFile(path, []byte(fmt.Sprintf(partGo, i)), 0644)
					if err != nil {
						panic(err)
					}
				}
			}
		}
	}
}
