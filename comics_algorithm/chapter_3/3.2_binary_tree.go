package chapter_3

import (
	"../utils"
)

type BinaryTreeNode struct {
	Data       int
	LeftChild  *BinaryTreeNode
	RightChild *BinaryTreeNode
}

var (
	EmptyNodeData = -1
)

func NewBinaryTree(inputs *[]int) *BinaryTreeNode {
	ipts := *inputs

	if len(ipts) == 0 {
		return nil
	}

	data := ipts[0]
	*inputs = ipts[1:]

	if data == EmptyNodeData {
		return nil
	}

	btn := &BinaryTreeNode{
		Data:       data,
		LeftChild:  NewBinaryTree(inputs),
		RightChild: NewBinaryTree(inputs),
	}

	return btn
}

type BinaryTreeIterator struct {
}

func (bti *BinaryTreeIterator) PreOrder(btn *BinaryTreeNode, result *[]int) {
	if btn == nil {
		return
	}

	*result = append(*result, btn.Data)
	bti.PreOrder(btn.LeftChild, result)
	bti.PreOrder(btn.RightChild, result)
}

func (bti *BinaryTreeIterator) InOrder(btn *BinaryTreeNode, result *[]int) {
	if btn == nil {
		return
	}

	bti.InOrder(btn.LeftChild, result)
	*result = append(*result, btn.Data)
	bti.InOrder(btn.RightChild, result)

}

func (bti *BinaryTreeIterator) PostOrder(btn *BinaryTreeNode, result *[]int) {
	if btn == nil {
		return
	}

	bti.PostOrder(btn.LeftChild, result)
	bti.PostOrder(btn.RightChild, result)
	*result = append(*result, btn.Data)
}

func (bti *BinaryTreeIterator) LevelOrder(btn *BinaryTreeNode, result *[]int) {
	queue := &utils.Queue{}

	queue.In(btn)
	for !queue.IsEmpty() {
		if item, err := queue.Out(); err == nil {
			treeNode := item.(*BinaryTreeNode)
			*result = append(*result, treeNode.Data)
			if treeNode.LeftChild != nil {
				queue.In(treeNode.LeftChild)
			}
			if treeNode.RightChild != nil {
				queue.In(treeNode.RightChild)
			}
		}
	}
}

func (bti *BinaryTreeIterator) PreOrderWithStack(btn *BinaryTreeNode, result *[]int) {
	stack := &utils.Stack{}

	treeNode := btn
	for treeNode != nil || !stack.IsEmpty() {
		// traverse left child and put in stack
		for treeNode != nil {
			*result = append(*result, treeNode.Data)
			stack.Push(treeNode)
			treeNode = treeNode.LeftChild
		}

		// if left child is empty
		if !stack.IsEmpty() {
			if item, err := stack.Pop(); err == nil {
				treeNode = item.(*BinaryTreeNode)
				treeNode = treeNode.RightChild
			}
		}
	}
}
