package day11

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	monkeys := parse(data)
	countersA := make([]int, len(monkeys))
	countersB := make([]int, len(monkeys))

	// part one
	for i := 0; i < 20; i++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].Items {
				countersA[m] += 1
				n1, err := strconv.Atoi(monkeys[m].Operation[0])
				if err != nil {
					n1 = item
				}

				n2, err := strconv.Atoi(monkeys[m].Operation[2])
				if err != nil {
					n2 = item
				}

				var newValue int
				if monkeys[m].Operation[1] == "*" {
					newValue = n1 * n2
				} else if monkeys[m].Operation[1] == "+" {
					newValue = n1 + n2
				} else {
					panic("Invalid operation")
				}

				newValue = newValue / 3
				if newValue%monkeys[m].Test == 0 {
					monkeys[monkeys[m].IfTrue].Items = append(monkeys[monkeys[m].IfTrue].Items, newValue)
				} else {
					monkeys[monkeys[m].IfFalse].Items = append(monkeys[monkeys[m].IfFalse].Items, newValue)
				}
			}
			monkeys[m].Items = make([]int, 0)
		}
	}

	// part one
	monkeys = parse(data)
	for i := 0; i < 10000; i++ {
		for m := 0; m < len(monkeys); m++ {
			for _, item := range monkeys[m].Items {
				countersB[m] += 1
				n1, err := strconv.Atoi(monkeys[m].Operation[0])
				if err != nil {
					n1 = item
				}

				n2, err := strconv.Atoi(monkeys[m].Operation[2])
				if err != nil {
					n2 = item
				}

				var newValue int
				if monkeys[m].Operation[1] == "*" {
					newValue = n1 * n2
				} else if monkeys[m].Operation[1] == "+" {
					newValue = n1 + n2
				} else {
					panic("Invalid operation")
				}

				if newValue%monkeys[m].Test == 0 {
					monkeys[monkeys[m].IfTrue].Items = append(monkeys[monkeys[m].IfTrue].Items, newValue)
				} else {
					monkeys[monkeys[m].IfFalse].Items = append(monkeys[monkeys[m].IfFalse].Items, newValue)
				}
			}
			monkeys[m].Items = make([]int, 0)
		}
	}
	// 1_327_451_967_748_545_085
	// 8_770_139_538_443_747_337
	// 18_446_744_073_709_551_615 (max u64)

	for i, monk := range monkeys {
		fmt.Println(countersB[i])
		fmt.Println(monk.Items)
	}

	sort.Ints(countersA)
	sort.Ints(countersB)
	length := len(countersA)

	return [2]interface{}{
		countersA[length-1] * countersA[length-2],
		countersB[length-1] * countersB[length-2],
	}
}

type Monkey struct {
	Test      int
	Items     []int
	IfTrue    int
	IfFalse   int
	Operation []string
}

func parse(data []string) []Monkey {
	monkeys := make([]Monkey, 0)

	curr := -1
	for _, monk := range data {
		instruction := strings.Split(strings.TrimSpace(monk), " ")

		if instruction[0] == "Monkey" {
			curr++
			monkeys = append(monkeys, Monkey{
				Items: make([]int, 0),
			})
			continue
		}

		if instruction[0] == "Starting" {
			for _, item := range instruction[2:] {
				num, err := strconv.Atoi(strings.Trim(item, ","))
				if err != nil {
					panic(err)
				}

				monkeys[curr].Items = append(monkeys[curr].Items, num)
			}
			continue
		}

		if instruction[0] == "Operation:" {
			monkeys[curr].Operation = instruction[3:]
			continue
		}

		if instruction[0] == "Test:" {
			num, err := strconv.Atoi(instruction[len(instruction)-1])
			if err != nil {
				panic(err)
			}
			monkeys[curr].Test = num
			continue
		}

		if instruction[0] == "If" && instruction[1] == "true:" {
			num, err := strconv.Atoi(instruction[len(instruction)-1])
			if err != nil {
				panic(err)
			}
			monkeys[curr].IfTrue = num
			continue
		}

		if instruction[0] == "If" && instruction[1] == "false:" {
			num, err := strconv.Atoi(instruction[len(instruction)-1])
			if err != nil {
				panic(err)
			}
			monkeys[curr].IfFalse = num
			continue
		}
	}

	return monkeys
}
