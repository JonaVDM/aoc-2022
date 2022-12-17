package day17

import (
	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadSingleLineFile(file)
	// data := ">>><<><>><<<>><>>><<<>>><<<><<<>><>><<>>"

	shapes := []Shape{
		// -
		{
			{0, 0},
			{1, 0},
			{2, 0},
			{3, 0},
		},

		// +
		{
			{1, 0},
			{0, 1},
			{1, 1},
			{1, 2},
			{2, 1},
		},

		// mirrored L
		{
			{0, 0},
			{1, 0},
			{2, 0},
			{2, 1},
			{2, 2},
		},

		// |
		{
			{0, 0},
			{0, 1},
			{0, 2},
			{0, 3},
		},

		// Square
		{
			{0, 0},
			{1, 0},
			{0, 1},
			{1, 1},
		},
	}

	shapeCounter := 0
	instructionCounter := 0
	chamber := make([][7]bool, 0)

	for i := 0; i < 2022; i++ {
		bottomLeft := Pos{2, len(chamber) + 3}

		for {
			// Check if we can go sideways
			if canMoveSideways(chamber, shapes[shapeCounter], bottomLeft, data[instructionCounter] == '>') {
				if data[instructionCounter] == '>' {
					bottomLeft.X += 1
				} else {
					bottomLeft.X -= 1
				}
			}

			// Check if we can go down
			moveDownCheck := canMoveDown(chamber, shapes[shapeCounter], bottomLeft)
			if moveDownCheck {
				bottomLeft.Y -= 1
			}

			instructionCounter++
			if instructionCounter >= len(data) {
				instructionCounter = 0
			}

			if !moveDownCheck {
				for _, node := range shapes[shapeCounter] {
					if bottomLeft.Y+node.Y+1 > len(chamber) {
						chamber = append(chamber, [7]bool{})
					}

					chamber[bottomLeft.Y+node.Y][bottomLeft.X+node.X] = true
				}

				shapeCounter++
				if shapeCounter >= len(shapes) {
					shapeCounter = 0
				}
				break
			}
		}
	}

	return [2]interface{}{
		len(chamber),
		0,
	}
}

type Pos struct {
	X int
	Y int
}

type Shape []Pos

func canMoveSideways(chamber [][7]bool, shape Shape, corner Pos, right bool) bool {
	height := len(chamber)

	if !right && corner.X-1 < 0 {
		return false
	}

	for _, pos := range shape {
		if right {
			if corner.X+pos.X+1 >= 7 {
				return false
			}

			if corner.Y+pos.Y < height && chamber[corner.Y+pos.Y][corner.X+pos.X+1] {
				return false
			}
		} else {
			if corner.Y+pos.Y < height && chamber[corner.Y+pos.Y][corner.X+pos.X-1] {
				return false
			}
		}
	}
	return true
}

func canMoveDown(chamber [][7]bool, shape Shape, corner Pos) bool {
	if corner.Y <= 0 {
		return false
	}

	height := len(chamber)

	for _, pos := range shape {
		if corner.Y+pos.Y-1 < height && chamber[corner.Y+pos.Y-1][corner.X+pos.X] {
			return false
		}
	}

	return true
}
