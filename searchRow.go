package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
)

func SearchRow() {
	logfile, err := os.OpenFile("app.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
if err != nil {
    log.Fatal(err)
}
logger := log.New(io.MultiWriter(os.Stdout, logfile), "INFO: ", log.Ldate|log.Ltime)
	logger.Println("Opening CSV")
	// Open the CSV file
	file, err := os.Open(CsvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	logger.Println("Opened")
	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read the header row to determine column names
	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Map to store the index of each column based on its name
	colIndex := make(map[string]int)
	for i, colName := range header {
		colIndex[colName] = i
	}
	logger.Println("Gathering User Input")
	// Prompt the user for a row number
	fmt.Print("Enter the row number: ")
	var rowNumberInput string
	fmt.Scanln(&rowNumberInput)
	logger.Printf("Input gathered. Row: %s", rowNumberInput)
	// Convert the user input to an integer
	rowNumber, err := strconv.Atoi(rowNumberInput)
	rowNumber = rowNumber - 2
	if err != nil {
		log.Fatal("Invalid row number. Please enter a valid integer.")
	}

	// Read the CSV until the desired row
	var currentRow int
	for currentRow < rowNumber {
		_, err := reader.Read()
		if err != nil {
			log.Fatal("Row number exceeds the number of rows in the CSV.")
		}
		currentRow++
	}

	// Read the data from the desired row
	row, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	// Create a RouteData struct and populate it with the values from the row
	routeData := RouteData{
		Origin:   row[colIndex["Orig"]],
		Route:    row[colIndex["Route String"]],
		Dest:     row[colIndex["Dest"]],
		Altitude: row[colIndex["Altitude"]],
		DCNTR:    row[colIndex["DCNTR"]],
		ACNTR:    row[colIndex["ACNTR"]],
	}

	// Print the data in struct format
	fmt.Printf("RouteData: %+v\n", routeData)
}
