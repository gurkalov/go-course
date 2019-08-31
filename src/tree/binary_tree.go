package tree

import "course/stack"

type BinaryTree struct {
	Tree
	root *Node
}

func (T *BinaryTree) Insert (key int) *Node {
	z := &Node{nil, nil, nil, key, 0}

	var y *Node
	x := T.root
	for x != nil {
		y = x
		if z.Key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}

	z.P = y

	if y == nil {
		T.root = z
	} else if z.Key < y.Key {
		y.Left = z
	} else {
		y.Right = z
	}

	return z
}

func (T *BinaryTree) transplant (u, v *Node) {
	if u.P == nil {
		T.root = v
	} else if u == u.P.Left {
		u.P.Left = v
	} else {
		u.P.Right = v
	}

	if v != nil {
		v.P = u.P
	}
}

func (T *BinaryTree) InOrder(x *Node) []int {
	var items []int
	stack := stack.Stack{}

	for x != nil || stack.GetSize() > 0 {
		for x != nil {
			stack.Push(x)
			x = x.Left
		}

		pop := stack.Pop()
		x = pop.(*Node)

		items = append(items, x.Key)
		x = x.Right
	}

	return items
}

func (T *BinaryTree) PreOrder (x *Node) []int {
	var items []int
	stack := stack.Stack{}

	if x != nil {
		stack.Push(x)
	}

	for stack.GetSize() > 0 {
		pop := stack.Pop()
		x = pop.(*Node)

		items = append(items, x.Key)

		if x.Right != nil {
			stack.Push(x.Right)
		}

		if x.Left != nil {
			stack.Push(x.Left)
		}
	}

	return items
}

func (T *BinaryTree) Find (key int) *Node {
	x := T.root

	for x != nil && key != x.Key {
		if key < x.Key {
			x = x.Left
		} else {
			x = x.Right
		}
	}
	return x
}

func (T *BinaryTree) Min (x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func (T *BinaryTree) Max (x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

func (T *BinaryTree) Delete (key int) {
	z := T.Find(key)
	if z == nil {
		return
	}

	if z.Left == nil {
		T.transplant(z, z.Right)
	} else if z.Right == nil {
		T.transplant(z, z.Left)
	} else {
		y := T.Min(z.Right)
		if y.P != z {
			T.transplant(y, y.Right)
			y.Right = z.Right
			y.Right.P = y
		}
		T.transplant(z, y)
		y.Left = z.Left
		y.Left.P = y
	}
}
