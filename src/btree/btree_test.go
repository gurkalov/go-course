package btree

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestSearch(t *testing.T) {
	nodes := []string{"A", "D", "F", "L", "N", "P", "H"}

	rootTree := &BTree{nil, 2}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range nodes {
		elem, i := rootTree.Search(rootTree.root, node)
		if elem == nil {
			t.Errorf("Not found element: %s", node)
		}

		if elem.Key[i] != node {
			t.Errorf("Not correct element expected: %s, but found: %s", node, elem.Key[i])
		}
	}
}

func TestSearch2(t *testing.T) {
	nodes := []string{"F", "S", "Q", "K", "C", "L", "H", "T", "V", "W", "M", "R", "N", "P", "A", "B", "X", "Y", "D", "Z", "E"}

	rootTree := &BTree{nil, 2}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range nodes {
		elem, i := rootTree.Search(rootTree.root, node)
		if elem == nil {
			t.Errorf("Not found element: %s", node)
		}

		if elem.Key[i] != node {
			t.Errorf("Not correct element expected: %s, but found: %s", node, elem.Key[i])
		}
	}
}

func TestDelete(t *testing.T) {
	nodes := []string{"A", "D", "F", "L", "N", "P", "H"}
	deleteNodes := []string{"A", "F", "N", "P", "H"}

	rootTree := &BTree{nil, 2}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range nodes {
		elem, i := rootTree.Search(rootTree.root, node)
		if elem == nil {
			t.Errorf("Not found element: %s", node)
		}

		if elem.Key[i] != node {
			t.Errorf("Not correct element expected: %s, but found: %s", node, elem.Key[i])
		}
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}


}