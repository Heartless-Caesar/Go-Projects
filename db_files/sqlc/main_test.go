package db

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"

	_ "github.com/lib/pq"
)

var testQueries *Queries

const (
	dbDriver = "postgres"
	host = "localhost"
	port = 5432
	user = "bank_db"
	password = "887887yua2"
	dbname = "bank_db" 
)


func TestMain(m *testing.M){
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+
		"password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)
	conn, err := sql.Open(dbDriver, psqlInfo)
	if err != nil {
		log.Fatal("Cannot connect to db")
	}

	testQueries = New(conn) 

	os.Exit(m.Run()) 
}
