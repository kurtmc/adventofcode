package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part1Solver struct {
	currentDir *Stack
	root       *File
}

func NewPart1Solver() *Part1Solver {
	return &Part1Solver{
		root: NewFile(true, "root", 0),
	}
}

func (s *Part1Solver) Line(l string) {

	if strings.HasPrefix(l, "$ ") { // a command
		if strings.HasPrefix(l, "$ cd ") {
			fmt.Println(l)
			dir := strings.Split(l, " ")[2]
			if dir == "/" {
				s.currentDir = NewStack()
				s.currentDir.Push(s.root)
			} else if dir == ".." {
				s.currentDir.Pop()
			} else {
				for _, v := range s.currentDir.Peek().contains {
					if v.name == dir {
						s.currentDir.Push(v)
					}
				}

			}

		}
		if strings.HasPrefix(l, "$ ls") {
			fmt.Printf("listing %s:\n", FileSliceAsPath(s.currentDir.stack))
		}
	} else {
		parts := strings.Split(l, " ")
		var file *File
		if parts[0] == "dir" {
			file = NewFile(true, parts[1], 0)
		} else {
			size, _ := strconv.Atoi(parts[0])
			file = NewFile(false, parts[1], size)
		}
		//fmt.Println(s.currentDir)
		s.currentDir.Peek().contains = append(s.currentDir.Peek().contains, file)

		//fmt.Printf("current directory: /%s\n", strings.Join(s.currentDir.stack, "/"))
		fmt.Printf("file %v\n", file)
	}

}

func FileSliceAsPath(s []*File) string {
	result := "/"
	for _, v := range s {
		if v.name == "root" {
			continue
		}
		result = result + "/" + v.name
	}
	return result
}

func (s *Part1Solver) End() string {
	CalcDirectorSizes(s.root)
	PrintFs(s.root, 0)
	return fmt.Sprintf("%d", SumAtMost(s.root, 100000))
}

type File struct {
	isDir    bool
	name     string
	contains []*File
	size     int
}

func NewFile(isDir bool, name string, size int) *File {
	file := &File{
		isDir: isDir,
		name:  name,
		size:  size,
	}

	if file.isDir {
		file.contains = make([]*File, 0)
	}

	return file
}

type Stack struct {
	stack []*File
}

func NewStack() *Stack {
	return &Stack{
		stack: make([]*File, 0),
	}
}

func (s *Stack) Push(e *File) {
	s.stack = append(s.stack, e)
}

func (s *Stack) Pop() *File {
	element := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return element
}

func (s *Stack) Peek() *File {
	element := s.stack[len(s.stack)-1]
	return element
}

func PrintFs(f *File, indent int) {
	for i := 0; i < indent; i++ {
		fmt.Printf(" ")
	}
	if f.name == "root" {
		fmt.Printf("- / ")
	} else {
		fmt.Printf("- %s ", f.name)
	}
	if f.isDir {
		fmt.Printf("(dir, totalSize=%d)\n", f.size)
		for _, v := range f.contains {
			PrintFs(v, indent+2)
		}
	} else {
		fmt.Printf("(file, size=%d)\n", f.size)
	}

}

func CalcDirectorSizes(f *File) int {
	if !f.isDir {
		return f.size
	}
	totalSize := 0
	for _, v := range f.contains {
		totalSize += CalcDirectorSizes(v)
	}
	f.size = totalSize
	return totalSize
}

func SumAtMost(f *File, limit int) int {
	if !f.isDir {
		return 0
	}

	sum := 0
	if f.size < limit {
		sum = f.size
	}
	for _, v := range f.contains {
		sum += SumAtMost(v, limit)
	}

	return sum
}
