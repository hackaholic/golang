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

func inOrderTraversal(root *BinTree) {
	// left-> root -> right =>>> nodes in increasing order
	// right -> root -> left =>>> nodes in decresing order
	if root == nil {
		return
	}
	inOrderTraversal(root.left)
	fmt.Print(root.value, " ")
	inOrderTraversal(root.right)
}

func walk(root *BinTree, ch chan int) {
	//println("walking....")
	//time.Sleep(time.Second)
	if root == nil {
		return
	}
	walk(root.left, ch)
	ch <- root.value
	walk(root.right, ch)
}

func compare(ch chan bool, ch1 chan int, ch2 chan int) {
	for {
		v1, ok1 := <-ch1 // channel returns value and bool(true/false)
		v2, ok2 := <-ch2
		fmt.Println("ch1: ", ch1, ok1, v1, "ch2", ch2, ok2, v2)
		if !ok1 || !ok2 {
			ch <- ok1 == ok2
			break
		}
		if v1 != v2 {
			fmt.Println("Not Equal on value")
			ch <- false
			break
		}
	}
}

func main() {
	arr_int1 := []int{5, 2, 6, 7, 1, 9, 0, 21, 67, 89, 12, 31, -1, 22}
	arr_int2 := []int{5, 2, 6, 7, 1, 0, 21, 67, 89, 12, 31, -1, 22, 9}

	//arr_int1 := []int{1, 2, 3, 4}
	//arr_int2 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11}

	fmt.Println("arr_int1:", arr_int1)
	println("Constructing Binary tree...")
	head1 := constructBinarySearchTree(arr_int1)
	fmt.Println("head1:", head1)
	println("In Order Traversal")
	inOrderTraversal(head1)
	println()
	fmt.Println("arr_int2:", arr_int1)
	println("Constructing Binary tree...")
	head2 := constructBinarySearchTree(arr_int2)
	fmt.Println("head:", head2)
	println("In Order Traversal")
	inOrderTraversal(head2)

	println()

	ch, ch1, ch2 := make(chan bool), make(chan int), make(chan int)
	go func() {
		walk(head1, ch1)
		close(ch1)
	}()
	go func() {
		walk(head2, ch2)
		close(ch2)
	}()
	fmt.Println("Comparing Binary Tree")
	go compare(ch, ch1, ch2)
	flag := <-ch
	fmt.Println(flag)
	close(ch)
}
