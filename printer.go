package main

import (
	"chicofilho/merkletree/merkleTree"
	"fmt"
)

func PrintMerkleTree(mt merkleTree.MerkleTree) {
	fmt.Println("====== The Tree ======")
	stack := make([]*merkleTree.Node, 0, len(mt.Leafs))
	stack = append(stack, mt.Root)
	nextStep := make([]*merkleTree.Node, 0, len(stack)*2)
	height := 0
	for len(stack) > 0 {
		el := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		fmt.Println(nodeWithSpaces(el, height))

		if !el.IsLeaf() {
			nextStep = append(nextStep, el.Left)
			nextStep = append(nextStep, el.Right)
		}
		if len(stack) == 0 && len(nextStep) > 0 {
			stack = nextStep
			nextStep = make([]*merkleTree.Node, 0, len(stack)*2)
			height++
		}
	}
	fmt.Println("====== Finished The Tree ======")
}

func nodeWithSpaces(node *merkleTree.Node, height int) string {
	return addSpaces(height) + node.GetHash().ToString()
}

func addSpaces(amount int) string {
	spaces := "  "
	for i := 0; i < amount; i++ {
		spaces += spaces
	}
	return spaces
}
