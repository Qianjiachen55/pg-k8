package db

import (
	"database/sql"
	"github.com/Qianjiachen55/pgK8/util"
	_ "github.com/lib/pq"
	"log"
	"os"
	"testing"
)

var testQueries *Queries
var testDB *sql.DB


func TestMain(m *testing.M){
	config,err := util.LoadConfig("../..")
	if err != nil{
		log.Fatal("cannot load config:",err)
	}
	testDB,err = sql.Open(config.DBDriver, config.DBSource)
	//defer conn.Close()
	if err != nil{
		log.Fatal("cannot connect to db: ",err)
	}

	testQueries = New(testDB)

	os.Exit(m.Run())
}