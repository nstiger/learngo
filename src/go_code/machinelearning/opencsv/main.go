package main
import (
	"fmt"
	"os"
	"log"
	"strconv"
	"encoding/csv"
)

func main() {
	//open the CSV
	f, err := os.Open("myfile.csv")
	if err != nil {
		log.Fatal(err)
	}

	//read in the CSV records
	r := csv.NewReader(f)
	records, err := r.ReadAll()
	if err != nil {
		log.Fatal(err)
	}

	//get the maximum value in the integer colnum
	var intMax int
	for _, record := range records {
		//Parse the integer value
		intVal, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatal(err)
		}

		//Replace the maximum value if appropriate
		if intVal > intMax {
			intMax = intVal
		}
	}

	//Print the maximum value
	fmt.Println(intMax)
}