package main

import (
	"fmt"

	"github.com/gonum/stat"
)

//2.3.4.1 测试统计
func main() {

	// Define observed and expected values. Most
	// of the time these will come from your
	// data (website visits, etc.).
	observed := []float64{48, 52}
	expected := []float64{50, 50}

	// Calculate the ChiSquare test statistic.
	chiSquare := stat.ChiSquare(observed, expected)

	fmt.Println(chiSquare)
}
