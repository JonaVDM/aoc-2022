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

func main() {
	printOutput("1", day01.Run())
	printOutput("2", day02.Run2())
	printOutput("3", day03.Run())
	printOutput("4", day04.Run())
	printOutput("5", day05.Run())
	printOutput("6", day06.Run())
}

func printOutput(day string, out [2]interface{}) {
	fmt.Printf("\n--- Day %s ---\nPart One: %v\nPart Two: %v\n", day, out[0], out[1])
}
