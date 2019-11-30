package models

// import (
//     "bufio"
//     "encoding/csv"
//     "encoding/json"
//     "fmt"
//     "io"
//     "log"
//     "os"
// )

// type ETF struct {
//     Name string 
//     Ticker string
//     Identifier string 
//     SEDOL float32
//     Weight float32
//     Sector string
//     Shares_Held float32
//     Local_Currency string
// }

// func main(){
//     csvFile, _ := os.Open("people.csv")
//     reader := csv.NewReader(bufio.NewReader(csvFile))
//     var etf []ETF
//     for {
//         line, error := reader.Read()
//         if error == io.EOF {
//             break
//         } else if error != nil {
//             log.Fatal(error)
//         }
//         etf = append(etf, ETF{
//             Name: line[0],
//             Ticker:  line[1],
//             Identifier: line[2] 
//             SEDOL: line[3]
//             Weight: line[4]
//             Sector: line[5]
//             Shares_Held: line[6]
//             Local_Currency: line[7]
//             },
//         })
//     }
//     etfJson, _ := json.Marshal(etf)
//     fmt.Println(string(etfJson))
// }


// // func AllETFs() ([]*ETF, error) {

//     // rows, err := db.Query("SELECT * FROM books")
//     // if err != nil {
//     //     return nil, err
//     // }
//     // defer rows.Close()

//     // bks := make([]*Book, 0)
//     // for rows.Next() {
//     //     bk := new(Book)
//     //     err := rows.Scan(&bk.Isbn, &bk.Title, &bk.Author, &bk.Price)
//     //     if err != nil {
//     //         return nil, err
//     //     }
//     //     bks = append(bks, bk)
//     // }
//     // if err = rows.Err(); err != nil {
//     //     return nil, err
//     // }
//     // return bks, nil
// // }


// // func (r *Reader) Read() (record []string, err error)
// // Read reads one record (a slice of fields) from r. If the record has an unexpected number of fields, Read returns the record along with 
// // the error ErrFieldCount. Except for that case, Read always returns either a non-nil record or a non-nil error, but not both. If there is
// //  no data left to be read, Read returns nil, io.EOF. If ReuseRecord is true, the returned slice may be shared between multiple calls to 
// //  Read.

















