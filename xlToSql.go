//go run xlToSql.go db.go
//go get github.com/360EntSecGroup-Skylar/excelize
package main

import (
    "encoding/csv"
    "fmt"
    "io"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "os"
)

// type ETF_Model struct{
//     Db *sql.DB
// }
// var dbObj *sql.DB

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
func csv_to_sql() {
    //open csv file
    f, err := os.Open("SPY_All_Holdings.csv")
    if err != nil {
        panic(err)
    }
    defer f.Close()
    rows, err := readSample(f)
    if err != nil {
        panic(err)
    }

// rk = 0
// ak := "8294981000.0"
// rk, err =strconv.Atoi(ak)
// rk1:=int64(rk)
// fmt.Println(rk1)
// fmt.Println(reflect.TypeOf(ak))

    //get a handle for a database
    fmt.Println("connecting to DB")
    db, err := sql.Open("mysql", "root:riddhi@tcp(127.0.0.1:3306)/stateStreet")
    if err != nil{
        panic(err.Error())
    }

    // // Skip foth to 14th row 
    // for i:=4; i<14; i++{
    //     fmt.Println(rows[i])
    //     for j:=0; j<8; j++{
    
    //         fmt.Println(reflect.TypeOf(rows[i][j]))
    // }}


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
    db.Exec("INSERT INTO ETFs(Name, Ticker, Identifier, SEDOL, Weight, Sector, Shares_Held, Local_Currency) VALUES(?, ?, ?, ?, ?, ?, ?, ?)", Name, Ticker, Identifier, SEDOL, Weight, Sector, Shares_Held, Local_Currency)
    fmt.Println("insert complete")
    defer db.Close()
}

func demo(d int) int{
    d=d+1
    return d
}




