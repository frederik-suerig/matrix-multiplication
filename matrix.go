package main

import "errors"

func MatrixMultiplication(matrix1, matrix2 [][]int) ([][]int, error) {
	rows := len(matrix1)
	cols := len(matrix2[0])

	if rows != cols {
		return nil, errors.New("invalid input, those matrixes can't be multiplied")
	}

	result := [][]int{}

	for i := 0; i < len(matrix1); i++ {
		result = append(result, []int{})
		for j := 0; j < len(matrix2[0]); j++ {
			for k := 0; k < len(matrix2); k++ {
				if len(result[i]) <= j {
					result[i] = append(result[i], matrix1[i][k]*matrix2[k][j])
				} else {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
				}
			}
		}

	}

	return result, nil
}
