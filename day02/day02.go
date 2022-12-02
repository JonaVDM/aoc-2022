package day02

import (
	"github.com/jonavdm/aoc-2022/utils"
)

func Run() [2]int {
	data := utils.ReadFile("02")

	// Part one
	scoresA := make(map[string]int)
	scoresA["A Z"] = 3
	scoresA["B X"] = 1
	scoresA["C Y"] = 2

	scoresA["C X"] = 7
	scoresA["A Y"] = 8
	scoresA["B Z"] = 9

	scoresA["A X"] = 4
	scoresA["B Y"] = 5
	scoresA["C Z"] = 6

	scoreA := 0
	for _, row := range data {
		scoreA += scoresA[row]
	}

	// Part two
	scoresB := make(map[string]int)
	scoresB["A X"] = 3
	scoresB["B X"] = 1
	scoresB["C X"] = 2

	scoresB["A Z"] = 8
	scoresB["B Z"] = 9
	scoresB["C Z"] = 7

	scoresB["A Y"] = 4
	scoresB["B Y"] = 5
	scoresB["C Y"] = 6

	scoreB := 0
	for _, row := range data {
		scoreB += scoresB[row]
	}

	return [2]int{
		scoreA,
		scoreB,
	}
}

func Run2() [2]int {
	// For part one the formula used was:
	// (a - b + 3) mod 3 = c
	// Where 0 means drawing, 1 losing and 2 winning.

	// For part two the formula used was:
	// (a + c + 1) mod 3 = b

	data := utils.ReadFile("02")
	scoreA := 0
	scoreB := 0

	// some magic numbers
	playerOne := 64
	playerTwo := 87

	for _, v := range data {
		p1 := int(v[0]) - playerOne
		p2 := int(v[2]) - playerTwo
		scoreA += p2

		out := (p1 - p2 + 3) % 3
		if out == 2 {
			scoreA += 6
		} else if out == 0 {
			scoreA += 3
		}

		if p2 == 2 {
			scoreB += 3
		} else if p2 == 3 {
			scoreB += 6
		}

		out = (p1 + p2 + 1) % 3
		if out == 0 {
			out = 3
		}

		scoreB += out
	}

	return [2]int{
		scoreA,
		scoreB,
	}
}
