// mysql.server start
// mysql -u root -p
//riddhi
// SHOW DATABASES;
// CREATE DATABASE stateStreet;
//USE stateStreet;
// CREATE TABLE ETF(Name VARCHAR(255) NOT NULL, 
// Ticker VARCHAR(10) NOT NULL,
// Identifier VARCHAR(30) NOT NULL, 
// SEDOL VARCHAR(30) NOT NULL, 
// Weight float NOT NULL, 
// Sector VARCHAR(100) NOT NULL, 
// Shares_Held float NOT NULL, 
// Local_Currency VARCHAR(10) NOT NULL, PRIMARY KEY(Ticker));

//or 

// CREATE TABLE ETFs(Name VARCHAR(255) NOT NULL, 
// Ticker VARCHAR(10) NOT NULL,
// Identifier VARCHAR(30) NOT NULL, 
// SEDOL VARCHAR(30) NOT NULL, 
// Weight VARCHAR(30) NOT NULL, 
// Sector VARCHAR(100) NOT NULL, 
// Shares_Held VARCHAR(30) NOT NULL, 
// Local_Currency VARCHAR(10) NOT NULL, PRIMARY KEY(Ticker));

//SHOW FIELDS FROM ETF;

// export DJANGO_PROFILE=local
// export MARIADB_USERNAME=root
// export MARIADB_PASSWORD=riddhi
// export MARIADB_DATABASE=stateStreet
// export BACKEND_HOST=localhost
// export MARIADB_PORT=3306
// export BACKEND_PORT=8001
// export MARIADB_HOST=localhost
// echo "local profile set"  


package main

import(
	"fmt"
	// "encoding/json"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	// "log"
	// "net/http"
	// "github.com/gorilla/mux"
	// "math/random"
	// "strconv" //converts to string
)
// func dbConn(){
func dbConn() (db *sql.DB, err error){

    dbDriver := "mysql"
    dbUser := "root"
    dbPass := "riddhi"
    dbName := "stateStreet"
    db, err := sql.Open(dbDriver, dbUser+":"+dbPass+"@/"+dbName)
// db, err := sql.Open("mysql", "username:password@tcp(127.0.0.1:3306)/test")
    if err != nil {
        panic(err.Error())
    }
    return
	// return db
}

func demo() {
    fmt.Println("HI")
}