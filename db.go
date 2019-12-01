 package main

import(
	"fmt"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func dbConn() (db *sql.DB, err error){

    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "riddhi"
    dbName := "stateStreet"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
    if err != nil {
        panic(err.Error())
    }
    return
}
