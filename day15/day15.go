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

	for _, row := range data {
		spl := strings.Split(row, " ")

		sensorX, _ := strconv.Atoi(spl[2][2 : len(spl[2])-1])
		sensorY, _ := strconv.Atoi(spl[3][2 : len(spl[3])-1])

		beaconX, _ := strconv.Atoi(spl[8][2 : len(spl[8])-1])
		beaconY, _ := strconv.Atoi(spl[9][2:])

		distance := utils.AbsInt(sensorX-beaconX) + utils.AbsInt(sensorY-beaconY)
		difference := utils.AbsInt(targetY - sensorY)
		if difference > distance {
			continue
		}

		sensorTarget[sensorX] = true

		for i := 1; i < distance-difference+1; i++ {
			sensorTarget[sensorX+i] = true
			sensorTarget[sensorX-i] = true
		}

		if beaconY == targetY {
			beaconList[beaconX] = true
		}
	}

	return [2]interface{}{
		len(sensorTarget) - len(beaconList),
		0,
	}
}
