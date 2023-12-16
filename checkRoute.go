package main 

import (
	"fmt"
	"encoding/csv"
	"strings"
	"log"
	"io"
	"bufio"
	"os"
)

func CheckRoute() {

	logger := log.New(os.Stdout, "INFO: ", log.Ldate|log.Ltime)

	logfile, err := os.Create("app.log")

	if err != nil {
		log.Fatal(err)
	}
	defer logfile.Close()
	logger.SetOutput(logfile)

	
	file, err := os.Open(CsvFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	reader := csv.NewReader(file)

	header, err := reader.Read()
	if err != nil {
		log.Fatal(err)
	}

	colIndex := make(map[string]int)
	for i, colName := range header {
		colIndex[colName] = i
	}
	logger.Println("Scanning Data")
	fmt.Print("Enter the Origin: ")

	var ori string
	var dest string
	fmt.Scanln(&ori)
	var criteria string = "Origin"

	var criteria2 string = "Dest"
	var criteria3 string
	var correctNum int
	fmt.Print("Enter the Destination: ")

	fmt.Scanln(&dest)
	fmt.Printf("Enter the Route: ")
	scanner := bufio.NewScanner(os.Stdin)
	if scanner.Scan() {
		criteria3 = scanner.Text()
	}

	var matchingRows []RouteData

	for {
		row, err := reader.Read()
		if err != nil {
			if err == io.EOF {
				break
			} else {
				log.Fatal(err)
			}
		}

		if row[colIndex[criteria]] == ori && row[colIndex[criteria2]] == dest {
			routeData := RouteData{
				Origin:   row[colIndex["Origin"]],
				Route:    row[colIndex["Route String"]],
				Dest:     row[colIndex["Dest"]],
				Altitude: row[colIndex["Altitude"]],
				Aircraft: row[colIndex["Aircraft"]],
				DCNTR:    row[colIndex["DCNTR"]],
				ACNTR:    row[colIndex["ACNTR"]],
			}

			if routeData.Altitude == "" {
				routeData.Aircraft = "None required!"
			}
			if routeData.Altitude == "" {
				routeData.Altitude = "None Required!"
			}

			matchingRows = append(matchingRows, routeData)

			if strings.TrimSpace(routeData.Route) == strings.TrimSpace(criteria3) {
				correctNum = correctNum + 1000
			} else {
				correctNum = correctNum - 1
			}
		}
	}
	

	if correctNum <= 0 {
		fmt.Printf("You entered this '%s', Here is the correct route(s): \n", criteria3)
		for _, route := range matchingRows {
			if route.Aircraft == "" {
				route.Aircraft = "N/A"
			}
			if route.Altitude == "" {
				route.Altitude = "N/A"
			}
			fmt.Printf("Route: %s, Aircraft Type: %s, Altitude: %s\n", route.Route, route.Aircraft, route.Altitude)
		}
	} else if correctNum > 0 {

		fmt.Println(`You have the correct route!`)
	} else {
		fmt.Println("No matching rows found.")
	}
}

func test() {
	hello := "hi"
	bye := "hi"

	if hello == bye {
		fmt.Println("hello")
	}
}
