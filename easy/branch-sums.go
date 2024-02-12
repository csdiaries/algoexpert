package main

// This is the struct of the input root. Do not edit it.
type BinaryTree struct {
	Value int
	Left  *BinaryTree
	Right *BinaryTree
}

func BranchSums(root *BinaryTree) []int {
    var result []int
    return BranchSumsHelper(root, 0, result)
}

func BranchSumsHelper(node *BinaryTree, currentSum int, result []int) []int {
    if node == nil {
        return result
    }
    
    currentSum += node.Value
    if node.Left == nil && node.Right == nil {
        return append(result, currentSum)
    }
    
    result = BranchSumsHelper(node.Left, currentSum, result)
    result = BranchSumsHelper(node.Right, currentSum, result)
    
    return result
}
    
   