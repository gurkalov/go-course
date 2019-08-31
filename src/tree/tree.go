package tree

type Tree interface {
	Insert (key int) *Node
	Find (key int) *Node
	Delete (key int)
	InOrder(x *Node) []int
	PreOrder (x *Node) []int
	Min (x *Node) *Node
	Max (x *Node) *Node
}

type Node struct {
	P *Node
	Left *Node
	Right *Node
	Key int
	height uint
}
