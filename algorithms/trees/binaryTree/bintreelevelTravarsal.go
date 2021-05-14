package main

import (
	"fmt"
)

type BinTree struct {
	val   int
	left  *BinTree
	right *BinTree
}

func enqueue(arr []*BinTree, node *BinTree) []*BinTree {
	arr = append(arr, node)
	return arr
}

func dequeue(arr []*BinTree) []*BinTree {
	if arr == nil {
		return nil
	}
	if len(arr) == 1 {
		return nil
	}
	return arr[1:]
}

func insertNode(root *BinTree, node *BinTree) {
	if root == nil {
		return
	}

	if root.val >= node.val {
		if root.left == nil {
			root.left = node
		} else {
			insertNode(root.left, node)
		}
	}

	if root.val < node.val {
		if root.right == nil {
			root.right = node
		} else {
			insertNode(root.right, node)
		}
	}
}

func constructBinarySearchTree(arr []int) *BinTree {
	var head *BinTree
	for i, v := range arr {
		if i == 0 {
			head = &BinTree{v, nil, nil}
		} else {
			insertNode(head, &BinTree{v, nil, nil})
		}
	}
	return head
}

func inorder(root *BinTree) {
	if root == nil {
		return
	}
	inorder(root.left)
	print(root.val, " ")
	inorder(root.right)
}

func levelTraversal(root *BinTree) {
	var arr []*BinTree
	var node *BinTree
	var level int
	if root == nil {
		return
	}
	arr = enqueue(arr, root)
	for len(arr) > 0 {
		node = arr[0]
		arr = dequeue(arr)
		fmt.Print(node.val, " ")
		if node.left != nil {
			arr = enqueue(arr, node.left)
			level += 1
		}
		if node.right != nil {
			arr = enqueue(arr, node.right)
		}
	}
}

func main() {
	var head *BinTree
	arr_int := []int{5, 2, 6, 7, 1, 9, 0, 21, 67, 89, 12, 31, -1, 22, 1}
	fmt.Println("input array:", arr_int)
	fmt.Println("Constructing Binary Search Tree")
	head = constructBinarySearchTree(arr_int)
	print("Inorder Traversal: ")
	inorder(head)
	fmt.Println("\nWe will do level traversal in Binary search Tree")
	levelTraversal(head)
	println()
}
