# Multi-Threaded Matrix Multiplication

Implementing a multi-threaded matrix multiplication and compare that to single-threaded algorithm.

- [x] Implement a simple matrix multiplication algorithm
- [x] Use goroutines to multi-thread the process 


## Results
Results of my benchmarks

1. When multiplying small matrixes, the overhead of goroutines is too big. Therefore, the the time per operation is 5 times higher (227 -> 1458 ns/op)
2. We can cut time per operation to a third when multiplying 50x50 matrixes (338420 -> 107934 ns/op)


```
$ go test -bench=.
goos: darwin
goarch: arm64
pkg: www.github.com/frederik-suerig/matrix-multiplication
BenchmarkSqaureMatrixMultiplication-10              	 4197498	       277.0 ns/op
BenchmarkMultiThreadedMatrixMultiplication-10       	  775122	      1468 ns/op
BenchmarkBigMatrixMultiplication-10                 	    3481	    338420 ns/op
BenchmarkBigMultiThreadedMatrixMultiplication-10    	    9504	    107934 ns/op
PASS
ok  	www.github.com/frederik-suerig/matrix-multiplication	5.006s
```