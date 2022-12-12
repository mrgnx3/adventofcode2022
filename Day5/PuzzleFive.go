package main

import (
	"log"
	"os"
	"regexp"
	"strconv"
	"strings"
)

/*
puzzle input:
Starting stacks + desired results
*/

func sanitize(file []byte) []string {
	layoutstr := strings.Split(string(file), "\n\n")[0]
	layoutstr = strings.Replace(layoutstr, "    ", "0", -1)
	layoutstr = strings.Replace(layoutstr, " ", "", -1)
	layoutstr = strings.Replace(layoutstr, "[", "", -1)
	layoutstr = strings.Replace(layoutstr, "]", "", -1)
	return strings.Split(layoutstr, "\n")
}

type direction struct {
	Original, Dest, Amount int
}

func generateDirection(line string) direction {
	re := regexp.MustCompile("[0-9]+")
	listofnums := re.FindAllString(line, -1)
	a, _ := strconv.Atoi(listofnums[0])
	o, _ := strconv.Atoi(listofnums[1])
	d, _ := strconv.Atoi(listofnums[2])
	return direction{
		Amount:   a,
		Original: o - 1,
		Dest:     d - 1,
	}
}

func readDirections(file []byte) []direction {
	var directions []direction
	directionsStr := strings.Split(string(file), "\n\n")[1]
	for _, line := range strings.Split(directionsStr, "\n") {
		if line != "" {
			directions = append(directions, generateDirection(line))
		}
	}
	return directions
}

func buildLayout(layoutLst []string) map[int][]string {
	lo := map[int][]string{}

	for i := len(layoutLst) - 1; i >= 0; i-- {
		for j := 0; j < len(layoutLst[i]); j++ {
			if string(layoutLst[i][j]) != "0" {
				lo[j] = append(lo[j], string(layoutLst[i][j]))
			}
		}
	}
	return lo
}

//move 1 from 2 to 1

func executeOrder(layout map[int][]string, d direction, secondPuzzle bool) map[int][]string {
	if secondPuzzle {
		for idx := d.Amount - 1; idx >= 0; idx-- {
			layout[d.Dest] = append(layout[d.Dest], layout[d.Original][len(layout[d.Original])-idx-1])
		}
		for idx := 0; idx < d.Amount; idx++ {
			layout[d.Original] = layout[d.Original][:len(layout[d.Original])-1]
		}

	} else {
		for idx := 0; idx < d.Amount; idx++ {
			layout[d.Dest] = append(layout[d.Dest], layout[d.Original][len(layout[d.Original])-1])
			layout[d.Original] = layout[d.Original][:len(layout[d.Original])-1]
		}
	}
	return layout
}

func reorder(layout map[int][]string, directions []direction, secondPuzzle bool) map[int][]string {
	newLayout := map[int][]string{}
	for _, direction := range directions {
		newLayout = executeOrder(layout, direction, secondPuzzle)
	}
	return newLayout
}

func displayResult(layout map[int][]string) string {
	ans := []string{}
	for i := 0; i < len(layout); i++ {
		ans = append(ans, "0")
	}

	for k, _ := range layout {
		ans[k] = layout[k][len(layout[k])-1]
	}
	return strings.Join(ans, "")
}

func PuzzleFive() map[string]string {
	//file, err := os.ReadFile("data/test.txt")
	file, err := os.ReadFile("data/input.txt")
	if err != nil {
		log.Printf("error reading file: %d", err)
	}
	layoutLst := sanitize(file)
	layout1 := buildLayout(layoutLst)
	layout2 := buildLayout(layoutLst)
	directions := readDirections(file)

	layout1 = reorder(layout1, directions, false)
	layout2 = reorder(layout2, directions, true)
	ans1 := displayResult(layout1)
	ans2 := displayResult(layout2)
	return map[string]string{"roundOne": ans1, "roundTwo": ans2}
}

func main() {
	log.Printf("PuzzleFive(): %v", PuzzleFive())
}
