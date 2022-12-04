package day04

import (
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run() [2]interface{} {
	data := utils.ReadFile("day04")

	partA := 0
	partB := 0

	for _, pair := range data {
		raw := strings.FieldsFunc(pair, func(r rune) bool {
			return r == '-' || r == ','
		})

		line := utils.ConvertToInts(raw)
		if (line[0] >= line[2] && line[1] <= line[3]) || (line[0] <= line[2] && line[1] >= line[3]) {
			partA += 1
		}

		// This one is faster but it shouldn't solve it but it does.
		// utils.Between(line[1], line[2], line[3]) || utils.Between(line[3], line[0], line[1])
		if utils.Overlap(utils.IntRange(line[0], line[1]+1), utils.IntRange(line[2], line[3]+1)) {
			partB += 1
		}
	}

	return [2]interface{}{
		partA,
		partB,
	}
}
