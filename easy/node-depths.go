package main

type BinaryTree struct {
	Value       int
	Left, Right *BinaryTree
}

func NodeDepths(root *BinaryTree) int {
	return nodeDepthsHelper(root, 0)
}

func nodeDepthsHelper(node *BinaryTree, depth int) int {
    if node == nil {
        return 0
    }
    
    return depth + nodeDepthsHelper(node.Left, depth + 1) + nodeDepthsHelper(node.Right, depth + 1)
}
