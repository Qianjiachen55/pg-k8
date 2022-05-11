package db

import (
	"database/sql"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:123@localhost:5432/simple_bank?sslmode=disable"
)

func TestMain(m *testing.M){
	var err error
	testDB,err = sql.Open(dbDriver, dbSource)
	//defer conn.Close()
	if err != nil{
		log.Fatal("cannot connect to db: ",err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}