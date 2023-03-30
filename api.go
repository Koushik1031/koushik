package main

import (
    "encoding/csv"
    "encoding/json"
    "fmt"
    "log"
    "net/http"
    "strconv"
    "strings"
)

type Record struct {
    ID      int    `json:"id"`
    Name    string `json:"name"`
    Age     int    `json:"age"`
    Country string `json:"country"`
}

var records []Record

func main() {
    // load data from CSV file
    records = loadCSV("dataset.csv")

    // setup HTTP routes
    http.HandleFunc("/records", getRecordsHandler)
    http.HandleFunc("/records/filter", filterRecordsHandler)
    http.HandleFunc("/records/create", createRecordHandler)
    http.HandleFunc("/records/update", updateRecordHandler)
    http.HandleFunc("/records/delete", deleteRecordHandler)

    // start server
    log.Fatal(http.ListenAndServe(":8080", nil))
}

func loadCSV(filename string) []Record {
    file, err := os.Open(filename)
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    reader := csv.NewReader(file)
    reader.Comma = ','

    lines, err := reader.ReadAll()
    if err != nil {
        log.Fatal(err)
    }

    var records []Record
    for i, line := range lines {
        if i == 0 {
            continue // skip header row
        }

        id, _ := strconv.Atoi(line[0])
        age, _ := strconv.Atoi(line[2])

        records = append(records, Record{
            ID:      id,
            Name:    line[1],
            Age:     age,
            Country: line[3],
        })
    }

    return records
}

func getRecordsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(records)
}

func filterRecordsHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    q := r.URL.Query()
    var filtered []Record
    for _, record := range records {
        if q.Get("id") != "" && strconv.Itoa(record.ID) != q.Get("id") {
            continue
        }
        if q.Get("name") != "" && !strings.Contains(strings.ToLower(record.Name), strings.ToLower(q.Get("name"))) {
            continue
        }
        if q.Get("age") != "" {
            age, err := strconv.Atoi(q.Get("age"))
            if err != nil || record.Age != age {
                continue
            }
        }
        if q.Get("country") != "" && !strings.Contains(strings.ToLower(record.Country), strings.ToLower(q.Get("country"))) {
            continue
        }

        filtered = append(filtered, record)
    }

    json.NewEncoder(w).Encode(filtered)
}

func createRecordHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    var record Record
    err := json.NewDecoder(r.Body).Decode(&record)
    if err != nil {
        http.Error(w, err.Error(), http.StatusBadRequest)
        return
    }

    // generate new ID
    maxID := 0
    for _, r := range records {
        if r.ID > maxID {
            maxID = r.ID
        }
    }
    record.ID = maxID + 1

    records = append(records, record)

    json.NewEncoder(w).Encode(record)
}

func updateRecordHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("
func deleteRecordHandler(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")

    id := r.URL.Query().Get("id")
    if id == "" {
        http.Error(w, "ID parameter is required", http.StatusBadRequest)
        return
    }

    var deleted Record
    for i, record := range records {
        if strconv.Itoa(record.ID) == id {
            deleted = record
            // remove the record from the slice
            records = append(records[:i], records[i+1:]...)
            break
        }
    }

    if (Record{}) == deleted {
        http.Error(w, fmt.Sprintf("Record with ID %s not found", id), http.StatusNotFound)
        return
    }

    json.NewEncoder(w).Encode(deleted)
}


   
