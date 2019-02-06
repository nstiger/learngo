/*
行列式、特征集向量求解，矩阵转置
*/

package main
import(
	"fmt"
	"gonum.org/v1/gonum/mat"
	"log"
)
func main() {
	//Create a new matrix a.
	a := mat.NewDense(3, 3, []float64{1, 2, 3, 0, 4, 5, 0, 0, 6})

	//Compute and output the transpose of the matrix
	ft := mat.Formatted(a.T(), mat.Prefix("      "))
	fmt.Printf("a^T = %v\n\n", ft)

	//Compute and output the determinant of a.
	//一个n×n矩阵的行列式等于其任意行(或列)的元素与对应的代数余子式乘积之和
	deta := mat.Det(a)
	fmt.Printf("det(a) = %.2f\n\n", deta)

	//Compute and output the inverse of a.
	aInverse := mat.NewDense(3, 3, nil)
	if err := aInverse.Inverse(a); err != nil {
		log.Fatal(err)
	}
	fi := mat.Formatted(aInverse, mat.Prefix("      "))
	fmt.Printf("a^-1 = %v \n\n", fi)
}