package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
)

func FindRoute2() {
	logfile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
    log.Fatal(err)
}
logger := log.New(io.MultiWriter(os.Stdout, logfile), "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Opening CSV")
	logger.Println("Opening CSV...")
	file, err := os.Open(CsvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger.Println("Opened")
	reader := csv.NewReader(file)
	logger.Println("Starting User Input")
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	colIndex := make(map[string]int)
	for i, colName := range header {
		colIndex[colName] = i
	}

	fmt.Print("Enter the Origin: ")

	var ori string
	var dest string
	fmt.Scanln(&ori)
	var criteria string = "Origin"

	var criteria2 string = "Dest"

	fmt.Print("Enter the Destination: ")
	fmt.Scanln(&dest)

	logger.Printf("User input complete. Origin: '%s' Destination: '%s'", ori, dest)
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
		if row[colIndex[criteria]] == ori && row[colIndex[criteria2]] == dest {
			// Create a RouteData struct and populate it with the values from the row
			routeData := RouteData{
				Origin:   row[colIndex["Origin"]],
				Route:    row[colIndex["Route String"]],
				Dest:     row[colIndex["Dest"]],
				Aircraft: row[colIndex["Aircraft"]],
				Altitude: row[colIndex["Altitude"]],
				DCNTR:    row[colIndex["DCNTR"]],
				ACNTR:    row[colIndex["ACNTR"]],
			}

			// Append the RouteData to the slice
			matchingRows = append(matchingRows, routeData)
		}
	}
	logger.Println("Reading Complete")
	// Print the data for matching rows
	for _, routeData := range matchingRows {
		fmt.Printf("RouteData: %+v\n", routeData)
	}

	// Print a message if no matching rows were found
	if len(matchingRows) == 0 {
		fmt.Println("No matching rows found.")
	}
}
