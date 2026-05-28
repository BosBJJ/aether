package database

import (
	"database/sql"
	"fmt"
	"time"

	_ "modernc.org/sqlite"
)



func MakeDB(path string) (*sql.DB, error) {
	db, err := sql.Open("sqlite", path)
	if err != nil {
    	return nil, fmt.Errorf("failed to open database at %s: %v", path, err)
	}
	err = db.Ping()
	if err != nil {
		return nil, fmt.Errorf("error creating DB: %v", err)
	}
	return db, nil
}

type ScanResult struct {
	Id 				int
	Timestamp 		string
	CommandType 	string
	Target 			string
	RawResult 		string
	Summary 		string
}


func CreateSchema(db *sql.DB) error {
	query := `CREATE TABLE IF NOT EXISTS scans (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
	timestamp TEXT,
	command_type TEXT,
	target TEXT,
	raw_result TEXT,
	summary TEXT);
	`

	_, err := db.Exec(query)
	return err
}

func SaveScan(db *sql.DB, scan ScanResult) error {
	scan.Timestamp = time.Now().UTC().Format(time.RFC3339)

	query := `INSERT INTO scans (timestamp, command_type, target, raw_result, summary)
	VALUES (?, ?, ?, ?, ?)`

	_, err := db.Exec(query, scan.Timestamp, scan.CommandType, scan.Target, scan.RawResult, scan.Summary)

	return err
}

func ListScans(db *sql.DB) ([]ScanResult, error) {
	query := `SELECT * FROM scans`

	rows, err := db.Query(query)
	if err != nil {
		return nil, fmt.Errorf("error :%v", err)
	}
	defer rows.Close()
	
	var results []ScanResult

	for rows.Next() {
		var result ScanResult
		err = rows.Scan(&result.Id, &result.Timestamp, &result.CommandType, &result.Target, &result.RawResult, &result.Summary)
		if err != nil {
			return nil, fmt.Errorf("error: %v", err)
		}
		results = append(results, result)
	}
	
	return results, nil
}

func GetScan(db *sql.DB, id int) (ScanResult, error) {
	query := `SELECT * FROM scans WHERE id = ?`

	row := db.QueryRow(query, id)
	
	var result ScanResult

	err := row.Scan(&result.Id, &result.Timestamp, &result.CommandType, &result.Target, &result.RawResult, &result.Summary)
	if err != nil {
		return ScanResult{}, fmt.Errorf("error fetching data: %v", err)
	}
	
	return result, nil
}

func DeleteScan(db *sql.DB, id int) error {
	query := `DELETE FROM scans WHERE id = ?`

	_, err := db.Exec(query, id)

	if err != nil {
		return fmt.Errorf("error deleting scan: %v", err)
	}
	

	return nil
}