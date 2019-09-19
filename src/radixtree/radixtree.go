package radixtree

import (
	"strings"
)

type RadixTree struct {
	Leaf bool
	Key string
	C map[string]*RadixTree
}

func BrokeKey (a, b string) (string, string, string) {
	byteA := []byte(a)
	byteB := []byte(b)
	lenByteA := len(byteA)
	lenByteB := len(byteB)
	minSize := lenByteA
	if lenByteB < minSize {
		minSize = lenByteB
	}

	i := 0
	for ; i < minSize; i++ {
		if byteA[i] != byteB[i] {
			return string(byteA[0:i]), string(byteA[i:]), string(byteB[i:])
		}
	}

	return string(byteA[0:i]), string(byteA[i:]), string(byteB[i:])
}

func FirstChar (str string) string {
	byte := []byte(str)
	return string(byte[0:1])
}

func (n *RadixTree) Insert (key string) {
	if n.Key == "" && n.C == nil {
		n.Key = key
		n.Leaf = true
	} else {
		if n.C == nil {
			n.C = make(map[string]*RadixTree)
		}

		commonPrefix, suffNewKey, suffNodeKey := BrokeKey(key, n.Key)
		if commonPrefix != n.Key {
			n.Key = suffNodeKey
			leaf := n.Leaf
			n.Leaf = false

			newParent := &RadixTree{}
			newParent.C = make(map[string]*RadixTree)
			newParent.Key = commonPrefix
			newParent.C[FirstChar(suffNewKey)] = &RadixTree{}
			newParent.C[FirstChar(suffNewKey)].Key = suffNewKey
			newParent.C[FirstChar(suffNewKey)].Leaf = true

			newParent.C[FirstChar(suffNodeKey)] = &RadixTree{leaf,n.Key, n.C}
			n.Key = newParent.Key
			n.C = newParent.C
		} else {
			suffixKey := strings.TrimPrefix(key, n.Key)
			char := FirstChar(suffixKey)

			if node, ok := n.C[char]; ok {
				node.Insert(suffixKey)
			} else {
				n.C[FirstChar(suffixKey)] = &RadixTree{}
				n.C[FirstChar(suffixKey)].Key = suffixKey
				n.C[FirstChar(suffixKey)].Leaf = true
			}
		}
	}
}

func (n *RadixTree) Search (key string) (string, bool) {
	if n.Leaf && n.Key == key {
		return n.Key, true
	}

	_, suffNewKey, _ := BrokeKey(key, n.Key)
	char := FirstChar(suffNewKey)

	if node, ok := n.C[char]; ok {
		find, ok := node.Search(suffNewKey)
		if ok {
			return n.Key + find, true
		}

		return find, false
	}

	return "", false
}

func (n *RadixTree) LongestPrefix (key string) (string, bool) {
	if n.Leaf && n.Key == key {
		return n.Key, true
	}

	commonPrefix, suffNewKey, _ := BrokeKey(key, n.Key)
	char := FirstChar(suffNewKey)

	if node, ok := n.C[char]; ok {
		find, ok := node.LongestPrefix(suffNewKey)
		if ok {
			return n.Key + find, true
		} else if n.Leaf {
			return n.Key, true
		}

		return find, false
	}  else if n.Leaf && n.Key == commonPrefix {
		return n.Key, true
	}

	return "", false
}
