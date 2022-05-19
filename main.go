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
	log.Println("config has been load")

	conn, err := sql.Open(config.DBDriver, config.DBSource)
	//defer conn.Close()
	if err != nil {
		log.Fatal("cannot connect to db: ", err)
	}
	log.Println("db conn init success!")

	store := db.NewStore(conn)
	server,err := api.NewServer(config, store)
	if err !=nil{
		log.Fatal("cannot create server: ",err)
	}

	err = server.Start(config.ServerAddress)
	if err != nil{
		log.Fatal("cannot start server:",err)
	}
}
