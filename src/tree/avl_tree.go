package tree

import "course/stack"

type AvlTree struct {
	Tree
	root *Node
}

func (T *AvlTree) Insert (key int) *Node {
	z := &Node{nil, nil, nil, key, 1}

	if T.root == nil {
		T.root = z
		return z
	}

	T.root = T.insert(T.root, z)
	return T.root
}

func (T *AvlTree) insert(p *Node, new *Node) *Node {
	if p == nil {
		return new
	}

	key := new.Key
	if key < p.Key {
		p.Left = T.insert(p.Left, new)
	} else {
		p.Right = T.insert(p.Right, new)
	}

	return p.Balance()
}

func (T *AvlTree) transplant (u, v *Node) {
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

func (T *AvlTree) InOrder(x *Node) []int {
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

func (T *AvlTree) PreOrder (x *Node) []int {
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

func (T *AvlTree) Find (key int) *Node {
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

func (x *Node) Height () uint {
	if x == nil {
		return 0
	}
	return x.height
}

func (x *Node) BFactor () int {
	return int(x.Right.Height() - x.Left.Height())
}

func (x *Node) FixHeight () {
	hLeft := x.Left.Height()
	hRight := x.Right.Height()

	if hLeft > hRight {
		x.height = hLeft + 1
	} else {
		x.height = hRight + 1
	}
}

func (p *Node) RotateRight() *Node {
	q := p.Left
	p.Left = q.Right
	q.Right = p
	p.FixHeight()
	q.FixHeight()

	return q
}

func (q *Node) RotateLeft() *Node {
	p := q.Right
	q.Right = p.Left
	p.Left = q
	q.FixHeight()
	p.FixHeight()

	return p
}

func (p *Node) Balance() *Node {
	p.FixHeight()

	if p.BFactor() == 2 {
		if p.Right.BFactor() < 0 {
			p.Right = p.Right.RotateRight()
		}
		return p.RotateLeft()
	}

	if p.BFactor() == -2 {
		if p.Left.BFactor() > 0 {
			p.Left = p.Left.RotateLeft()
		}
		return p.RotateRight()
	}

	return p
}

func (T *AvlTree) RemoveMin (p *Node) *Node {
	if p.Left == nil {
		return p.Right
	}

	p.Left = T.RemoveMin(p.Left)

	return p.Balance()
}

func (T *AvlTree) Min (x *Node) *Node {
	for x.Left != nil {
		x = x.Left
	}
	return x
}

func (T *AvlTree) Max (x *Node) *Node {
	for x.Right != nil {
		x = x.Right
	}
	return x
}

func (T *AvlTree) remove (p *Node, key int) *Node {
	if p == nil {
		return nil
	}

	if key < p.Key {
		p.Left = T.remove(p.Left, key)
	} else if key > p.Key {
		p.Right = T.remove(p.Right, key)
	} else {
		q := p.Left
		r := p.Right

		if r == nil {
			return q
		}

		min := T.Min(r)
		min.Right = T.RemoveMin(r)
		min.Left = q
		return min.Balance()
	}

	return p.Balance()
}

func (T *AvlTree) Delete (key int) {
	z := T.Find(key)
	if z == nil {
		return
	}

	T.remove(T.root, key)
}
