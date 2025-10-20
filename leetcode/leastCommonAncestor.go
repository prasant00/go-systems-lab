package main

import (
	"fmt"
	"strings"
)

func lowestCommonAncestor(p, q *EduTreeNode) *EduTreeNode {
	ptr1, ptr2 := p, q

	for ptr1 != ptr2 {
		if ptr1.parent != nil {
			ptr1 = ptr1.parent
		} else {
			ptr1 = q
		}

		if ptr2.parent != nil {
			ptr2 = ptr2.parent
		} else {
			ptr2 = p
		}
	}

	return ptr1
}

// Driver code
func main() {
	inputTrees := []int{100, 50, 200, 25, 75, 350}

	inputNodes := []int{25, 75}
	p := inputNodes[0]
	q := inputNodes[1]

	tree := &EduBinaryTree{root: createBinaryTree(inputTrees)}
	displayTree(tree.root)
	lca := LowestCommonAncestor(p, q)
	fmt.Printf("\n\tLowest common ancestor: %d\n", lca.data)
	fmt.Println(strings.Repeat("-", 100))

}
