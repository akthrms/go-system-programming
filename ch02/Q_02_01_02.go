package main

import (
	"encoding/csv"
	"os"
)

func main() {
	file, err := os.Create("text.csv")
	if err != nil {
		panic(err)
	}

	writer := csv.NewWriter(file)
	writer.Write([]string{"encoding", "csv"})
	writer.Write([]string{"hello", "world"})
	writer.Flush()
}
