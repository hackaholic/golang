package main

import (
	"fmt"
)

type BinTree struct {
	value int
	left  *BinTree
	right *BinTree
}

//{5, 2, 6, 7, 1, 9}
func insertBinNode(root *BinTree, node *BinTree) {
	if root == nil {
		return
	}
	if root.value >= node.value {
		if root.left == nil {
			root.left = node
		} else {
			insertBinNode(root.left, node)
		}
	}
	if root.value < node.value {
		if root.right == nil {
			root.right = node
		} else {
			insertBinNode(root.right, node)
		}
	}
}

func constructBinarySearchTree(arr []int) *BinTree {
	/*
	  Construct Binary tree from the array
	*/
	var head, node *BinTree
	for i, v := range arr {
		node = &BinTree{value: v}
		if i == 0 {
			head = node
			continue
		}
		insertBinNode(head, node)
	}
	return head
}

func postOrderTraversal(root *BinTree) {
	// left -> right -> root
	if root == nil {
		return
	}
	postOrderTraversal(root.left)
	postOrderTraversal(root.right)
	fmt.Print(root.value, " ")
}

func preOrderTraversal(root *BinTree) {
	// root -> left -> right
	if root == nil {
		return
	}
	fmt.Print(root.value, " ")
	preOrderTraversal(root.left)
	preOrderTraversal(root.right)
}

var inorder []BinTree

//  binary search trees (BST), Inorder traversal gives nodes in increasing order
func inOrderTraversal(root *BinTree) {
	// left-> root -> right =>>> nodes in increasing order
	// right -> root -> left =>>> nodes in decresing order
	if root == nil {
		return
	}
	inOrderTraversal(root.left)
	fmt.Print(root.value, " ")
	inorder = append(inorder, *root)
	inOrderTraversal(root.right)
}

func getHeight(root *BinTree) int64 {

	if root == nil {
		return -1
	}
	heightL := getHeight(root.left) + 1
	heightR := getHeight(root.right) + 1
	if heightL > heightR {
		return heightL
	}
	return heightR

	/*
		      h, node
		      0, 1  -> 2^0
			  1, 2  -> 2^1
			  2, 4  -> 2^2
			  3, 8  -> 2^3
			  h-1,   -> 2^h
			  h,     -> 2^h+1

			           0
				   0        0
				0     0   0	   0
			   0  0 0  0 0 0  0 0
		    maximum no of nodes at height h
			n = 2^h

			total number of nodes possible in binary tree of heigh h
		    a + ar + ar^2 + ar^3 ==> a*(1-r^n)/1-r
		    1 + 2  +  4   +  8 .... =>
			a=1, r=2
			n = 1(1-2^h+1)/1-2 => 2^h+1 - 1
			n = 2^(h+1) - 1

			minimum height possible for n nodes
			n+1 = 2^(h+1)
			log2(n+1) = h+1
			h = log2(n+1) - 1


			internal_nodes = total_nodes - leaf_nodes
			internal_nodes = 2^(h+1) -1  - 2^h => 2.2^h - 2^h -1 => 2^h(2-1) -1 => 2^h - 1
			internal_nodes = 2^h - 1

	*/

}

func countNodes(root *BinTree) int {
	if root == nil {
		return 0
	}
	return 1 + (countNodes(root.left) + countNodes(root.right))
}

func countLeafNodes(root *BinTree) int {
	if root == nil {
		return 0
	}
	if root.left == nil && root.right == nil {
		return 1
	}
	return countLeafNodes(root.left) + countLeafNodes(root.right)
}

func getLeafNodes(root *BinTree) []BinTree {
	if root != nil {
		if root.left == nil && root.right == nil {
			return []BinTree{*root}
		}
		var temp []BinTree
		for _, v := range getLeafNodes(root.left) {
			temp = append(temp, v)
		}
		for _, v := range getLeafNodes(root.right) {
			temp = append(temp, v)
		}
		return temp
	}

	return nil
}

func getMaxNode(root *BinTree) BinTree {
	if root.right == nil && root.left == nil {
		return *root
	}
	return getMaxNode(root.right)
}

func searchNode(root *BinTree, x int) bool {
	if root == nil {
		return false
	}
	if root.value == x {
		return true
	}
	if root.value > x {
		return false || searchNode(root.left, x)
	} else {
		return false || searchNode(root.right, x)
	}
}

func main() {
	arr_int := []int{5, 2, 6, 7, 1, 9, 0, 21, 67, 89, 12, 31, -1, 22}
	/*

		             5
					2 6
				   1   7
				  0      9
			   -1         21
						12   67
						  31   89
		                22
	*/
	fmt.Println("arr_int:", arr_int)
	println("Constructing Binary tree...")
	head := constructBinarySearchTree(arr_int)
	fmt.Println("head:", head)
	println("Post Order Traversal")
	postOrderTraversal(head)
	println("\nPre Order Traversal")
	preOrderTraversal(head)
	println("\nIn Order Traversal")
	inOrderTraversal(head)
	println("\n nodes in inorder sequence:")
	fmt.Println(inorder)
	// height of a tree
	// Height: Imagine measuring a person's height, we do it from toe to head (leaf to root).
	// 2. Depth: Imagine measuring depth of a sea, we do it from earth's surface to ocean bed (root to leaf)

	fmt.Println("height:", getHeight(head))
	fmt.Println("Leaf nodes count: ", countLeafNodes(head))
	fmt.Println("Leaf Nodes:", getLeafNodes(head))
	fmt.Println("Max Node:", getMaxNode(head))
	fmt.Println("Total Nodes:", countNodes(head))
	fmt.Println("searching 6:", searchNode(head, 6))
	fmt.Println("searching 33:", searchNode(head, 33))
	//fmt.Println(flag)
}
