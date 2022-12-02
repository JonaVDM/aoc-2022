package main

import (
	"fmt"

	"github.com/jonavdm/aoc-2022/day01"
)

func main() {
	printOutput("1", day01.Run())
}

func printOutput(day string, out [2]int) {
	fmt.Printf("--- Day %s ---\nPart One: %d\nPart Two: %d\n", day, out[0], out[1])
}
