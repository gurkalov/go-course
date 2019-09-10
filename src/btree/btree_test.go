package btree

import (
	"os"
	"reflect"
	"strconv"
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
	deleteNodes := []string{"A", "F", "N", "P", "H", "L", "D"}

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

func TestDeleteE(t *testing.T) {
	nodes := []string{"P", "C", "T", "G", "M", "X", "A", "B", "D", "Q", "R", "U", "V", "E", "F", "J", "K", "Y", "Z", "N", "O"}
	deleteNodes := []string{"E"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:6], []string{"C", "G", "M", "R", "V"}) {
		t.Error("Error, expected: [C G M R V]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"D", "F"}) {
		t.Error("Error, expected: [D F]")
	}
}

func TestDeleteD(t *testing.T) {
	nodes := []string{"P", "C", "T", "G", "M", "X", "A", "B", "D", "Q", "R", "U", "V", "E", "F", "J", "K", "Y", "Z", "N", "O"}
	deleteNodes := []string{"D"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:3], []string{"C", "G"}) {
		t.Error("Error, expected: [C G]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"E", "F"}) {
		t.Error("Error, expected: [E F]")
	}
}

func TestDeleteL(t *testing.T) {
	nodes := []string{"A", "B", "C", "D", "E", "F", "G", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V"}
	deleteNodes := []string{"L"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"P", "S"}) {
		t.Error("Error, expected: [P S]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].C[1].Key[1:5], []string{"K", "M", "N", "O"}) {
		t.Error("Error, expected: [K M N O]")
	}
}

func TestDeleteK(t *testing.T) {
	nodes := []string{"A", "B", "C", "D", "E", "F", "G", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T", "U", "V"}
	deleteNodes := []string{"K"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"P", "S"}) {
		t.Error("Error, expected: [P S]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].C[1].Key[1:5], []string{"L", "M", "N", "O"}) {
		t.Error("Error, expected: [L M N O]")
	}
}

func TestDeleteX(t *testing.T) {
	nodes := []string{"P", "C", "G", "T", "X", "A", "B", "E", "F", "J", "K", "L", "N", "O", "Q", "R", "U", "V", "Y", "Z", "H", "D", "I", "W"}
	deleteNodes := []string{"X"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:4], []string{"P", "T", "W"}) {
		t.Error("Error, expected: [P T W]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].C[3].Key[1:3], []string{"U", "V"}) {
		t.Error("Error, expected: [U V]")
	}
}

func TestDeleteT(t *testing.T) {
	nodes := []string{"P", "C", "G", "T", "X", "A", "B", "E", "F", "J", "K", "L", "N", "O", "Q", "R", "U", "V", "Y", "Z", "H", "D", "I", "W"}
	deleteNodes := []string{"T"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:4], []string{"P", "U", "X"}) {
		t.Error("Error, expected: [P U X]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].C[3].Key[1:3], []string{"V", "W"}) {
		t.Error("Error, expected: [V W]")
	}
}

func TestDeleteP(t *testing.T) {
	nodes := []string{"P", "C", "G", "T", "X", "A", "B", "E", "F", "J", "K", "L", "N", "O", "Q", "R", "U", "V", "Y", "Z", "H", "D", "I", "W"}
	deleteNodes := []string{"P"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"T", "X"}) {
		t.Error("Error, expected: [T X]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].C[1].Key[1:5], []string{"N", "O", "Q", "R"}) {
		t.Error("Error, expected: [N O Q R]")
	}
}

func TestDeleteRecurse(t *testing.T) {
	nodes := []string{"P", "C", "T", "G", "M", "X", "A", "B", "D", "Q", "R", "U", "V", "E", "F", "J", "K", "Y", "Z", "N", "O"}
	deleteNodes := []string{"Z", "Y", "X"}

	rootTree := &BTree{nil, 2}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	deleteNodes = []string{"R", "Q"}
	for _, node := range deleteNodes {
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}

	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"K", "O"}) {
		t.Error("Error, expected: [K O]")
	}

	if !reflect.DeepEqual(rootTree.root.C[2].C[3].Key[1:2], []string{"P"}) {
		t.Error("Error, expected: [P]")
	}
}

func TestDeleteTurn(t *testing.T) {
	nodes := []string{"A", "D", "F", "L", "B", "C", "E", "H", "J", "G", "I", "K", "N", "M", "Q", "R", "S", "T", "P", "X", "Y", "Z"}

	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for _, node := range nodes {
		rootTree.Insert(node)
	}

	node := "Z"
	rootTree.Delete(rootTree.root, node)
	elem, _ := rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:6], []string{"D", "H", "K", "N", "R"}) {
		t.Error("Error, expected: [D H K N R]")
	}

	node = "R"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:6], []string{"D", "H", "K", "N", "S"}) {
		t.Error("Error, expected: [D H K N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[6].Key[1:4], []string{"T", "X", "Y"}) {
		t.Error("Error, expected: [T X Y]")
	}

	node = "H"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:6], []string{"D", "G", "K", "N", "S"}) {
		t.Error("Error, expected: [D G K N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"E", "F"}) {
		t.Error("Error, expected: [E F]")
	}

	node = "K"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:5], []string{"D", "G", "N", "S"}) {
		t.Error("Error, expected: [D G N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[3].Key[1:5], []string{"I", "J", "L", "M"}) {
		t.Error("Error, expected: [I J L M]")
	}

	node = "E"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:5], []string{"C", "G", "N", "S"}) {
		t.Error("Error, expected: [D G N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"D", "F"}) {
		t.Error("Error, expected: [D F]")
	}

	node = "F"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:5], []string{"C", "I", "N", "S"}) {
		t.Error("Error, expected: [C I N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[3].Key[1:4], []string{"J", "L", "M"}) {
		t.Error("Error, expected: [J L M]")
	}

	node = "D"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:5], []string{"C", "J", "N", "S"}) {
		t.Error("Error, expected: [C J N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:3], []string{"G", "I"}) {
		t.Error("Error, expected: [G I]")
	}

	node = "G"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:4], []string{"C", "N", "S"}) {
		t.Error("Error, expected: [C N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[2].Key[1:5], []string{"I", "J", "L", "M"}) {
		t.Error("Error, expected: [I J L M]")
	}

	node = "T"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:4], []string{"C", "N", "S"}) {
		t.Error("Error, expected: [C N S]")
	}
	if !reflect.DeepEqual(rootTree.root.C[4].Key[1:3], []string{"X", "Y"}) {
		t.Error("Error, expected: [X Y]")
	}

	node = "X"
	rootTree.Delete(rootTree.root, node)
	elem, _ = rootTree.Search(rootTree.root, node)
	if elem != nil {
		t.Errorf("Found deleted element: %s", node)
	}

	if !reflect.DeepEqual(rootTree.root.Key[1:3], []string{"C", "N"}) {
		t.Error("Error, expected: [C N]")
	}
	if !reflect.DeepEqual(rootTree.root.C[3].Key[1:5], []string{"P", "Q", "S", "Y"}) {
		t.Error("Error, expected: [P Q S Y]")
	}
}

func TestDeleteLargeTree2(t *testing.T) {
	rootTree := &BTree{nil, 2}
	rootTree.Create()

	for i := 1000; i < 10000; i++ {
		rootTree.Insert(strconv.Itoa(i))
	}

	var node string
	for i := 1000; i < 10000; i++ {
		node = (strconv.Itoa(i))
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}
}

func TestDeleteLargeTree3(t *testing.T) {
	rootTree := &BTree{nil, 3}
	rootTree.Create()

	for i := 1000; i < 10000; i++ {
		rootTree.Insert(strconv.Itoa(i))
	}

	var node string
	for i := 1000; i < 10000; i++ {
		node = (strconv.Itoa(i))
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}
}

func TestDeleteLargeTree4(t *testing.T) {
	rootTree := &BTree{nil, 4}
	rootTree.Create()

	for i := 1000; i < 10000; i++ {
		rootTree.Insert(strconv.Itoa(i))
	}

	var node string
	for i := 1000; i < 10000; i++ {
		node = (strconv.Itoa(i))
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}
}

func TestDeleteLargeTree5(t *testing.T) {
	rootTree := &BTree{nil, 5}
	rootTree.Create()

	for i := 10000; i < 100000; i++ {
		rootTree.Insert(strconv.Itoa(i))
	}

	var node string
	for i := 10000; i < 100000; i++ {
		node = (strconv.Itoa(i))
		rootTree.Delete(rootTree.root, node)
		elem, _ := rootTree.Search(rootTree.root, node)
		if elem != nil {
			t.Errorf("Found deleted element: %s", node)
		}
	}
}