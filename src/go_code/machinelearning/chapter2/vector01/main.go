package main
import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/floats"
)

func main() {
	//Initialize a "vector" via a slice
	var myvector []float64

	//Add a couple of components to the vector
	myvector = append(myvector, 11.0)
	myvector = append(myvector, 5.2)

	//output the results to stdout
	fmt.Println(myvector)

		//use gonum mat
		//Create a new vector value
//		myvector01 := mat.NewVector(2, []float64{11.0, 5.2})


	//Initialize a couple of "vectors" represented as slice
	vectorA := []float64{11.0, 5.2, -1.3}
	vectorB := []float64{-7.2, 4.2, 5.1}

	//Compute the dot product of A and B
	//(https://en.wikipedia.org/wiki/Dot_product)
	dotProduct := floats.Dot(vectorA, vectorB)
	fmt.Printf("The dot product of A and B is: %0.2f\n", dotProduct)

	//Scale each element of A by 1.5
	floats.Scale(1.5, vectorA)
	fmt.Printf("Scale A by 1.5 gives: %v\n", vectorA)

	//Compute the norm/length of B.
	normB := floats.Norm(vectorB, 2)
	fmt.Printf("The norm/length of B is: %0.2f\n", normB)

	//gonum mat
	vectorAm := mat.NewVector(3, []float64{11.0,5.2,-1.3}) 
}