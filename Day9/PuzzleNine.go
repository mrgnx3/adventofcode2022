package main

import (
	"2022/utils"
	"fmt"
	"golang.org/x/exp/constraints"
	"log"
)

/*
Puzzle input
Hypothetical series of motions

H = Head
T = Tail

find head
find tail
move head 1
move tail 1 to follow
*/

type position struct {
	x, y int
}

type Set[T comparable] map[T]struct{}

func BuildSet[T comparable]() Set[T] {
	return make(Set[T])
}

func (s Set[T]) Add(value T) {
	s[value] = struct{}{}
}

func Abs[T constraints.Integer | constraints.Float](x T) T {
	if x < 0 {
		return -x
	}
	return x
}

type State struct {
	n    int
	rope []position
	path Set[position]
}

func NewState(n int) *State {
	res := &State{
		n:    n,
		rope: make([]position, n),
		path: BuildSet[position](),
	}
	res.path.Add(res.rope[n-1])
	return res
}

func (s *State) Move(dir byte) {
	switch dir {
	case 'U':
		s.rope[0].y++
	case 'D':
		s.rope[0].y--
	case 'L':
		s.rope[0].x--
	case 'R':
		s.rope[0].x++
	}
}

func (s *State) MoveTail() {
	for i := 1; i < s.n; i++ {
		delta := position{s.rope[i-1].x - s.rope[i].x, s.rope[i-1].y - s.rope[i].y}
		if Abs(delta.x) <= 1 && Abs(delta.y) <= 1 {
			return
		}
		if delta.y > 0 {
			s.rope[i].y++
		} else if delta.y < 0 {
			s.rope[i].y--
		}
		if delta.x > 0 {
			s.rope[i].x++
		} else if delta.x < 0 {
			s.rope[i].x--
		}
	}
	s.path.Add(s.rope[s.n-1])
}

func run(input []string, n int) int {
	state := NewState(n)
	for _, line := range input {
		var dir byte
		var nb int
		fmt.Sscanf(line, "%c %d", &dir, &nb)
		for i := 0; i < nb; i++ {
			state.Move(dir)
			state.MoveTail()
		}
	}
	return len(state.path)
}

func PuzzleNine() map[string]int {
	//input := utils.ReadInput(utils.GetWD() + "/Day9/data/test.txt")
	input := utils.ReadInput(utils.GetWD() + "/Day9/data/input.txt")

	ans1 := run(input, 2)
	ans2 := run(input, 10)

	return map[string]int{"ans1": ans1, "ans2": ans2}
}

func main() {
	log.Printf("PuzzleNine(): %v", PuzzleNine())
}
