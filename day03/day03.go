package day03

import (
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run() [2]int {
	data := utils.ReadFile("03")

	// part a
	sumA := 0
	for _, v := range data {
		spl := len(v) / 2
		oth := v[spl:]
		for _, r := range v[:spl] {
			if strings.Contains(oth, string(r)) {
				sumA += getPriority(r)
				break
			}
		}
	}

	// part b
	sumB := 0
	for i := 0; i < len(data); i += 3 {
		for _, v := range data[i] {
			if strings.Contains(data[i+1], string(v)) && strings.Contains(data[i+2], string(v)) {
				sumB += getPriority(v)
				break
			}
		}
	}

	return [2]int{
		sumA, sumB,
	}
}

func getPriority(letter rune) int {
	if letter > 96 {
		return int(letter) - 96
	}
	return int(letter) - 38
}
