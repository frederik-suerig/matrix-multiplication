package main

import "errors"

func MatrixMultiplication(matrix1, matrix2 [][]int) ([][]int, error) {
	if checkMatrixSize(matrix1, matrix2) {
		return nil, errors.New("invalid input, those matrixes can't be multiplied")
	}

	columns := len(matrix2[0])

	result := [][]int{}

	for i := 0; i < len(matrix1); i++ {
		result = append(result, make([]int, columns))
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}

	}

	return result, nil
}

func checkMatrixSize(m1, m2 [][]int) bool {
	rows := len(m1)
	cols := len(m2[0])

	return rows != cols
}
