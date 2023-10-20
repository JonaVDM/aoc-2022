package day14

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	caveA := parse(data)
	caveB := parse(data)

	bottom := 0
	for k := range caveA {
		if k > bottom {
			bottom = k
		}
	}

	return [2]interface{}{
		partOne(caveA, bottom),
		partTwo(caveB, bottom+2),
	}
}

func parse(file []string) map[int]map[int]bool {
	// data[y][x] = ...
	// where 0 = air and 1 = something else
	data := make(map[int]map[int]bool)

	for _, v := range file {
		points := strings.Split(v, " -> ")

		node := strings.Split(points[0], ",")

		prevX, _ := strconv.Atoi(node[0])
		prevY, _ := strconv.Atoi(node[1])

		for i := 1; i < len(points); i++ {
			node := strings.Split(points[i], ",")

			currX, _ := strconv.Atoi(node[0])
			currY, _ := strconv.Atoi(node[1])

			// if x changed
			if prevX-currX != 0 {
				large := max(currX, prevX)
				small := min(currX, prevX)

				for x := small; x <= large; x++ {
					if _, ok := data[currY]; !ok {
						data[currY] = make(map[int]bool)
					}

					data[currY][x] = true
				}
			}

			// if y changed
			if prevY-currY != 0 {
				large := max(currY, prevY)
				small := min(currY, prevY)

				for y := small; y <= large; y++ {
					if _, ok := data[y]; !ok {
						data[y] = make(map[int]bool)
					}

					data[y][currX] = true
				}
			}

			prevX, prevY = currX, currY
		}
	}

	return data
}

func partOne(cave map[int]map[int]bool, bottom int) int {
	counter := 0

	for {
		x, y := 500, 0

		for {
			if _, ok := cave[y+1]; !ok {
				cave[y+1] = make(map[int]bool)
			}

			// managed to hit the bottom
			if y >= bottom {
				return counter
			}

			// below is air, don't need to check
			if !cave[y+1][x] {
				y += 1
				continue
			}

			// below is not air, left is air
			if !cave[y+1][x-1] {
				y += 1
				x -= 1
				continue
			}

			// below is not air, right is air
			if !cave[y+1][x+1] {
				y += 1
				x += 1
				continue
			}

			// below is not air, neither is left or right
			cave[y][x] = true
			break
		}
		counter++
	}
}

func partTwo(cave map[int]map[int]bool, bottom int) int {
	counter := 0

	for {
		x, y := 500, 0

		counter++
		for {
			if _, ok := cave[y+1]; !ok {
				cave[y+1] = make(map[int]bool)
			}

			// below is the ground
			if y+1 == bottom {
				cave[y][x] = true
				break
			}

			// below is air
			if !cave[y+1][x] {
				y += 1
				continue
			}

			// below is not air, left is air
			if !cave[y+1][x-1] {
				y += 1
				x -= 1
				continue
			}

			// below is not air, right is air
			if !cave[y+1][x+1] {
				y += 1
				x += 1
				continue
			}

			// return when the current is also filled
			if y == 0 {
				return counter
			}

			// below is not air, neither is left or right
			cave[y][x] = true
			break
		}
	}
}
