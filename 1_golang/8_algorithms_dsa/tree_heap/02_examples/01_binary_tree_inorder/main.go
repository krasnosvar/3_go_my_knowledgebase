package main

import "fmt"

type Node struct {
	Val  int
	L, R *Node
}

func inorder(n *Node) {
	if n == nil {
		return
	}
	inorder(n.L)
	fmt.Print(n.Val, " ")
	inorder(n.R)
}

func main() {
	root := &Node{Val: 2, L: &Node{Val: 1}, R: &Node{Val: 3}}
	inorder(root)
	fmt.Println()
}
