package bst

import (
	"goleetcode/datastruct/trees"
)

func SearchRec(root *trees.TreeNode, val int) *trees.TreeNode {
	if root.Val == val {
		return root
	}

	if root.Val < val {
		return SearchRec(root.Right, val)
	} else if root.Val > val {
		return SearchRec(root.Left, val)
	}
	return nil
}

// TODO finish
func SearchIterative(root *trees.TreeNode, val int) *trees.TreeNode {
	return nil
}
