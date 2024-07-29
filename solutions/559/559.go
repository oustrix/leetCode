package _559

import (
	"slices"
)

type Node struct {
	Val      int
	Children []*Node
}

func maxDepth(root *Node) int {
	if root == nil {
		return 0
	}

	return calcMax(*root, 1)
}

func calcMax(node Node, current int) int {
	if len(node.Children) == 0 {
		return current
	}

	vals := make([]int, 0)
	for _, v := range node.Children {
		vals = append(vals, calcMax(*v, current+1))
	}

	return slices.Max(vals)

}
