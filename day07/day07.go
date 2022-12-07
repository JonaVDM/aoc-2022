package day07

import (
	"math"
	"strconv"
	"strings"

	"github.com/jonavdm/aoc-2022/utils"
)

func Run(file string) [2]interface{} {
	data := utils.ReadFile(file)

	root := ParseFileSystem(data)

	required := 30000000 - (70000000 - root.Size)

	return [2]interface{}{
		FindDirs(&root),
		FindSmallest(&root, required),
	}
}

func FindDirs(node *Node) int {
	if !node.IsDir {
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
		if !child.IsDir {
			continue
		}

		sizes = append(sizes, FindSmallest(child, required))
	}

	return utils.Min(sizes)
}

func ParseFileSystem(data []string) Node {
	root := Node{
		Name:     "/",
		Children: make([]*Node, 0),
		Parent:   nil,
		Size:     0,
		IsDir:    true,
	}

	cur := &root

	for _, v := range data {
		std := strings.Split(v, " ")

		// Is command
		if std[0] == "$" && std[1] == "ls" {
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
				Name:     std[1],
				Children: make([]*Node, 0),
				Parent:   cur,
				Size:     0,
				IsDir:    true,
			})

			continue
		}

		// Is file
		size, err := strconv.Atoi(std[0])
		if err != nil {
			panic(err)
		}

		cur.Children = append(cur.Children, &Node{
			Name:     std[0],
			Children: nil,
			Parent:   cur,
			Size:     size,
			IsDir:    false,
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

	return root
}

type Node struct {
	Name     string
	Children []*Node
	Parent   *Node
	Size     int
	IsDir    bool
}
