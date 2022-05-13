package main

import (
	"database/sql"
	"github.com/Qianjiachen55/pgK8/api"
	db "github.com/Qianjiachen55/pgK8/db/sqlc"
	"github.com/Qianjiachen55/pgK8/util"
	_ "github.com/lib/pq"
	"log"
)



func main() {
	config,err := util.LoadConfig(".")
	if err != nil{
		log.Fatal("cannot load config:",err)
	}

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	//defer conn.Close()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}

	store := db.NewStore(conn)
	server := api.NewServer(store)

	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("cannot start server:",err)
	}
}
