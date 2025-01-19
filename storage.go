package main

import "database/sql"

type DatabaseStorage struct {
	db *sql.DB
}

func NewDatabaseStorage(db *sql.DB) *DatabaseStorage {
	return &DatabaseStorage{db: db}
}

func (s *DatabaseStorage) Save(data Saveable) error {
	// Encode data and save it to the database
	return nil
}

func (s *DatabaseStorage) Load(data Saveable) error {
	// Load data from the database and decode it
	return nil
}
