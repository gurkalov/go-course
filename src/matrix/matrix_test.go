package matrix

import (
	"os"
	"testing"
)

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}

func BenchmarkRow(b *testing.B) {
	m := CreateMatrix(100)
	for i := 0; i < b.N; i++ {
		RowMulty(m)
	}
}

func BenchmarkCol(b *testing.B) {
	m := CreateMatrix(100)
	for i := 0; i < b.N; i++ {
		ColMulty(m)
	}
}
