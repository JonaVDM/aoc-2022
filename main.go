package main

import (
	"flag"
	"fmt"

	"github.com/jonavdm/aoc-2022/day01"
	"github.com/jonavdm/aoc-2022/day02"
	"github.com/jonavdm/aoc-2022/day03"
	"github.com/jonavdm/aoc-2022/day04"
	"github.com/jonavdm/aoc-2022/day05"
	"github.com/jonavdm/aoc-2022/day06"
	"github.com/jonavdm/aoc-2022/day07"
	"github.com/jonavdm/aoc-2022/day08"
	"github.com/jonavdm/aoc-2022/day09"
	"github.com/jonavdm/aoc-2022/day10"
	"github.com/jonavdm/aoc-2022/day11"
	"github.com/jonavdm/aoc-2022/day12"
	"github.com/jonavdm/aoc-2022/day14"
	"github.com/jonavdm/aoc-2022/day15"
	"github.com/jonavdm/aoc-2022/day16"
	"github.com/jonavdm/aoc-2022/day17"
	"github.com/jonavdm/aoc-2022/day18"
	"github.com/jonavdm/aoc-2022/day20"
	"github.com/jonavdm/aoc-2022/day21"
)

type Runner struct {
	Day      int
	Function func(string) [2]interface{}
	File     string
}

func main() {
	onlyDay := flag.Int("day", -1, "Specify the day")
	replacedInput := flag.String("file", "", "Run with a different input")
	flag.Parse()

	runners := []Runner{
		{1, day01.Run, "day01"},
		{2, day02.Run2, "day02"},
		{3, day03.Run, "day03"},
		{4, day04.Run, "day04"},
		{5, day05.Run, "day05"},
		{6, day06.Run, "day06"},
		{7, day07.Run, "day07"},
		{8, day08.Run, "day08"},
		{9, day09.Run, "day09"},
		{10, day10.Run, "day10"},
		{11, day11.Run, "day11"},
		{12, day12.Run, "day12"},
		{14, day14.Run, "day14"},
		{15, day15.Run, "day15"},
		{16, day16.Run, "day16"},
		{17, day17.Run, "day17"},
		{18, day18.Run, "day18"},
		// {19, day19.Run, "day19"},
		{20, day20.Run, "day20"},
		{21, day21.Run, "day21"},
	}

	for _, runner := range runners {
		if *onlyDay > 0 && runner.Day == *onlyDay {
			file := runner.File
			if *replacedInput != "" {
				file = *replacedInput
			}
			out := runner.Function(file)
			printOutput(runner.Day, out)
		}

		if *onlyDay == -1 {
			out := runner.Function(runner.File)
			printOutput(runner.Day, out)
		}
	}
}

func printOutput(day int, out [2]interface{}) {
	fmt.Printf("\n--- Day %d ---\nPart One: %v\nPart Two: %v\n", day, out[0], out[1])
}
