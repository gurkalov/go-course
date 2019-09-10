package btree

type BTree struct {
	root *Node
	t uint
}

type Tree interface {
	Insert (key int) *Node
	Find (key int) *Node
	Delete (key int)
}

type Node struct {
	Leaf bool
	N uint
	Key []string
	C []*Node
}

func (T *BTree) Create () {
	x := &Node{}
	x.Leaf = true
	x.N = 0
	x.Key = make([]string, 2 * T.t)
	x.C = make([]*Node, 2 * T.t + 1)
	T.root = x
}

func (T *BTree) InsertNonfull (x *Node, k string) {
	i := x.N
	if x.Leaf {
		for i >= 1 && k < x.Key[i] {
			x.Key[i + 1] = x.Key[i]
			i = i - 1
		}
		x.Key[i + 1] = k
		x.N = x.N + 1
	} else {
		for i >= 1 && k < x.Key[i] {
			i = i - 1
		}
		i = i + 1
		if x.C[i].N == 2 * T.t - 1 {
			T.SplitChild(x, i)
			if k > x.Key[i] {
				i = i + 1
			}
		}
		T.InsertNonfull(x.C[i], k)
	}
}

func (T *BTree) SplitChild (x *Node, i uint) {
	z := &Node{}

	y := x.C[i]
	z.Leaf = y.Leaf
	z.N = T.t - 1
	z.Key = make([]string, 2 * T.t)
	z.C = make([]*Node, 2 * T.t + 1)

	for j := 1; j <= int(T.t) - 1; j++ {
		z.Key[j] = y.Key[j + int(T.t)]
	}

	if !y.Leaf {
		for j := 1; j <= int(T.t); j++ {
			z.C[j] = y.C[j + int(T.t)]
		}
	}

	y.N = T.t - 1
	for j := x.N + 1; j >= i + 1; j-- {
		x.C[j + 1] = x.C[j]
	}
	x.C[i + 1] = z

	for j := x.N; j >= i; j-- {
		x.Key[j + 1] = x.Key[j]
	}

	x.Key[i] = y.Key[T.t]
	// clear y part that moved to z
	for j := y.N + 1; j < 2 * T.t; j++ {
		y.Key[j] = ""
		y.C[j + 1] = nil
	}

	x.N = x.N + 1
}

func (T *BTree) Insert (k string) {
	r := T.root

	if r.N == 2 * T.t - 1 {
		s := &Node{}
		T.root = s
		s.Leaf = false
		s.N = 0
		s.C = make([]*Node, 2 * T.t + 1)
		s.Key = make([]string, 2 * T.t)
		s.C[1] = r
		T.SplitChild(s, 1)
		T.InsertNonfull(s, k)
	} else {
		T.InsertNonfull(r, k)
	}
}

func (T *BTree) Search (x *Node, k string) (*Node, uint) {
	i := uint(1)
	for i <= x.N && k > x.Key[i] {
		i = i + 1
	}
	if i <= x.N && k == x.Key[i] {
		return x, i
	} else if x.Leaf {
		return nil, 0
	} else {
		if x.C[i] == nil {
			return nil, 0
		}
		return T.Search(x.C[i], k)
	}
}

func (x *Node) RemoveFromLeaf (i uint) {
	for j := i; j < x.N; j++ {
		x.Key[j] = x.Key[j + 1]
	}
	x.Key[x.N] = ""
	for j := i; j <= x.N; j++ {
		x.C[j] = x.C[j + 1]
	}
	x.C[x.N + 1] = nil
	x.SetN(x.N - 1)
}

func (x *Node) MergePair (p, y *Node) {
	x.Key[x.N + 1] = p.Key[1]
	for j := uint(1); j <= y.N; j++ {
		x.Key[x.N + 1 + j] = y.Key[j]
	}

	for j := uint(1); j <= y.N + 1; j++ {
		x.C[x.N + 1 + j] = y.C[j]
	}

	x.SetN(x.N + y.N + 1)
	y = nil
}

func (x *Node) SetN (n uint) {
	start := uint(n + 2)
	if n == 0 {
		start = 1
	}
	for j := start; j < uint(len(x.C)); j++ {
		x.C[j] = nil
	}
	x.N = n
}

func (x *Node) MergeSibling (n, y *Node, i uint) {
	x.Key[x.N + 1] = n.Key[i]
	for k := uint(1); k <= y.N; k++ {
		y.Key[x.N + 1 + k] = y.Key[k]
	}

	for k := uint(1); k <= x.N; k++ {
		y.Key[k] = x.Key[k]
	}
	y.Key[x.N + 1] = n.Key[i]
	y.SetN(x.N + 1 + y.N)
	n.RemoveFromLeaf(i)
}

func (T *BTree) NeedMergeAll (n *Node) bool {
	if n == nil {
		return false
	}
	if n.N == 1 && n.C[1] != nil && n.C[1].N == T.t - 1 && n.C[2] != nil && n.C[2].N == T.t - 1 {
		return true
	}
	return false
}

func (T *BTree) TryUp (n *Node, i uint) bool {
	if n.C[i].N > T.t - 1 {
		n.Key[i] = n.C[i].Key[n.C[i].N]
		n.C[i].RemoveFromLeaf(n.C[i].N)
		return true
	}

	if n.C[i + 1].N > T.t - 1 {
		n.Key[i] = n.C[i + 1].Key[1]
		n.C[i + 1].RemoveFromLeaf(1)
		return true
	}

	return false
}

func (T *BTree) TryTurn (n *Node, i uint, key string) bool {
	if n.C[i] == nil || n.C[i].N > T.t - 1 {
		return false
	}

	j := uint(1)
	for j <= n.C[i].N && key > n.C[i].Key[j] {
		j = j + 1
	}

	if n.C[i].Key[j] == key {
		if i > 1 && n.C[i - 1] != nil && n.C[i - 1].N > T.t - 1 {
			n.C[i].Key[j] = n.Key[i - 1]
			n.Key[i - 1] = n.C[i - 1].Key[n.C[i - 1].N]
			n.C[i - 1].RemoveFromLeaf(n.C[i - 1].N)
			return true
		} else if n.C[i + 1] != nil && n.C[i + 1].N > T.t - 1 {
			n.C[i].RemoveFromLeaf(j)
			n.C[i].Key[n.C[i].N + 1] = n.Key[i]
			n.C[i].SetN(n.C[i].N + 1)
			n.Key[i] = n.C[i + 1].Key[1]
			n.C[i + 1].RemoveFromLeaf(1)
			return true
		}
	}

	return false
}

func (T *BTree) TryMerge (n *Node, i uint, key string) bool {
	if (n.C[i] != nil && n.C[i].N > T.t - 1) || (i <= n.N && n.C[i + 1] != nil && n.C[i + 1].N > T.t - 1) || (i > 1 && n.C[i - 1] != nil && n.C[i - 1].N > T.t - 1) {
		return false
	}

	if i <= n.N {
		n.C[i].Key[n.C[i].N+1] = n.Key[i]
		n.C[i].C[n.C[i].N+2] = n.C[i + 1].C[1]
		for j := uint(1); j <= n.C[i+1].N; j++ {
			n.C[i].Key[n.C[i].N+1+j] = n.C[i+1].Key[j]
			n.C[i].C[n.C[i].N+2+j] = n.C[i+1].C[j + 1]
		}

		n.C[i].SetN(n.C[i].N + 1 + n.C[i+1].N)
		j := uint(i)
		for ; j < n.N; j++ {
			n.Key[j] = n.Key[j+1]
			n.C[j+1] = n.C[j+2]
		}
		n.Key[j] = ""
		n.SetN(n.N - 1)
		T.Delete(n.C[i], key)
	} else {
		n.C[i - 1].Key[n.C[i].N + 1] = n.Key[i - 1]
		for j := uint(1); j <= n.C[i].N; j++ {
			n.C[i - 1].Key[n.C[i].N + 1 + j] = n.C[i].Key[j]
		}

		n.C[i - 1].SetN(n.C[i-1].N + 1 + n.C[i].N)
		j := uint(i)
		for ; j < n.N; j++ {
			n.Key[j] = n.Key[j+1]
			n.C[j+1] = n.C[j+2]
		}
		n.Key[j - 1] = ""
		n.SetN(n.N - 1)

		T.Delete(n.C[i - 1], key)
	}
	return true
}

func (T *BTree) Delete (n *Node, key string) {
	i := uint(1)
	for i <= n.N && key > n.Key[i] {
		i = i + 1
	}

	if !n.Leaf {
		if T.NeedMergeAll(n) {
			n.C[1].MergePair(n, n.C[2])
			if n == T.root {
				T.root = n.C[1]
			}
			T.Delete(n.C[1], key)
		} else if T.NeedMergeAll(n.C[i]) {
			n.C[i].C[1].MergePair(n.C[i], n.C[i].C[2])
			n.C[i] = n.C[i].C[1]

			T.Delete(n.C[i], key)
		} else if i <= n.N && key == n.Key[i] {
			if !T.TryUp(n, i) {
				n.C[i].MergeSibling(n, n.C[i + 1], i)
				T.Delete(n.C[i], key)
			}
		} else {
			if T.TryMerge(n, i, key) {
				return
			} else if T.TryTurn(n, i, key) {
				return
			}

			T.Delete(n.C[i], key)
		}
	} else {
		if i <= n.N && key == n.Key[i] {
			n.RemoveFromLeaf(i)
		}
	}
}
