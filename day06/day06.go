package day06

import (
	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadSingleLineFile(file)

	partA := -1
	for i := 0; i < len(data)-4; i++ {
		if !utils.UniqueString(data[i : i+4]) {
			continue
		}

		partA = i + 4
		break
	}

	partB := -1
	for i := 0; i < len(data)-14; i++ {
		if !utils.UniqueString(data[i : i+14]) {
			continue
		}

		partB = i + 14
		break
	}

	return [2]interface{}{
		partA,
		partB,
	}
}
