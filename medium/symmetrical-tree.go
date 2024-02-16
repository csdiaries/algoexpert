package main

// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func SymmetricalTree(tree *BinaryTree) bool {
	return helper(tree.Left, tree.Right)
}

func helper(left, right *BinaryTree) bool {
	if left != nil && right != nil && left.Value == right.Value {
        return helper(left.Left, right.Right) && helper(left.Right, right.Left)
    }

	return left == right
}
