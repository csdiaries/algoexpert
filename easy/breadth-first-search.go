package main

// Do not edit the class below except
// for the breadthFirstSearch method.
// Feel free to add new properties
// and methods to the class.
type Node struct {
	Name     string
	Children []*Node
}

func (n *Node) BreadthFirstSearch(array []string) []string {
    queue := []*Node{n} 
    for len(queue) > 0 {
        currentNode := queue[0]
        queue = queue[1:]
        array = append(array, currentNode.Name) 
        for _, child := range currentNode.Children {
            queue = append(queue, child)
        }
    }

    return array
}
