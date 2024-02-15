package main

import (
    "math"
)
// This is an input class. Do not edit.
type BinaryTree struct {
	Value int

	Left  *BinaryTree
	Right *BinaryTree
}

func BinaryTreeDiameter(tree *BinaryTree) int {
    var max int
    
    helper(tree, &max)
    return max
}

func helper(node *BinaryTree, maxDiameter *int) int {
	if node == nil {
        return 0
    }

    left := helper(node.Left, maxDiameter)
    right := helper(node.Right, maxDiameter)

    *maxDiameter = int(math.Max(float64(*maxDiameter), float64(left+right)))
    return 1 + int(math.Max(float64(left), float64(right)))
}