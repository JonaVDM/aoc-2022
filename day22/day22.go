package day22

import (
	"fmt"
	"strconv"

	"github.com/jonavdm/aoc-2022/utils"
)

type Solver struct {
	Grid         map[int]map[int]bool
	Instructions []string
	X, Y, Dir    int
}

func Run(file string) [2]interface{} {
	data := utils.ReadFileRaw(file)
	fmt.Println(len(data))

	solver := parse(data)

	return [2]interface{}{
		solver.PartOne(),
		0,
	}
}

func parse(data []string) *Solver {
	grid := make(map[int]map[int]bool)

	for y, row := range data {
		grid[y] = make(map[int]bool)
		if row == "" {
			break
		}
		for x, col := range row {
			if col == '.' {
				grid[y][x] = false
			} else if col == '#' {
				grid[y][x] = true
			}
		}
	}

	instructions := make([]string, 0)
	var inst string
	for _, c := range data[len(data)-2] {
		if c == 'R' || c == 'L' {
			instructions = append(instructions, inst)
			instructions = append(instructions, string(c))
			inst = ""
			continue
		}

		inst += string(c)
	}

	instructions = append(instructions, inst)
	return &Solver{
		Grid:         grid,
		Instructions: instructions,
	}
}

func (s *Solver) PartOne() int {
	for i := 0; true; i++ {
		if _, ok := s.Grid[0][i]; ok {
			s.X = i
			break
		}
	}

	for _, instr := range s.Instructions {
		if instr == "L" {
			s.Dir += 3
			s.Dir %= 4
			continue
		}
		if instr == "R" {
			s.Dir += 1
			s.Dir %= 4
			continue
		}

		num, _ := strconv.Atoi(instr)
		for i := 0; i < num; i++ {
			if !s.Move() {
				break
			}
		}
	}

	return 1000*(s.Y+1) + 4*(s.X+1) + s.Dir
}

// Move moves x and y one space towards direction
// Returns false if it cant continue
func (s *Solver) Move() bool {
	x, y := s.X, s.Y

	// move temp
	switch s.Dir {
	case 0:
		x += 1

		if _, ok := s.Grid[y][x]; ok {
			break
		}
		for i := x - 1; true; i-- {
			if _, ok := s.Grid[y][i]; ok {
				continue
			}

			x = i + 1
			break
		}

	case 1:
		y += 1

		if _, ok := s.Grid[y][x]; ok {
			break
		}
		for i := y - 1; true; i-- {
			if _, ok := s.Grid[i][x]; ok {
				continue
			}

			y = i + 1
			break
		}

	case 2:
		x -= 1

		if _, ok := s.Grid[y][x]; ok {
			break
		}
		for i := x + 1; true; i++ {
			if _, ok := s.Grid[y][i]; ok {
				continue
			}

			x = i - 1
			break
		}

	case 3:
		y -= 1

		if _, ok := s.Grid[y][x]; ok {
			break
		}
		for i := y + 1; true; i++ {
			if _, ok := s.Grid[i][x]; ok {
				continue
			}

			y = i - 1
			break
		}
	}

	if wall := s.Grid[y][x]; wall {
		return false
	}

	s.X = x
	s.Y = y
	return true
}
