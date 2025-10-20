package main

type EduBinaryTree struct {
	root *EduTreeNode
}

func createBinaryTree(nodes []int) *EduTreeNode {
	if len(nodes) == 0 || nodes[0] == 0 {
		return nil
	}
	root := &EduTreeNode{data: nodes[0]}
	queue := []*EduTreeNode{root}
	i := 1
	for i < len(nodes) {
		curr := queue[0]
		queue = queue[1:]
		if i < len(nodes) && nodes[i] != 0 {
			curr.left = &EduTreeNode{data: nodes[i], parent: curr}
			queue = append(queue, curr.left)
		}
		i++
		if i < len(nodes) && nodes[i] != 0 {
			curr.right = &EduTreeNode{data: nodes[i], parent: curr}
			queue = append(queue, curr.right)
		}
		i++
	}
	return root
}

func (tree *EduBinaryTree) find(root *EduTreeNode, value int) *EduTreeNode {
	if root == nil {
		return nil
	}
	queue := []*EduTreeNode{root}
	for len(queue) > 0 {
		currentNode := queue[0]
		queue = queue[1:]
		if currentNode.data == value {
			return currentNode
		}
		if currentNode.left != nil {
			queue = append(queue, currentNode.left)
		}
		if currentNode.right != nil {
			queue = append(queue, currentNode.right)
		}
	}
	return nil
}
