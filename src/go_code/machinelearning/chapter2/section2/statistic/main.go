package main
import (
	"github.com/montanaflynn/stats"
	"gonum.org/v1/gonum/stat"
	"github.com/kniren/gota/dataframe"
	"fmt"
	"os"
	"log"
)
//集中趋势测量法
func main() {
	//Open the CSV file
	irisFile, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer irisFile.Close()

	//Create a dataframe from the CSV file.
	irisDF := dataframe.ReadCSV(irisFile)

	//Get the float values from the "sepal_length" column as
	//we will be looking at the measures for this variable.
	sepalLength := irisDF.Col("sepal_length").Float()
	/*	petalLength := irisDF.Col("petal_length").Float()
		Output:
	petal Length Summary Statistics:
	Mean value: 3.76
	Mode value: 1.50
	Mode Count: 14
	Median value: 4.35
	*/

	//Calculate the Mean of the variable.均值
	meanVal := stat.Mean(sepalLength, nil)

	//Calculate the Mode of the variable.众数和个数
	modeVal, modeCount := stat.Mode(sepalLength, nil)

	//Calculate the Median of the variable.中位数
	medianVal, err := stats.Median(sepalLength)
	if err != nil {
		log.Fatal(err)
	}

	//Output the results to standard out.
	fmt.Printf("\nSepal Length Summary Statistics:\n")
	fmt.Printf("Mean value: %.2f\n", meanVal)
	fmt.Printf("Mode value: %.2f\n", modeVal)
	fmt.Printf("Mode Count: %v\n", modeCount)
	fmt.Printf("Median value: %.2f\n", medianVal)
}