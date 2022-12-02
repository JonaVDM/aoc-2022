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
