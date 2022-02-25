package main

import (
	"reflect"
	"testing"
)

func TestMatrixMultiplication(t *testing.T) {
	matrix1 := [][]int{{1, 3, 0}}

	matrix2 := [][]int{{2}, {4}, {1}}

	got := MatrixMultiplication(matrix1, matrix2)
	want := [][]int{{14}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Expected %v, got %v", want, got)
	}
}
