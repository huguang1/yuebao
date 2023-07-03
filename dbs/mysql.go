package dbs

import (
	"database/sql"
	"../libs"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Conns *sql.DB

func init() {
	var err error
	username := libs.Conf.Read("mysql", "username")
	password := libs.Conf.Read("mysql", "password")
	dataname := libs.Conf.Read("mysql", "dataname")
	port := libs.Conf.Read("mysql", "port")
	host := libs.Conf.Read("mysql", "host")
	dns := username + ":" + password + "@tcp(" + host + ":" + port + ")/" + dataname + "?parseTime=true"
	Conns, err = sql.Open("mysql", dns)
	if err != nil {
		log.Fatal(err.Error())
	}
	err = Conns.Ping()
	if err != nil {
		log.Fatal(err.Error())
	}
	Conns.SetMaxIdleConns(20)
	Conns.SetMaxOpenConns(20)
}
