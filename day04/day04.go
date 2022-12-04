package day04

import (
	"fmt"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run() [2]interface{} {
	data := utils.ReadFile("day04")
	fmt.Println(len(data))

	return [2]interface{}{
		0,
		0,
	}
}
