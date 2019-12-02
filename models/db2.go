package models

// import (
//     "fmt"
//     "encoding/csv"

//     "net/http"
// )

// func readSample(url string) ([][]string, error) {
// 	resp, err := http.Get(url)
// 	if err != nil {return nil, err}

// 	defer resp.Body.Close()
// 	reader := csv.NewReader(resp.Body)
// 	reader.LazyQuotes = true
// 	rows, err := reader.ReadAll()
// 	if err != nil {
// 		return nil, err
// 	}
// 	return rows, nil
// }

// // func csv_to_sql(db *sql.DB) {
// func Csv_to_sql() (){
//     url := "https://us.spdrs.com/site-content/xls/SPY_All_Holdings.xls"
// 	rows, err := readSample(url)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// for idx, row := range data {
// 	// 	// skip header
// 	// 	if idx == 0 {
// 	// 		continue
// 	// 	}

// 	// 	if idx == 6 {
// 	// 		break
// 	// 	}

// 	// 	fmt.Println(row[2])
// 	// }


//     for i:=4; i<14; i++{
//         var Name string = rows[i][0]
//         var Ticker string = rows[i][1] 
//         var Identifier string = rows[i][2]
//         var SEDOL string = rows[i][3]
//         var Weight string = rows[i][4]
//         // Weight, err := strconv.ParseFloat(rows[i][4], 64)
//         //     if err != nil{
//         //     panic(err.Error())
//         //     }
//         var Sector string = rows[i][5]
//         var Shares_Held string = rows[i][6]
//         // Shares_Held, err := strconv.ParseInt(rows[i][6], 10, 64)
//         //     if err != nil{
//         //     panic(err.Error())
//         //     }
//         var Local_Currency string = rows[i][7]  
//         fmt.Println(Name,Ticker,Identifier,SEDOL,Weight,Sector,Shares_Held,Local_Currency)

// }//func close

// }

//problem! dotdefender
//https://stackoverflow.com/questions/31326659/golang-csv-error-bare-in-non-quoted-field




