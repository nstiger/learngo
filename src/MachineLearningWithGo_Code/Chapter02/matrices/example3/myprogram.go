package main

import (
	"fmt"
	"log"

	"github.com/gonum/matrix/mat64"
)

func main() {

	// Create a new matrix a.
	a := mat64.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	// Compute and output the transpose of the matrix.
	//a的转置矩阵
	ft := mat64.Formatted(a.T(), mat64.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	// Compute and output the determinant of a.
	//a的行列式
	deta := mat64.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	// Compute and output the inverse of a.
	//a的逆矩阵
	aInverse := mat64.NewDense(0, 0, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat64.Formatted(aInverse, mat64.Prefix("       "))
	fmt.Printf("a^-1 = %v\n\n", fi)
}
