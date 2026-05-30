/*
Copyright © 2026 https://github.com/BosBJJ/aether
*/
package main

import (
	"github.com/BosBJJ/aether/cmd"
	"github.com/BosBJJ/aether/cmd/config"
	"github.com/BosBJJ/aether/internal/database"
	"log"
	"os"
)

func main() {
	dbPath := os.Getenv("AETHER_DB_PATH")
	if dbPath == "" {
		dbPath = "aether.db"
	}

	db, err := database.MakeDB(dbPath)
	if err != nil {
		log.Fatalf("unable to create database: %v", err)
	}
	config.DB = db
	err = database.CreateSchema(db) 
	if err != nil {
		log.Fatalf("unable to create database: %v", err)
	}
	cmd.Execute()
}

