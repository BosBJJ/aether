/*
Copyright © 2026 NAME HERE <EMAIL ADDRESS>
*/
package main

import (
	"aether/cmd"
	"aether/cmd/config"
	"aether/internal/database"
	"log"
)

func main() {
	db, err := database.MakeDB("aether.db")
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

