package main

import (
	"reflect"
	"testing"
)

func TestMatrixMultiplication(t *testing.T) {
	t.Run("multiply 1x3 by 3x1 matrix", func(t *testing.T) {
		matrix1 := [][]int{{1, 3, 0}}

		matrix2 := [][]int{{2}, {4}, {1}}

		got, err := MatrixMultiplication(matrix1, matrix2)
		want := [][]int{{14}}

		if err != nil {
			t.Errorf("Received an error %q when it shouldn't", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("multiply 3x2 by 2x3 matrix", func(t *testing.T) {
		m1 := [][]int{{3, 2, 1}, {1, 0, 2}}
		m2 := [][]int{{1, 2}, {0, 1}, {4, 0}}

		got, err := MatrixMultiplication(m1, m2)
		want := [][]int{{7, 8}, {9, 2}}

		if err != nil {
			t.Errorf("Received an error %q when it shouldn't", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("multiply 4x4 by 4x4 matrix", func(t *testing.T) {
		m1 := [][]int{{5, 2, 6, 1}, {0, 6, 2, 0}, {3, 8, 1, 4}, {1, 8, 5, 6}}
		m2 := [][]int{{7, 5, 8, 0}, {1, 8, 2, 6}, {9, 4, 3, 8}, {5, 3, 7, 9}}

		got, err := MatrixMultiplication(m1, m2)
		want := [][]int{{96, 68, 69, 69}, {24, 56, 18, 52}, {58, 95, 71, 92}, {90, 107, 81, 142}}

		if err != nil {
			t.Errorf("Received an error %q when it shouldn't", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("expect error when matrices rows and columns don't match", func(t *testing.T) {
		m1 := [][]int{{3, 2, 1}, {1, 0, 2}}
		m2 := [][]int{{3}, {1}}

		_, err := MatrixMultiplication(m1, m2)

		if err == nil {
			t.Errorf("Expected error, but got %v", err)
		}
	})
}

func TestMultiThreadMatrixMultiplication(t *testing.T) {
	t.Run("multiply a small sqaure matrix", func(t *testing.T) {
		m1 := [][]int{{5, 2, 6, 1}, {0, 6, 2, 0}, {3, 8, 1, 4}, {1, 8, 5, 6}}
		m2 := [][]int{{7, 5, 8, 0}, {1, 8, 2, 6}, {9, 4, 3, 8}, {5, 3, 7, 9}}

		got, err := MultiThreadedMatrixMultiplication(m1, m2)
		want := [][]int{{96, 68, 69, 69}, {24, 56, 18, 52}, {58, 95, 71, 92}, {90, 107, 81, 142}}

		if err != nil {
			t.Errorf("Received an error %q when it shouldn't", err)
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("Expected %v, got %v", want, got)
		}
	})

	t.Run("expect error when matrices rows and columns don't match", func(t *testing.T) {
		m1 := [][]int{{3, 2, 1}, {1, 0, 2}}
		m2 := [][]int{{3}, {1}}

		_, err := MultiThreadedMatrixMultiplication(m1, m2)

		if err == nil {
			t.Errorf("Expected error, but got %v", err)
		}
	})
}

func BenchmarkMatrixMultiplication(b *testing.B) {
	m1 := [][]int{{3, 2, 1}, {1, 0, 2}}
	m2 := [][]int{{1, 2}, {0, 1}, {4, 0}}
	for i := 0; i < b.N; i++ {
		MatrixMultiplication(m1, m2)
	}
}

func BenchmarkSqaureMatrixMultiplication(b *testing.B) {
	m1 := [][]int{{5, 2, 6, 1}, {0, 6, 2, 0}, {3, 8, 1, 4}, {1, 8, 5, 6}}
	m2 := [][]int{{7, 5, 8, 0}, {1, 8, 2, 6}, {9, 4, 3, 8}, {5, 3, 7, 9}}

	for i := 0; i < b.N; i++ {
		MatrixMultiplication(m1, m2)
	}
}

func BenchmarkMultiThreadedMatrixMultiplication(b *testing.B) {
	m1 := [][]int{{5, 2, 6, 1}, {0, 6, 2, 0}, {3, 8, 1, 4}, {1, 8, 5, 6}}
	m2 := [][]int{{7, 5, 8, 0}, {1, 8, 2, 6}, {9, 4, 3, 8}, {5, 3, 7, 9}}

	for i := 0; i < b.N; i++ {
		MultiThreadedMatrixMultiplication(m1, m2)
	}
}
