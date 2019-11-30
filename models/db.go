//go run xlToSql.go db.go
//go get github.com/360EntSecGroup-Skylar/excelize
package models

import (
    "fmt"
    "encoding/csv"
    "io"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "os"
)

func readSample(rs io.ReadSeeker) ([][]string, error) {
    // Read remaining rows
    r := csv.NewReader(rs)
    rows, err := r.ReadAll()
    if err != nil {
        return nil, err
    }
    return rows, nil
}

// func csv_to_sql(db *sql.DB) {
func Csv_to_sql() (){
    //open csv file
    f, err := os.Open("SPY_All_Holdings.csv")
    if err != nil {panic(err)}
    defer f.Close()
    rows, err := readSample(f)
    if err != nil {panic(err)}

    //get a handle for a database
    fmt.Println("connecting to DB")
    Db, err := sql.Open("mysql", "root:riddhi@tcp(127.0.0.1:3306)/stateStreet")
    if err != nil{panic(err.Error())}

    Db.Exec("TRUNCATE TABLE ETFs")
    // perform a db.Query insert
    for i:=4; i<14; i++{
        var Name string = rows[i][0]
        var Ticker string = rows[i][1] 
        var Identifier string = rows[i][2]
        var SEDOL string = rows[i][3]
        var Weight string = rows[i][4]
        // Weight, err := strconv.ParseFloat(rows[i][4], 64)
        //     if err != nil{
        //     panic(err.Error())
        //     }
        var Sector string = rows[i][5]
        var Shares_Held string = rows[i][6]
        // Shares_Held, err := strconv.ParseInt(rows[i][6], 10, 64)
        //     if err != nil{
        //     panic(err.Error())
        //     }
        var Local_Currency string = rows[i][7]  
        Db.Exec("INSERT INTO ETFs(Name, Ticker, Identifier, SEDOL, Weight, Sector, Shares_Held, Local_Currency) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", Name, Ticker, Identifier, SEDOL, Weight, Sector, Shares_Held, Local_Currency)
           }//for close
        defer Db.Close()
}//func close

func ConnectDatabase() (string){
    var connectionString string = "root:riddhi@tcp(127.0.0.1:3306)/stateStreet"
    return connectionString
}






