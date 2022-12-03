package main

import (
	"fmt"

	"github.com/jonavdm/aoc-2022/day01"
	"github.com/jonavdm/aoc-2022/day02"
	"github.com/jonavdm/aoc-2022/day03"
)

func main() {
	printOutput("1", day01.Run())
	printOutput("2", day02.Run2())
	printOutput("3", day03.Run())
}

func printOutput(day string, out [2]int) {
	fmt.Printf("\n--- Day %s ---\nPart One: %d\nPart Two: %d\n", day, out[0], out[1])
}
