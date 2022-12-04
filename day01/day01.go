package day01

import (
	"sort"
	"strconv"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run() [2]interface{} {
	data := utils.ReadFile("01")
	cals := findCal(data)

	sort.Ints(cals)
	l := len(cals)

	return [2]interface{}{
		utils.Max(cals),
		cals[l-1] + cals[l-2] + cals[l-3],
	}
}

func findCal(data []string) []int {
	elfs := make([]int, 0)
	elf := 0

	for _, inp := range data {
		if inp == "" {
			elfs = append(elfs, elf)
			elf = 0
			continue
		}

		i, err := strconv.Atoi(inp)
		if err != nil {
			panic("error parsing string")
		}
		elf += i
	}

	return elfs
}
