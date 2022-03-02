package main

import (
	"errors"
	"sync"
)

type TwoDimensionalMatrix [][]int

func MatrixMultiplication(matrix1, matrix2 TwoDimensionalMatrix) (TwoDimensionalMatrix, error) {
	rows := len(matrix1)
	columns := len(matrix2[0])

	if rows != columns {
		return nil, errors.New("invalid input, those matrixes can't be multiplied")
	}

	result := TwoDimensionalMatrix{}

	for i := 0; i < rows; i++ {
		result = append(result, make([]int, columns))
		for j := 0; j < columns; j++ {
			for k := 0; k < len(matrix2); k++ {
				result[i][j] += matrix1[i][k] * matrix2[k][j]
			}
		}

	}

	return result, nil
}

func MultiThreadedMatrixMultiplication(matrix1, matrix2 TwoDimensionalMatrix) (TwoDimensionalMatrix, error) {
	rows := len(matrix1)
	columns := len(matrix2[0])

	if rows != columns {
		return nil, errors.New("invalid input, those matrixes can't be multiplied")
	}

	var wg sync.WaitGroup
	wg.Add(rows)

	result := TwoDimensionalMatrix{}

	for i := 0; i < rows; i++ {
		result = append(result, make([]int, columns))
		go func(i int) {
			defer wg.Done()
			for j := 0; j < columns; j++ {
				for k := 0; k < len(matrix2); k++ {
					result[i][j] += matrix1[i][k] * matrix2[k][j]
				}
			}
		}(i)
	}

	wg.Wait()

	return result, nil
}
