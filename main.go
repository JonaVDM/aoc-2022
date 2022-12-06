package main

import (
	"fmt"

	"github.com/jonavdm/aoc-2022/day01"
	"github.com/jonavdm/aoc-2022/day02"
	"github.com/jonavdm/aoc-2022/day03"
	"github.com/jonavdm/aoc-2022/day04"
	"github.com/jonavdm/aoc-2022/day05"
	"github.com/jonavdm/aoc-2022/day06"
)

type Runner struct {
	Day      int
	Function func(string) [2]interface{}
	File     string
}

func main() {
	runners := []Runner{
		{1, day01.Run, "day01"},
		{2, day02.Run2, "day02"},
		{3, day03.Run, "day03"},
		{4, day04.Run, "day04"},
		{5, day05.Run, "day05"},
		{6, day06.Run, "day06"},
	}

	for _, runner := range runners {
		out := runner.Function(runner.File)
		printOutput(runner.Day, out)
	}
}

func printOutput(day int, out [2]interface{}) {
	fmt.Printf("\n--- Day %d ---\nPart One: %v\nPart Two: %v\n", day, out[0], out[1])
}
