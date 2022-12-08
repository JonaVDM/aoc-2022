package day08

import (
	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	grid := utils.ConverToGridInts(data)

	count := 0
	scores := make([]int, 0)
	for i, row := range grid {
		for j := range row {
			if validTree(grid, i, j) {
				count++
			}

			scores = append(scores, countTrees(grid, i, j))
		}
	}

	return [2]interface{}{
		count,
		utils.Max(scores),
	}
}

func countTrees(data [][]int, i, j int) int {
	up := 0
	prev := data[i][j]
	for u := i - 1; u >= 0; u-- {
		up++
		if data[u][j] >= prev {
			break
		}
	}

	down := 0
	prev = data[i][j]
	for u := i + 1; u < len(data); u++ {
		down++
		if data[u][j] >= prev {
			break
		}
	}

	left := 0
	prev = data[i][j]
	for u := j - 1; u >= 0; u-- {
		left++
		if data[i][u] >= prev {
			break
		}
	}

	right := 0
	prev = data[i][j]
	for u := j + 1; u < len(data); u++ {
		right++
		if data[i][u] >= prev {
			break
		}
	}

	return left * right * up * down
}

func validTree(data [][]int, i, j int) bool {
	if i == 0 || i == len(data[0])-1 || j == 0 || j == len(data)-1 {
		return true
	}

	found := false

	// up
	prev := data[i][j]
	for u := i - 1; u >= 0; u-- {
		if data[u][j] >= prev {
			found = false
			break
		}
		found = true
	}
	if found {
		return true
	}

	// down
	prev = data[i][j]
	for u := i + 1; u < len(data); u++ {
		if data[u][j] >= prev {
			found = false
			break
		}
		found = true
	}
	if found {
		return true
	}

	// left
	prev = data[i][j]
	for u := j - 1; u >= 0; u-- {
		if data[i][u] >= prev {
			found = false
			break
		}
		found = true
	}
	if found {
		return true
	}

	// right
	prev = data[i][j]
	for u := j + 1; u < len(data); u++ {
		if data[i][u] >= prev {
			found = false
			break
		}
		found = true
	}

	return found
}
