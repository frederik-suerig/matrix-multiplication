package main

import (
	"reflect"
	"testing"
)

func TestMatrixMultiplication(t *testing.T) {
	t.Run("multiply 1x3 by 3x1 matrix", func(t *testing.T) {
		matrix1 := [][]int{{1, 3, 0}}

		matrix2 := [][]int{{2}, {4}, {1}}

		got := MatrixMultiplication(matrix1, matrix2)
		want := [][]int{{14}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("multiply 3x2 by 2x3 matrix", func(t *testing.T) {
		m1 := [][]int{{3, 2, 1}, {1, 0, 2}}
		m2 := [][]int{{1, 2}, {0, 1}, {4, 0}}

		got := MatrixMultiplication(m1, m2)
		want := [][]int{{7, 8}, {9, 2}}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})
}
