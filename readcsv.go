package services

import (
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
)

func ReadCSV(filePath string) {
	Db := setupDatabase() // Open database connection once

	// Open the CSV file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("Error: failed to open the file %v", err)
	}
	defer file.Close()

	// Read all records from the CSV file
	reader := csv.NewReader(file)
	records, err := reader.ReadAll()
	if err != nil {
		log.Fatalf("Failed to read CSV file: %v", err)
	}

	// Parse CSV records and insert into the database
	for _, record := range records[1:] { // Skip the header row
		// Error handling for each conversion
		siteID, err := strconv.Atoi(record[0])
		if err != nil {
			log.Printf("Error converting SiteID: %v", err)
			continue // Skip this record if conversion fails
		}

		fixletID, err := strconv.Atoi(record[1])
		if err != nil {
			log.Printf("Error converting FixletID: %v", err)
			continue // Skip this record if conversion fails
		}

		relevantComputerCount, err := strconv.Atoi(record[4])
		if err != nil {
			log.Printf("Error converting RelevantComputerCount: %v", err)
			continue // Skip this record if conversion fails
		}

		// Create a new Record struct
		user := Record{
			SiteID:                uint(siteID),
			FixletID:              int64(fixletID), // Ensure FixletID can handle large values
			Name:                  record[2],
			Criticality:           record[3],
			RelevantComputerCount: relevantComputerCount,
		}

		// Insert record into the database
		if err := Db.Create(&user).Error; err != nil {
			log.Printf("Error inserting record into DB: %v", err)
			continue // Skip this record if DB insertion fails
		}
	}

	fmt.Println("CSV data has been successfully inserted into the database!")
}
