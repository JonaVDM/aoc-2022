package day10

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	total := 0
	x := 1
	cycle := 0
	index := 0
	length := len(data)
	next := false
	drawing := "\n"
	position := 0

	for {
		if index >= length {
			break
		}

		// during the cycle
		cycle++
		if (cycle-20)%40 == 0 {
			total += (cycle * x)
		}

		if position == 40 {
			position = 0
			drawing += "\n"
		}

		if utils.AbsInt(position-x) <= 1 {
			drawing += "#"
		} else {
			drawing += "."
		}

		position++

		// after the cycle
		if next {
			next = false
			spl := strings.Split(data[index], " ")
			num, err := strconv.Atoi(spl[1])
			if err != nil {
				panic("Could convert number: " + spl[1])
			}

			x += num
			index++
			continue
		}

		if data[index] == "noop" {
			index++
			continue
		}

		next = true
	}

	return [2]interface{}{
		total,
		drawing,
	}
}
