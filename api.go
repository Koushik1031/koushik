package main

import (
	"encoding/csv"
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
)

type Record struct {
	ID    int    `csv:"id"`
	Name  string `csv:"name"`
	Email string `csv:"email"`
}

var records []Record

func loadDataset() error {
	// Open the dataset file
	file, err := os.Open("dataset.csv")
	if err != nil {
		return err
	}
	defer file.Close()

	// Parse the CSV data into records
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.TrimLeadingSpace = true
	lines, err := reader.ReadAll()
	if err != nil {
		return err
	}
	records = make([]Record, len(lines))
	for i, line := range lines {
		id, err := strconv.Atoi(line[0])
		if err != nil {
			return err
		}
		records[i] = Record{
			ID:    id,
			Name:  line[1],
			Email: line[2],
		}
	}

	return nil
}

func getRecordsHandler(w http.ResponseWriter, r *http.Request) {
	// Write the records as CSV to the response
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=dataset.csv")
	writer := csv.NewWriter(w)
	writer.Write([]string{"id", "name", "email"})
	for _, record := range records {
		writer.Write([]string{strconv.Itoa(record.ID), record.Name, record.Email})
	}
	writer.Flush()
}

func getRecordsWithFilterHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the query params
	q := r.URL.Query()
	filter := q.Get("filter")
	value := q.Get("value")

	// Write the matching records as CSV to the response
	w.Header().Set("Content-Type", "text/csv")
	w.Header().Set("Content-Disposition", "attachment; filename=dataset.csv")
	writer := csv.NewWriter(w)
	writer.Write([]string{"id", "name", "email"})
	for _, record := range records {
		fieldValue := ""
		switch filter {
		case "id":
			fieldValue = strconv.Itoa(record.ID)
		case "name":
			fieldValue = record.Name
		case "email":
			fieldValue = record.Email
		default:
			http.Error(w, fmt.Sprintf("Invalid filter: %s", filter), http.StatusBadRequest)
			return
		}
		if strings.Contains(strings.ToLower(fieldValue), strings.ToLower(value)) {
			writer.Write([]string{strconv.Itoa(record.ID), record.Name, record.Email})
		}
	}
	writer.Flush()
}

func createRecordHandler(w http.ResponseWriter, r *http.Request) {
	// Parse the request body
	reader := csv.NewReader(r.Body)
	reader.FieldsPerRecord = 3
	reader.TrimLeadingSpace = true
	lines, err := reader.ReadAll()
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	if len(lines) != 1 {
		http.Error(w, "Invalid number of records in request", http.StatusBadRequest)
		return
	}
	line := lines[0]

	// Create a new record
	id, err := strconv.Atoi(line[0])
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	record := Record{
		ID:    id,
		Name:  line[1],
		Email:


   
