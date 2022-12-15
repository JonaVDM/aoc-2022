package day15

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	rows := parse(data)

	return [2]interface{}{
		partOne(rows),
		partTwo(rows),
	}
}

type Row struct {
	SensorX int
	SensorY int

	BeaconX int
	BeaconY int

	Manhatten int
}

func parse(data []string) []Row {
	rows := make([]Row, 0)
	for _, row := range data {
		spl := strings.Split(row, " ")

		sensorX, _ := strconv.Atoi(spl[2][2 : len(spl[2])-1])
		sensorY, _ := strconv.Atoi(spl[3][2 : len(spl[3])-1])

		beaconX, _ := strconv.Atoi(spl[8][2 : len(spl[8])-1])
		beaconY, _ := strconv.Atoi(spl[9][2:])

		rows = append(rows, Row{
			sensorX,
			sensorY,
			beaconX,
			beaconY,
			utils.AbsInt(sensorX-beaconX) + utils.AbsInt(sensorY-beaconY),
		})
	}

	return rows
}

func partOne(rows []Row) int {
	targetY := 2000000
	sensorTarget := make(map[int]bool)
	beaconList := make(map[int]bool)

	for _, row := range rows {
		difference := utils.AbsInt(targetY - row.SensorY)
		if difference > row.Manhatten {
			continue
		}

		sensorTarget[row.SensorX] = true

		for i := 1; i < row.Manhatten-difference+1; i++ {
			sensorTarget[row.SensorX+i] = true
			sensorTarget[row.SensorX-i] = true
		}

		if row.BeaconY == targetY {
			beaconList[row.BeaconX] = true
		}
	}

	return len(sensorTarget) - len(beaconList)
}

func partTwo(rows []Row) int {
	matching := make(map[string]int)

	for i := 0; i < len(rows); i++ {
		for x := i + 1; x < len(rows); x++ {
			dis := utils.AbsInt(rows[i].SensorX-rows[x].SensorX) + utils.AbsInt(rows[i].SensorY-rows[x].SensorY)
			if dis != rows[i].Manhatten+rows[x].Manhatten+2 {
				continue
			}

			var big Row
			var small Row

			if rows[i].Manhatten < rows[x].Manhatten {
				// i is smaller
				big = rows[x]
				small = rows[i]
			} else {
				// x is smaller
				big = rows[i]
				small = rows[x]
			}

			x := small.Manhatten + 1
			y := 0

			for {
				var key string = "woeps"
				// Top left
				if small.SensorX < big.SensorX && small.SensorY < big.SensorY {
					key = fmt.Sprintf("%d,%d", small.SensorX+x, small.SensorY+y)
				}

				// Top right
				if small.SensorX > big.SensorX && small.SensorY < big.SensorY {
					key = fmt.Sprintf("%d,%d", small.SensorX-x, small.SensorY+y)
				}

				// Bottom left
				if small.SensorX < big.SensorX && small.SensorY > big.SensorY {
					key = fmt.Sprintf("%d,%d", small.SensorX+x, small.SensorY-y)
				}

				// Bottom right
				if small.SensorX > big.SensorX && small.SensorY > big.SensorY {
					key = fmt.Sprintf("%d,%d", small.SensorX-x, small.SensorY-y)
				}

				if key == "woeps" {
					panic("woeps")
				}

				matching[key]++

				if x == 0 {
					break
				}

				y++
				x--
			}
		}
	}

	max := 0
	maxKey := ""
	for k, value := range matching {
		if value > max {
			max = value
			maxKey = k
		}
	}

	key := strings.Split(maxKey, ",")
	x, _ := strconv.Atoi(key[0])
	y, _ := strconv.Atoi(key[1])

	return 4000000*x + y
}
