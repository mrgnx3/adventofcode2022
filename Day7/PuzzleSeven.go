package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
puzzle input:
fs navigation

lines that begin with $ is commands you executed
/ is root
cd means the same as always
ls means the same
123 abc means the current dir contains a file abc sized 123
dir xyz contains a dir named xyz

read file
parse values
map fs
sum values of each dir
find every dir with a sum of less than 100000
sum of every dir with less than 100000

*/

const (
	Root             = "/"
	MaxFolderSize    = 100000
	TotalDiskSpace   = 70000000
	MinRequiredSpace = 30000000
)

type file struct {
	Name string
	Size int
}

type dir struct {
	Name     string
	Parent   *dir
	Children []*dir
	Files    []file
	Size     int
}

func (d *dir) populate(dirName string) {
	newChild := true
	for _, child := range d.Children {
		if child.Name == dirName {
			newChild = false
			break
		}
	}
	if newChild {
		d.Children = append(d.Children, &dir{
			Name:   dirName,
			Parent: d,
		})
	}

}

func (d *dir) populateFile(command string) {
	size, _ := strconv.Atoi(strings.Split(command, " ")[0])
	filename := strings.Split(command, " ")[1]
	if !contains(d.Files, filename) {
		d.Files = append(d.Files, file{
			Name: filename,
			Size: size,
		})
	}
}

func (d dir) findRoot() *dir {
	root := &d
	for root.Parent != nil {
		root = root.Parent
	}
	return root
}

func (d *dir) Contains(newDir string) bool {
	for _, child := range d.Children {
		if child.Name == newDir {
			return true
		}
	}
	return false
}

func (d *dir) get(newDir string) *dir {
	for _, child := range d.Children {
		if child.Name == newDir {
			return child
		}
	}
	return &dir{}
}

func contains(files []file, filename string) bool {
	for _, file := range files {
		if filename == file.Name {
			return true
		}
	}
	return false
}

func parseInput(command string, currentDir *dir) *dir {
	switch {
	case strings.HasPrefix(command, "$"):
		currentDir = execute(strings.TrimPrefix(command, "$ "), currentDir)
	case strings.HasPrefix(command, "dir"):
		currentDir.populate(strings.TrimPrefix(command, "dir "))
	default:
		currentDir.populateFile(command)
	}
	return currentDir
}

func execute(command string, currentDir *dir) *dir {
	switch {
	case strings.HasPrefix(command, "cd"):
		strippedCommand := strings.TrimPrefix(command, "cd ")
		if strippedCommand == "/" {
			if currentDir != nil {
				return currentDir.findRoot()
			} else {
				return &dir{
					Name: Root,
				}
			}
		}
		newDir := cd(strippedCommand, currentDir)
		if strippedCommand == ".." {
			return currentDir.Parent
		} else {
			if currentDir.Contains(newDir) {
				return currentDir.get(newDir)
			}
		}

	case strings.HasPrefix(command, "ls"):
		//		do nothing
	}
	return currentDir
}

func cd(command string, currentDir *dir) string {
	switch command {
	case "..":
		if currentDir.Name != Root {
			return currentDir.Parent.Name
		} else {
			return currentDir.Name
		}
	default:
		return command
	}
}

func sum(array []int) int {
	result := 0
	for _, v := range array {
		result += v
	}
	return result
}

func (d *dir) calculateSize() {
	childTotalSize := 0
	for _, directory := range d.Children {
		directory.calculateSize()
		childTotalSize += directory.Size
	}
	d.Size = d.getFileSize() + childTotalSize
}

func (d *dir) getFileSize() int {
	total := 0
	for _, file := range d.Files {
		total += file.Size
	}
	return total
}

func (d *dir) findDirs() int {
	totalSize := 0
	for _, directory := range d.Children {
		totalSize += directory.findDirs()
	}
	if d.Size <= MaxFolderSize {
		totalSize += d.Size
	}
	return totalSize
}

func (d *dir) getSmallestSufficientDirectory(minSize, requiredSize int) int {
	if d.Size >= requiredSize && d.Size < minSize {
		minSize = d.Size
	}
	for _, directory := range d.Children {
		newMin := directory.getSmallestSufficientDirectory(minSize, requiredSize)
		if newMin >= requiredSize && newMin < minSize {
			minSize = newMin
		}
	}
	return minSize
}

func PuzzleSeven() map[string]int {
	datafile, err := os.Open("data/input.txt")
	//datafile, err := os.Open("data/test.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}
	fileScanner := bufio.NewScanner(datafile)
	defer func(file *os.File) {
		err := file.Close()
		if err != nil {

		}
	}(datafile)

	var root *dir

	for fileScanner.Scan() {
		root = parseInput(fileScanner.Text(), root)
	}
	log.Printf("currentDir: %v", root)
	root = root.findRoot()
	root.calculateSize()
	ans1 := root.findDirs()
	usedSpace := TotalDiskSpace - root.Size
	ans2 := root.getSmallestSufficientDirectory(TotalDiskSpace, MinRequiredSpace-usedSpace)
	return map[string]int{"ans1": ans1, "ans2": ans2}
}

func main() {
	log.Printf("PuzzleSeven(): %v", PuzzleSeven())
}

//6999588
