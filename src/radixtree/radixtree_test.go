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
		"e": &RadixTree{false,"e", map[string]*RadixTree{
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

	if !reflect.DeepEqual(*tree.C["a"], RadixTree{true,"amber", nil}) {
		t.Error("Error!")
	}
}

func TestSearch(t *testing.T) {
	nodes := []string{"test", "team", "amber", "testjob", "testwork", "testart", "testworld"}
	notLeafNodes := []string{"te", "testwor"}
	wrongNodes := []string{"focus", "tea", "t", "testj", "a", "am", "testworker", "testworlord"}

	tree := RadixTree{}
	for _, node := range nodes {
		tree.Insert(node)
	}

	for _, node := range nodes {
		find, _ := tree.Search(node)
		if find != node {
			t.Errorf("Not found element: %s", node)
		}
	}

	for _, node := range notLeafNodes {
		find, _ := tree.Search(node)
		if find != "" {
			t.Errorf("Found element: %s, but it isn't leaf %s", node, find)
		}
	}

	for _, node := range wrongNodes {
		find, _ := tree.Search(node)
		if find != "" {
			t.Errorf("Found element: %s, but it is wrong %s", node, find)
		}
	}
}

func TestLongestPrefix(t *testing.T) {
	nodes := []string{"test", "team", "amber", "testjob", "testwork", "testart", "testworld"}
	longestPrefix := map[string]string{
		"teamer": "team",
		"testjob": "testjob",
		"testjobber": "testjob",
		"teambet": "team",
		"testing": "test",
		"test": "test",
		"testworing": "test",
		"testworker": "testwork",
		"testworlord": "test",
		"tea": "",
	}

	tree := RadixTree{}
	for _, node := range nodes {
		tree.Insert(node)
	}

	for key, prefix := range longestPrefix {
		find, _ := tree.LongestPrefix(key)
		if find != prefix {
			t.Errorf("Error longest prefix: %s for %s, but found: %s", prefix, key, find)
		}
	}
}
