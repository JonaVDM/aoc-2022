package day03

import (
	"github.com/jonavdm/aoc-2022/utils"
)

func Run() [2]int {
	data := utils.ReadFile("03")

	// part a
	sumA := 0
	for _, v := range data {
		spl := len(v) / 2

		var letter byte
		for i := 0; i < spl; i++ {
			for j := spl; j < len(v); j++ {
				if v[i] == v[j] {
					letter = v[i]
					break
				}
				if letter != 0 {
					break
				}
			}
		}

		sumA += getPriority(letter)
	}

	// part b
	sumB := 0
	for i := 0; i < len(data); i += 3 {
		var letter byte

	Outer:
		for x := 0; x < len(data[i]); x++ {
			for y := 0; y < len(data[i+1]); y++ {
				for z := 0; z < len(data[i+2]); z++ {
					if data[i][x] == data[i+1][y] && data[i][x] == data[i+2][z] {
						letter = data[i][x]
						break Outer
					}
				}
			}
		}

		sumB += getPriority(letter)
	}

	return [2]int{
		sumA, sumB,
	}
}

func getPriority(letter byte) int {
	if letter > 96 {
		return int(letter) - 96
	}
	return int(letter) - 38
}
