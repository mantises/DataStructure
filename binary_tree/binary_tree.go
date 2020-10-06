package binary_tree

import (
	"fmt"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func visit(node *TreeNode) {
	if node != nil {
		fmt.Println(node.Val)
	}
}

func PreOrderTraversalByRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	visit(root)
	PreOrderTraversalByRecursive(root.Left)
	PreOrderTraversalByRecursive(root.Right)
}

func InOrderTraversalByRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	InOrderTraversalByRecursive(root.Left)
	visit(root)
	InOrderTraversalByRecursive(root.Right)
}

func PostOrderTraversalByRecursive(root *TreeNode) {
	if root == nil {
		return
	}
	PostOrderTraversalByRecursive(root.Left)
	PostOrderTraversalByRecursive(root.Right)
	visit(root)
}

func PreOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			// pre order visit
			visit(p)
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			p = p.Right
		}
	}
}

func InOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0, 0)
	p := root
	for p != nil || len(stack) != 0 {
		for p != nil {
			stack = append(stack, p)
			p = p.Left
		}
		if len(stack) != 0 {
			p = stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// in order visit
			visit(p)
			p = p.Right
		}
	}
}

func PostOrderTraversal(root *TreeNode) {
	if root == nil {
		return
	}
	stack := make([]*TreeNode, 0, 0)
	stack = append(stack, root)
	var pre *TreeNode
	for len(stack) != 0 {
		p := stack[len(stack)-1]
		if (p.Left == nil && p.Right == nil) || (pre != nil && (pre == p.Left || pre == p.Right)) {
			visit(p)
			stack = stack[:len(stack)-1]
			pre = p
		} else {
			if p.Right != nil {
				stack = append(stack, p.Right)
			}
			if p.Left != nil {
				stack = append(stack, p.Left)
			}
		}
	}
}

func LevelOrderTraversal(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	var res [][]int
	var d1 []*TreeNode
	d1 = append(d1, root)
	for len(d1) != 0 {
		var t []int
		var d2 []*TreeNode
		for i, v := range d1 {
			t = append(t, v.Val)
			if v.Left != nil {
				d2 = append(d2, v.Left)
			}
			if v.Right != nil {
				d2 = append(d2, v.Right)
			}
			if i == len(d1) - 1 {
				res = append(res, t)
				d1, d2 = d2, d1
			}
		}
	}
	return res
}
