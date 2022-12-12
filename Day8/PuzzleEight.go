package main

import (
	"2022/utils"
	"log"
	"strconv"
	"strings"
)

//var rotation = 0

type tree struct {
	Height      int
	Visible     bool
	ScenicScore int
}

type grid struct {
	Trees    [][]*tree
	Visible  int
	BestView int
}

func (g *grid) CountVisible() {
	for i, row := range g.Trees {
		for j, t := range row {
			if i == 0 || j == 0 || i == len(g.Trees)-1 || j == len(row)-1 {
				t.Visible = true
			}
		}
	}
	for i := 0; i < 4; i++ {
		g.count()
		g.rotate()
	}

	count := 0
	for _, row := range g.Trees {
		for _, t := range row {
			if t.Visible {
				count += 1
			}
		}
	}
	g.Visible = count
}

func (g *grid) count() {
	var highest int
	for i := 0; i < len(g.Trees)-1; i++ {
		highest = 0
		for j := 0; j < len(g.Trees[i])-1; j++ {
			currentHeight := g.Trees[i][j].Height
			if currentHeight > highest {
				g.Trees[i][j].Visible = true
				highest = g.Trees[i][j].Height
			}
		}
	}
}

func (g *grid) rotate() {
	for i, j := 0, len(g.Trees)-1; i < j; i, j = i+1, j-1 {
		g.Trees[i], g.Trees[j] = g.Trees[j], g.Trees[i]
	}
	for i := 0; i < len(g.Trees); i++ {
		for j := 0; j < i; j++ {
			g.Trees[i][j], g.Trees[j][i] = g.Trees[j][i], g.Trees[i][j]
		}
	}
	//g.print()
}

func (g *grid) SortViewCount() {
	for i := 0; i < 4; i++ {
		g.sortView()
		g.rotate()
	}

}

func (g *grid) sortView() {
	for _, row := range g.Trees {
		for i := 0; i < len(row); i++ {
			count := 0
			for j := i + 1; j < len(row); j++ {
				if row[i].Height > row[j].Height {
					count += 1
				}
				if row[i].Height <= row[j].Height {
					count += 1
					break
				}
			}
			row[i].ScenicScore *= count
		}
	}
	//rotation += 1
}

func (g *grid) FindBestView() {
	highest := 0
	for _, row := range g.Trees {
		for _, t := range row {
			if t.ScenicScore > highest {
				highest = t.ScenicScore
			}
		}
	}
	g.BestView = highest
}

func (g *grid) print() {
	var rows []string
	for _, row := range g.Trees {
		var tr []string
		for _, h := range row {
			tr = append(tr, "["+strconv.Itoa(h.Height)+","+strconv.Itoa(h.ScenicScore)+"]")
		}
		rows = append(rows, strings.Join(tr, ","))
	}
	log.Print("*******************")
	for _, row := range rows {
		log.Print(row)
	}
	log.Print("*******************")

}

func buildGrid(input []string) grid {
	var row grid
	for rowIdx, line := range input {
		row.Trees = append(row.Trees, []*tree{})
		for _, char := range line {
			charI, _ := strconv.Atoi(string(char))
			row.Trees[rowIdx] = append(row.Trees[rowIdx], &tree{
				Height:      charI,
				Visible:     false,
				ScenicScore: 1,
			})
		}
	}
	return row
}

func PuzzleEight() map[string]int {
	//input := utils.ReadInput(utils.GetWD() + "/Day8/data/test.txt")
	input := utils.ReadInput(utils.GetWD() + "/Day8/data/input.txt")
	treeGrid := buildGrid(input)
	treeGrid.CountVisible()
	treeGrid.SortViewCount()
	treeGrid.FindBestView()

	return map[string]int{"ans1": treeGrid.Visible, "ans2": treeGrid.BestView}
}

func main() {
	log.Printf("PuzzleEight(): %v", PuzzleEight())
}
