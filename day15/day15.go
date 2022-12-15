package day15

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	targetY := 2000000

	sensorTarget := make(map[int]bool)
	beaconList := make(map[int]bool)

	rows := parse(data)

	// part one
	for _, row := range rows {
		distance := utils.AbsInt(row.SensorX-row.BeaconX) + utils.AbsInt(row.SensorY-row.BeaconY)
		difference := utils.AbsInt(targetY - row.SensorY)
		if difference > distance {
			continue
		}

		sensorTarget[row.SensorX] = true

		for i := 1; i < distance-difference+1; i++ {
			sensorTarget[row.SensorX+i] = true
			sensorTarget[row.SensorX-i] = true
		}

		if row.BeaconY == targetY {
			beaconList[row.BeaconX] = true
		}
	}

	ret := make(chan int)
	for i := 0; i < 4000000; i += 100000 {
		go findPartTwo(i, i+100000-1, rows, ret)
	}

	b := 0
	for {
		res := <-ret
		if res == -1 {
			continue
		}
		b = res
		break
	}

	return [2]interface{}{
		len(sensorTarget) - len(beaconList),
		b,
	}
}

type Row struct {
	SensorX int
	SensorY int

	BeaconX int
	BeaconY int
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
		})
	}

	return rows
}

func findPartTwo(start, end int, rows []Row, ret chan int) {
	for i := start; i < end; i++ {
		line := make([]bool, 4000000)

		for _, row := range rows {
			distance := utils.AbsInt(row.SensorX-row.BeaconX) + utils.AbsInt(row.SensorY-row.BeaconY)
			difference := utils.AbsInt(i - row.SensorY)
			if difference > distance {
				continue
			}

			line[row.SensorX] = true

			for i := 1; i < distance-difference+1; i++ {
				if row.SensorX-i >= 0 {
					line[row.SensorX-i] = true
				}

				if row.SensorX+i < 4000000 {
					line[row.SensorX+i] = true
				}
			}
		}

		for x, l := range line {
			if !l {
				ret <- 4000000*x + i
			}
		}
	}

	ret <- -1
}
