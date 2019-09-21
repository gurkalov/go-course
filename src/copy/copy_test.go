package copy

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func TestCopy1(t *testing.T) {
	CopyOne("data/file.txt", "data/copy1.txt")
	CopyTwo("data/file.txt", "data/copy2.txt")
}
