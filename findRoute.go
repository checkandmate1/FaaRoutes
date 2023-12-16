package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"io"
)	
	

func FindRoute() {
	logfile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
    log.Fatal(err)
}
logger := log.New(io.MultiWriter(os.Stdout, logfile), "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Opening CSV")
	file, err := os.Open(CsvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger.Println("CSV Opened")
	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	colIndex := make(map[string]int)
	for i, colName := range header {
		colIndex[colName] = i
	}

	logger.Println("Gathering User Input...")
	fmt.Print("Enter the criteria to search for (Origin, Route, Dest, Altitude, DCNTR, ACNTR): ")
	var criteria string
	fmt.Scanln(&criteria)

	fmt.Printf("Enter the value for %s: ", criteria)
	var searchValue string
	fmt.Scanln(&searchValue)
	logger.Printf("Input complete. Criteria: '%s' and value is '%s'", criteria, searchValue)
	var matchingRows []RouteData

	// Read all rows from the CSV file
	logger.Println("Reading CSV")
	for {
		// Read the data from the row
		row, err := reader.Read()
		if err != nil {
			break // Break if there are no more rows
		}

		// Check if the specified criteria column contains the search value
		if row[colIndex[criteria]] == searchValue {
			// Create a RouteData struct and populate it with the values from the row
			routeData := RouteData{
				Origin:   row[colIndex["Origin"]],
				Route:    row[colIndex["Route String"]],
				Dest:     row[colIndex["Dest"]],
				Altitude: row[colIndex["Altitude"]],
				DCNTR:    row[colIndex["DCNTR"]],
				ACNTR:    row[colIndex["ACNTR"]],
			}

			// Append the RouteData to the slice
			matchingRows = append(matchingRows, routeData)
		}
	}
	logger.Println("CSV Read")
	// Print the data for matching rows
	for _, routeData := range matchingRows {
		fmt.Printf("RouteData: %+v\n", routeData)
	}

	// Print a message if no matching rows were found
	if len(matchingRows) == 0 {
		fmt.Println("No matching rows found.")
	}
}
