package radixtree

import (
	"os"
	"reflect"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestBrokeKey(t *testing.T) {
	p, sA, sB := BrokeKey("testwork", "test")
	if p != "test" || sA != "work" || sB != "" {
		t.Errorf("Not correct broke string: [test] [work] [] , but found: [%s] [%s] [%s]", p, sA, sB)
	}

	p, sA, sB = BrokeKey("test", "team")
	if p != "te" || sA != "st" || sB != "am" {
		t.Errorf("Not correct broke string: [te] [st] [am] , but found: [%s] [%s] [%s]", p, sA, sB)
	}

	p, sA, sB = BrokeKey("t", "test")
	if p != "t" || sA != "" || sB != "est" {
		t.Errorf("Not correct broke string: [t] [] [est] , but found: [%s] [%s] [%s]", p, sA, sB)
	}

	p, sA, sB = BrokeKey("test", "t")
	if p != "t" || sA != "est" || sB != "" {
		t.Errorf("Not correct broke string: [t] [est] [] , but found: [%s] [%s] [%s]", p, sA, sB)
	}
}

func TestInsert(t *testing.T) {

	tree := RadixTree{}
	tree.Insert("test")
	tree.Insert("testjob")
	tree.Insert("testwork")
	tree.Insert("team")
	tree.Insert("teller")
	tree.Insert("tor")
	tree.Insert("tar")
	tree.Insert("tera")
	tree.Insert("teamer")

	expectedTree := RadixTree{false,"t", map[string]*RadixTree{
		"e": &RadixTree{true,"e", map[string]*RadixTree{
			"s": &RadixTree{true,"st", map[string]*RadixTree{
				"j": &RadixTree{true,"job", nil},
				"w": &RadixTree{true,"work", nil},
			}},
			"a": &RadixTree{true,"am", map[string]*RadixTree{
				"e": &RadixTree{true,"er", nil},
			}},
			"l": &RadixTree{true,"ller", nil},
			"r": &RadixTree{true,"ra", nil},
		}},
		"o": &RadixTree{true,"or", nil},
		"a": &RadixTree{true, "ar", nil},
	}}

	if !reflect.DeepEqual(tree, expectedTree) {
		t.Error("Error!")
	}
}

func TestInsertRootLevel(t *testing.T) {

	tree := RadixTree{}
	tree.Insert("test")
	tree.Insert("team")
	tree.Insert("amber")

	if !reflect.DeepEqual(tree.Key, "") {
		t.Error("Error!")
	}

	if !reflect.DeepEqual(*tree.C["a"], RadixTree{false,"amber", nil}) {
		t.Error("Error!")
	}
}
