package day09

import (
	"fmt"
	"strconv"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	ropeOne := [2][2]int{}
	ropeTwo := [10][2]int{}

	beenOne := make(map[string]bool)
	beenOne["0,0"] = true

	beenTwo := make(map[string]bool)
	beenTwo["0,0"] = true

	for _, row := range data {
		instruction := row[0]
		amount, err := strconv.Atoi(row[2:])
		if err != nil {
			panic(err)
		}

		for i := 0; i < amount; i++ {
			if instruction == 'U' {
				ropeOne[0][0] -= 1
				ropeTwo[0][0] -= 1
			} else if instruction == 'D' {
				ropeOne[0][0] += 1
				ropeTwo[0][0] += 1
			} else if instruction == 'L' {
				ropeOne[0][1] -= 1
				ropeTwo[0][1] -= 1
			} else if instruction == 'R' {
				ropeOne[0][1] += 1
				ropeTwo[0][1] += 1
			} else {
				panic("Invalid instruction")
			}

			// part one
			if calculateDistance(ropeOne[0], ropeOne[1]) == 2 {
				ropeOne[1] = moveTowards(ropeOne[0], ropeOne[1])
				beenOne[fmt.Sprintf("%d,%d", ropeOne[1][0], ropeOne[1][1])] = true
			}

			// part two
			for i := 1; i <= 9; i++ {
				if calculateDistance(ropeTwo[i-1], ropeTwo[i]) == 2 {
					ropeTwo[i] = moveTowards(ropeTwo[i-1], ropeTwo[i])
				}
			}
			beenTwo[fmt.Sprintf("%d,%d", ropeTwo[9][0], ropeTwo[9][1])] = true
		}
	}

	return [2]interface{}{
		len(beenOne),
		len(beenTwo),
	}
}

func calculateDistance(head, tail [2]int) int {
	// https://en.wikipedia.org/wiki/Chebyshev_distance
	return utils.Max([]int{utils.AbsInt(head[0] - tail[0]), utils.AbsInt(head[1] - tail[1])})
}

func moveTowards(head, tail [2]int) [2]int {
	// up
	if head[1] > tail[1] {
		tail[1] += 1
	}

	// down
	if head[1] < tail[1] {
		tail[1] -= 1
	}

	// left
	if head[0] < tail[0] {
		tail[0] -= 1
	}

	// right
	if head[0] > tail[0] {
		tail[0] += 1
	}

	return tail
}
