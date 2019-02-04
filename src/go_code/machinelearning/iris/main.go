package main
import (
	"fmt"
	"encoding/csv"
	"os"
	"log"
)
func main() {
	//open the iris dataset file 
	f, err := os.Open("iris.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	//Create a new CSV reader reading from the opened file
	reader := csv.NewReader(f)

	//Assume we don't know the number of fields per line. By setting
	//FieldsPerRecord negative, each row may have a variable
	//number of fields.
	reader.FieldsPerRecord = -1
	//Read in all of the CSV records.
	rawCSVData, err := reader.ReadAll()
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; i < len(rawCSVData); i++ {
		for j := 0; j < len(rawCSVData[i]); j++ {
			fmt.Print(rawCSVData[i][j], "\t")

		}
		fmt.Println()
	}


}