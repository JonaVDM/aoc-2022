package day09

import (
	"fmt"
	"strconv"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	headX, headY := 0, 0
	prevHeadX, prevHeadY := 0, 0
	tailX, tailY := 0, 0
	been := make(map[string]bool)
	been["0,0"] = true

	for _, row := range data {
		instruction := row[0]
		amount, err := strconv.Atoi(row[2:])
		if err != nil {
			panic(err)
		}

		for i := 0; i < amount; i++ {
			prevHeadX = headX
			prevHeadY = headY

			if instruction == 'U' {
				headY += 1
			} else if instruction == 'D' {
				headY -= 1
			} else if instruction == 'L' {
				headX -= 1
			} else if instruction == 'R' {
				headX += 1
			} else {
				panic("Invalid instruction")
			}

			// part one
			if utils.AbsInt(tailX-headX) == 2 || utils.AbsInt(tailY-headY) == 2 {
				tailX = prevHeadX
				tailY = prevHeadY

				been[fmt.Sprintf("%d,%d", tailX, tailY)] = true
			}
		}
	}

	return [2]interface{}{
		len(been),
		0,
	}
}
