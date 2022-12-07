package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Part2Solver struct {
	diskSize          int
	freeSpaceRequired int
	currentDir        *Stack
	root              *File
}

func NewPart2Solver() *Part2Solver {
	return &Part2Solver{
		root:              NewFile(true, "root", 0),
		diskSize:          70000000,
		freeSpaceRequired: 30000000,
	}
}

func (s *Part2Solver) Line(l string) {

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

func (s *Part2Solver) End() string {
	CalcDirectorSizes(s.root)
	PrintFs(s.root, 0)

	freeSpace := s.diskSize - s.root.size
	required := s.freeSpaceRequired - freeSpace
	p := GetPotentialDirectories(s.root, required)
	var min *File
	for _, f := range p {
		if min == nil || f.size < min.size {
			min = f
		}
	}
	fmt.Printf("%s (%d)\n", min.name, min.size)

	return fmt.Sprintf("%d", min.size)
}

func GetPotentialDirectories(f *File, required int) []*File {
	result := make([]*File, 0)
	if !f.isDir {
		return result
	}

	if f.size > required {
		result = append(result, f)
	}
	for _, v := range f.contains {
		r := GetPotentialDirectories(v, required)
		if r != nil {
			result = append(result, r...)
		}
	}

	return result
}
