package tree

import (
	"fmt"
	"reflect"
	"testing"
)

func TestAvlMax(t *testing.T) {
	nodes := []int{5, 7, 6, 9, 3, 2, 1, 4, 8}
	rootTree := &AvlTree{}
	for _, node := range nodes {
		rootTree.Insert(node)
	}

	if 9 != rootTree.Max(rootTree.root).Key {
		t.Errorf("Max element not uqual: %d", 9)
	}
}

func TestAvlMin(t *testing.T) {
	nodes := []int{5, 7, 6, 9, 3, 2, 1, 4, 8}
	rootTree := &AvlTree{}
	for _, node := range nodes {
		rootTree.Insert(node)
	}

	if 1 != rootTree.Min(rootTree.root).Key {
		t.Errorf("Min element not uqual: %d", 1)
	}
}

func TestAvlInOrder(t *testing.T) {
	nodes := []int{5, 7, 6, 9, 3, 2, 1, 4, 8}
	rootTree := &AvlTree{}
	for _, node := range nodes {
		rootTree.Insert(node)
	}

	result := rootTree.InOrder(rootTree.root)
	if !reflect.DeepEqual(result, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}) {
		fmt.Println(result)
		t.Error("Error")
	}
}

func TestAvlPreOrder(t *testing.T) {
	nodes := []int{5, 7, 6, 9, 3, 2, 1, 4, 8}
	rootTree := &AvlTree{}
	for _, node := range nodes {
		rootTree.Insert(node)
	}

	result := rootTree.PreOrder(rootTree.root)
	if !reflect.DeepEqual(result, []int{6, 3, 2, 1, 5, 4, 8, 7, 9}) {
		fmt.Println(result)
		t.Error("Error")
	}
}

func TestAvlFind(t *testing.T) {
	nodes := []int{5, 7, 6, 9, 3, 2, 1, 4, 8}
	rootTree := &AvlTree{}
	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range nodes {
		if node != rootTree.Find(node).Key {
			t.Errorf("Not found element: %d", node)
		}
	}
}

func TestAvlDelete(t *testing.T) {
	nodes := []int{5, 7, 6, 9, 3, 2, 1, 4, 8}
	deletedNodes := []int{5, 9, 1, 8}
	rootTree := &AvlTree{}
	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range nodes {
		if node != rootTree.Find(node).Key {
			t.Errorf("Not found element: %d", node)
		}
	}

	for _, node := range deletedNodes {
		if node != rootTree.Find(node).Key {
			t.Errorf("Not found element: %d", node)
		}
		rootTree.Delete(node)
		if rootTree.Find(node) != nil {
			t.Errorf("Found deleted element: %d", node)
		}
	}

	result := rootTree.InOrder(rootTree.root)
	if !reflect.DeepEqual(result, []int{2, 3, 4, 6, 7}) {
		fmt.Println(result)
		t.Error("Error")
	}

	result = rootTree.PreOrder(rootTree.root)
	if !reflect.DeepEqual(result, []int{6, 3, 2, 4, 7}) {
		fmt.Println(result)
		t.Error("Error")
	}
}
