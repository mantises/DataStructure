package binary_tree

import (
	"fmt"
	"testing"
)

func TestPreOrderTraversal(t *testing.T) {
	n1 := &TreeNode{Val: 10}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 55}
	n5 := &TreeNode{Val: -4}
	n6 := &TreeNode{Val: -2}
	n7 := &TreeNode{Val: 190}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n5
	n3.Right = n4
	n5.Right = n6
	n4.Left = n7
	PreOrderTraversalByRecursive(n1)
	fmt.Println("--------------------")
	PreOrderTraversal(n1)
}

func TestInOrderTraversal(t *testing.T) {
	n1 := &TreeNode{Val: 10}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 55}
	n5 := &TreeNode{Val: -4}
	n6 := &TreeNode{Val: -2}
	n7 := &TreeNode{Val: 190}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n5
	n3.Right = n4
	n5.Right = n6
	n4.Left = n7
	InOrderTraversalByRecursive(n1)
	fmt.Println("--------------------")
	InOrderTraversal(n1)
}

func TestPostOrderTraversal(t *testing.T) {
	n1 := &TreeNode{Val: 10}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 55}
	n5 := &TreeNode{Val: -4}
	n6 := &TreeNode{Val: -2}
	n7 := &TreeNode{Val: 190}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n5
	n3.Right = n4
	n5.Right = n6
	n4.Left = n7
	PostOrderTraversalByRecursive(n1)
	fmt.Println("--------------------")
	PostOrderTraversal(n1)
}

func TestLevelOrderTraversal(t *testing.T) {
	n1 := &TreeNode{Val: 10}
	n2 := &TreeNode{Val: 3}
	n3 := &TreeNode{Val: 5}
	n4 := &TreeNode{Val: 55}
	n5 := &TreeNode{Val: -4}
	n6 := &TreeNode{Val: -2}
	n7 := &TreeNode{Val: 190}
	n1.Left = n2
	n1.Right = n3
	n2.Left = n5
	n3.Right = n4
	n5.Right = n6
	n4.Left = n7
	fmt.Println(LevelOrderTraversal(n1))
}
