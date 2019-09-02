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
		return T.Search(x.C[i], k)
	}
}

func (T *BTree) Delete (x *Node, k string) {
	i := uint(1)
	for i <= x.N && k > x.Key[i] {
		i = i + 1
	}
	if i <= x.N && k == x.Key[i] {
		x.Key[i] = ""
		return
	} else if x.Leaf {
		return
	} else {
		T.Delete(x.C[i], k)
	}
}
