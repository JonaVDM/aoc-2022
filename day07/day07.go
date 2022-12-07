package day07

import (
	"math"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	root := Node{
		Name:     "/",
		Children: make([]*Node, 0),
		Parent:   nil,
		Type:     "dir",
		Size:     0,
	}

	cur := &root

	for _, v := range data {
		std := strings.Split(v, " ")
		// fmt.Println(std)

		// Is command
		if std[0] == "$" && std[1] == "ls" {
			// 	fmt.Println("List command")
			continue
		}

		if std[0] == "$" && std[1] == "cd" {
			if std[2] == ".." {
				cur = cur.Parent
			}

			for _, c := range cur.Children {
				if c.Name == std[2] {
					cur = c
					break
				}
			}

			continue
		}

		if std[0] == "$" {
			panic("Unkown command")
		}

		// Is directory
		if std[0] == "dir" {
			cur.Children = append(cur.Children, &Node{
				Type:     "dir",
				Name:     std[1],
				Children: make([]*Node, 0),
				Parent:   cur,
				Size:     0,
			})

			continue
		}

		// Is file
		size, err := strconv.Atoi(std[0])
		if err != nil {
			panic(err)
		}

		cur.Children = append(cur.Children, &Node{
			Type:     "file",
			Name:     std[0],
			Children: nil,
			Parent:   cur,
			Size:     size,
		})

		parent := cur
		for {
			parent.Size += size
			if parent.Name == "/" {
				break
			}
			parent = parent.Parent
		}
	}

	required := 30000000 - (70000000 - root.Size)

	return [2]interface{}{
		FindDirs(&root),
		FindSmallest(&root, required),
	}
}

func FindDirs(node *Node) int {
	if node.Type == "file" {
		return 0
	}

	s := 0
	if node.Size <= 100000 {
		s += node.Size
	}

	for _, v := range node.Children {
		s += FindDirs(v)
	}

	return s
}

func FindSmallest(node *Node, required int) int {
	if node.Size < required {
		return math.MaxInt
	}

	sizes := make([]int, 0)
	sizes = append(sizes, node.Size)

	for _, child := range node.Children {
		if child.Type != "dir" {
			continue
		}

		sizes = append(sizes, FindSmallest(child, required))
	}

	return utils.Min(sizes)
}

type Node struct {
	Type     string
	Name     string
	Children []*Node
	Parent   *Node
	Size     int
}
