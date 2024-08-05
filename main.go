package main

import (
	"eth_usd/csv_process"
	"fmt"
	"os"
)

func main() {
	_, err := csv_process.LoadDataFrom("prices.csv")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading prices from CSV file: %v\n", err)
		os.Exit(1)
	}

}
