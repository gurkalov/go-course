package matrix

import "math/rand"

type Matrix [][]int

func CreateMatrix(n int) Matrix {
	matrix := make(Matrix, n)
	for i := 0; i < n; i++ {
		matrix[i] = make([]int, n)
		for j := 0; j < n; j++ {
			matrix[i][j] = (rand.Int() % 8) + 1
		}
	}
	return matrix
}

func RowMulty (a Matrix) int {
	n := len(a)
	res := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res *= a[i][j]
		}
	}

	return res
}

func ColMulty (a Matrix) int {
	n := len(a)
	res := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			res *= a[j][i]
		}
	}

	return res
}
