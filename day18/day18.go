package day18

import (
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	points := parse(data)

	score := len(points) * 6
	for i := 0; i < len(points); i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i].x == points[j].x && points[i].y == points[j].y && utils.AbsInt(points[i].z-points[j].z) <= 1 {
				score -= 2
				continue
			}
			if points[i].y == points[j].y && points[i].z == points[j].z && utils.AbsInt(points[i].x-points[j].x) <= 1 {
				score -= 2
				continue
			}
			if points[i].z == points[j].z && points[i].x == points[j].x && utils.AbsInt(points[i].y-points[j].y) <= 1 {
				score -= 2
				continue
			}
		}
	}

	return [2]interface{}{
		score,
		0,
	}
}

type Point struct {
	x, y, z int
}

func parse(data []string) []*Point {
	points := make([]*Point, 0)

	for _, line := range data {
		spl := strings.Split(line, ",")

		x, err := strconv.Atoi(spl[0])
		if err != nil {
			panic(err)
		}

		y, err := strconv.Atoi(spl[1])
		if err != nil {
			panic(err)
		}

		z, err := strconv.Atoi(spl[2])
		if err != nil {
			panic(err)
		}

		points = append(points, &Point{
			x: x,
			y: y,
			z: z,
		})
	}

	return points
}
