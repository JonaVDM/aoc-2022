package day20

import (
	"fmt"
	"strconv"

	"github.com/jonavdm/aoc-2022/utils"
)

type Item struct {
	Prev *Item
	Next *Item

	OPrev *Item
	ONext *Item

	Value int
}

func (i *Item) Move(size int) {
	// We can ignore the 0 case
	if i.Value == 0 {
		return
	}

	i.Prev.Next = i.Next
	i.Next.Prev = i.Prev
	curr := i

	// Negative
	if i.Value < 0 {
		// amount := utils.AbsInt(i.Value) % size
		for x := 0; x > i.Value%(size-1); x-- {
			curr = curr.Prev
		}

		i.Prev = curr.Prev
		i.Next = curr
	}

	// Positive
	if i.Value > 0 {
		for x := 0; x < i.Value%(size-1); x++ {
			curr = curr.Next
		}

		i.Next = curr.Next
		i.Prev = curr
	}

	i.Next.Prev = i
	i.Prev.Next = i
}

func (i *Item) Print() {
	values := make([]int, 0)
	values = append(values, i.Value)

	curr := i.Next
	for {
		if i == curr {
			break
		}

		values = append(values, curr.Value)
		curr = curr.Next
	}

	fmt.Println(values)
}

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	return [2]interface{}{
		partOne(data),
		partTwo(data),
	}
}

func partOne(data []string) int {
	root := parse(data)
	curr := root
	var zero *Item

	amount := len(data)
	for i := 0; i < amount; i++ {
		if curr.Value == 0 {
			zero = curr
		}
		curr.Move(amount)
		curr = curr.ONext
	}

	total := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			zero = zero.Next
		}

		total += zero.Value
	}

	return total
}

func partTwo(data []string) int {
	root := parseBig(data)
	curr := root
	var zero *Item

	amount := len(data)
	for i := 0; i < amount*10; i++ {
		if curr.Value == 0 {
			zero = curr
		}
		curr.Move(amount)
		curr = curr.ONext
	}

	total := 0
	for i := 0; i < 3; i++ {
		for j := 0; j < 1000; j++ {
			zero = zero.Next
		}

		total += zero.Value
	}

	return total
}

func parse(data []string) *Item {
	root := &Item{}
	curr := root
	list := make([]int, len(data))

	for i, v := range data {
		num, _ := strconv.Atoi(v)
		list[i] = num
		curr.Next = &Item{
			Prev:  curr,
			OPrev: curr,
			Value: num,
		}

		curr.ONext = curr.Next
		curr = curr.Next
	}

	root = root.Next
	root.Prev = curr
	root.OPrev = curr
	curr.Next = root
	curr.ONext = root
	return root
}

func parseBig(data []string) *Item {
	root := &Item{}
	curr := root
	list := make([]int, len(data))

	for i, v := range data {
		num, _ := strconv.Atoi(v)
		list[i] = num
		curr.Next = &Item{
			Prev:  curr,
			OPrev: curr,
			Value: num * 811589153,
		}

		curr.ONext = curr.Next
		curr = curr.Next
	}

	root = root.Next
	root.Prev = curr
	root.OPrev = curr
	curr.Next = root
	curr.ONext = root
	return root
}
