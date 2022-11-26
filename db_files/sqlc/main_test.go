package db

import (
	"database/sql"
	"log"
	"testing"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	dbSource = "postgresql://root:secret@localhost:5432/bank_db?sslmode=disable"
)

func TestMain(m *testing.M){
	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db")
	}

	testQueries = New(conn) 

	m.Run()
}