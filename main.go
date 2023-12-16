package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
)

// Define a struct to represent the data in each row

type RouteData struct {
	Origin   string
	Route    string
	Altitude string
	Aircraft string
	Dest     string
	DCNTR    string
	ACNTR    string
}

var base string


var CsvFileName string

func main() {
	
	executablePath, err := os.Executable()
	if err != nil {
		fmt.Println("Error getting executable path:", err)
		return
	}

	currentDir := filepath.Dir(executablePath)

	relativePath := "faadata.csv"

	filePath := filepath.Join(currentDir, relativePath)

	fmt.Println(filePath)
	CsvFileName = filePath

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	logfile, err := os.Create("app.log")

	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	logger.SetOutput(logfile)

	logger.Println("Scanning Choice")
	var choice int
	fmt.Println("1 or 2 or 3 or 4")
	fmt.Scanln(&choice)

	if choice == 1 {
		SearchRow()
	}
	if choice == 2 {
		FindRoute()
	}
	if choice == 3 {
		FindRoute2()
	}
	if choice == 4 {
		CheckRoute()
	} else {
		fmt.Println("Invalid Choice")
	}

}
