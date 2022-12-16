package day16

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)
	root, worth, distances := parse(data)

	paths := getPaths(root, worth[1:], distances, 0)

	fmt.Println(len(paths))

	max := 0

	for _, path := range paths {

		score := calculateScore(path, distances)

		if max < score {
			max = score
		}
	}

	return [2]interface{}{
		max,
		0,
	}
}

func parse(data []string) (*Node, []*Node, map[string]map[string]int) {
	var root *Node
	valves := make(map[string]*Node)
	worthIt := make([]*Node, 1)

	for _, row := range data {
		spl := strings.Split(row, " ")

		rate, _ := strconv.Atoi(spl[4][5 : len(spl[4])-1])
		node := &Node{
			Rate: rate,
			Id:   spl[1],
		}

		if rate > 0 {
			worthIt = append(worthIt, node)
		}

		for _, connection := range spl[9:] {
			conn := strings.Trim(connection, ",")
			other, ok := valves[conn]
			if !ok {
				continue
			}
			other.Connected = append(other.Connected, node)
			node.Connected = append(node.Connected, other)
		}

		if spl[1] == "AA" {
			root = node
			worthIt[0] = node
		}

		valves[spl[1]] = node
	}

	distances := make(map[string]map[string]int)
	for i := 0; i < len(worthIt); i++ {
		if _, ok := distances[worthIt[i].Id]; !ok {
			distances[worthIt[i].Id] = make(map[string]int)
		}
		for j := i + 1; j < len(worthIt); j++ {
			if _, ok := distances[worthIt[j].Id]; !ok {
				distances[worthIt[j].Id] = make(map[string]int)
			}

			distance := getDistance(worthIt[i], worthIt[j])

			distances[worthIt[i].Id][worthIt[j].Id] = distance
			distances[worthIt[j].Id][worthIt[i].Id] = distance
		}

	}

	return root, worthIt, distances
}

type Node struct {
	Connected []*Node
	Rate      int
	Id        string

	// For counting the length, just ignore this
	Parent *Node
}

func getDistance(start, end *Node) int {
	queue := make([]*Node, 0)
	visited := make(map[string]bool)

	queue = append(queue, start)
	visited[start.Id] = true
	var item *Node
	start.Parent = nil

	for len(queue) > 0 {
		item = queue[0]
		queue = queue[1:]

		if item == end {
			break
		}

		for _, child := range item.Connected {
			if _, ok := visited[child.Id]; ok {
				// ignore if already visisted
				continue
			}

			visited[child.Id] = true
			child.Parent = item
			queue = append(queue, child)
		}
	}

	counter := 0
	for item.Parent != nil {
		counter++
		item = item.Parent
	}

	return counter
}

func getPaths(head *Node, usable []*Node, distances map[string]map[string]int, count int) [][]*Node {
	nodes := make([][]*Node, 0)

	if len(usable) == 0 {
		return [][]*Node{{head}}
	}

	for k, v := range usable {
		if count+distances[head.Id][v.Id] >= 30 {
			continue
		}

		slice := usable[k+1:]
		slice = append(slice, usable[:k]...)

		rec := getPaths(v, slice, distances, count+distances[head.Id][v.Id])
		for _, child := range rec {
			child = append([]*Node{head}, child...)
			nodes = append(nodes, child)
		}
		if (len(rec)) == 0 {
			nodes = append(nodes, []*Node{head})
		}
	}

	return nodes
}

func calculateScore(path []*Node, distances map[string]map[string]int) int {
	prev := "AA"
	score := 0
	over := 0

	for _, node := range path[1:] {
		over += distances[prev][node.Id]
		over++
		score += node.Rate * (30 - over)

		prev = node.Id
	}

	return score
}
