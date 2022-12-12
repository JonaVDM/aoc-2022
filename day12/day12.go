package day12

import (
	"fmt"

	"github.com/jonavdm/aoc-2022/utils"
)

// Using https://en.wikipedia.org/wiki/Breadth-first_search

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	grid, start, end := parse(data)

	path := bfs(grid, start, end)
	a := 0
	if path == nil {
		panic("No path was found")
	}
	for path.Parent != nil {
		a++
		path = path.Parent
	}

	path = parTwo(grid, end)
	b := 0
	if path == nil {
		panic("No path was found")
	}
	for path.Parent != nil {
		b++
		path = path.Parent
	}

	return [2]interface{}{
		a,
		b,
	}
}

type Coord struct {
	X      int
	Y      int
	Parent *Coord
}

// parse returns the grid, starting position and ending position
func parse(data []string) ([][]int, Coord, Coord) {
	rows := len(data)
	cols := len(data[0])

	start := Coord{}
	end := Coord{}

	grid := make([][]int, rows)
	for y := 0; y < rows; y++ {
		grid[y] = make([]int, cols)
		for x := 0; x < cols; x++ {
			if data[y][x] == 'S' {
				start = Coord{x, y, nil}
				grid[y][x] = 0
				continue
			}

			if data[y][x] == 'E' {
				end = Coord{x, y, nil}
				grid[y][x] = 25
				continue
			}

			grid[y][x] = int(data[y][x] - 'a')
		}
	}

	return grid, start, end
}

func bfs(grid [][]int, start, end Coord) *Coord {
	visited := make(map[string]bool)
	queue := make([]*Coord, 0)

	queue = append(queue, &start)
	visited[fmt.Sprint(start.X, start.Y)] = true

	counter := 0

	for len(queue) > 0 {
		counter++
		// pop from queue
		item := queue[0]
		queue = queue[1:]

		// Is end
		if item.X == end.X && item.Y == end.Y {
			return item
		}

		children := [][]int{
			{item.X + 1, item.Y},
			{item.X - 1, item.Y},
			{item.X, item.Y + 1},
			{item.X, item.Y - 1},
		}
		for _, child := range children {
			// check if out of bounce
			if child[0] < 0 || child[0] >= len(grid[0]) || child[1] < 0 || child[1] >= len(grid) {
				continue
			}

			strcoords := fmt.Sprint(child[0], child[1])

			// check if not already visited
			if _, ok := visited[strcoords]; ok {
				continue
			}

			// Custom check
			if grid[child[1]][child[0]]-grid[item.Y][item.X] > 1 {
				continue
			}

			visited[strcoords] = true
			queue = append(queue, &Coord{
				X:      child[0],
				Y:      child[1],
				Parent: item,
			})
		}
	}

	return nil
}

func parTwo(grid [][]int, start Coord) *Coord {
	visited := make(map[string]bool)
	queue := make([]*Coord, 0)

	queue = append(queue, &start)
	visited[fmt.Sprint(start.X, start.Y)] = true

	counter := 0

	for len(queue) > 0 {
		counter++
		// pop from queue
		item := queue[0]
		queue = queue[1:]

		// Is end
		if grid[item.Y][item.X] == 0 {
			return item
		}

		children := [][]int{
			{item.X + 1, item.Y},
			{item.X - 1, item.Y},
			{item.X, item.Y + 1},
			{item.X, item.Y - 1},
		}
		for _, child := range children {
			// check if out of bounce
			if child[0] < 0 || child[0] >= len(grid[0]) || child[1] < 0 || child[1] >= len(grid) {
				continue
			}

			strcoords := fmt.Sprint(child[0], child[1])

			// check if not already visited
			if _, ok := visited[strcoords]; ok {
				continue
			}

			// Custom check
			if grid[item.Y][item.X]-grid[child[1]][child[0]] > 1 {
				continue
			}

			visited[strcoords] = true
			queue = append(queue, &Coord{
				X:      child[0],
				Y:      child[1],
				Parent: item,
			})
		}
	}

	return nil
}
