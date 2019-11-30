package main
// Go MySQL Driver is an implementation of Go's database/sql/driver interface. 
// You only need to import the driver and can use the full database/sql API then.
import(
    "fmt"
    // "Users/riddhikakadiya/go/src/github.com/riddhikakadiya/spdrs/models/db"
    "encoding/json"
    "github.com/riddhikakadiya/spdrs/models"
    // "encoding/csv"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "log"
)
// type ETF struct {
//     Name string `json:”Name”`
//     Ticker string `json:”Ticker”`
//     Identifier string `json:”Identifier”`
//     SEDOL float32 `json:”SEDOL”`
//     Weight float32 `json:”Weight”`
//     Sector string `json:”Sector”`
//     Shares_Held float32 `json:”Shares_Held”`
//     Local_Currency string `json:”Local_Currency”`
// }
type ETF struct {
    Name string 
    Ticker string 
    Identifier string 
    SEDOL string 
    Weight string 
    Sector string 
    Shares_Held string 
    Local_Currency string 
}

type Ticker struct {
	Ticker string
}

func getJson(in []byte) []byte {
	var raw map[string]interface{}
	json.Unmarshal(in, &raw)
	out, _ := json.Marshal(raw)
	//println(string(out))
	return out
}

func availableETFs(w http.ResponseWriter, r *http.Request){  
	// var connectionString string = models.ConnectDatabase()
    Db, err := sql.Open("mysql", "root:riddhi@tcp(127.0.0.1:3306)/stateStreet")
    if err != nil{panic(err.Error())}

	fmt.Println("before select query")
	rows, err := Db.Query("SELECT * FROM ETFs")
    if err != nil{panic(err.Error())}
    defer rows.Close()


    etfs := make([]*ETF, 0)
    for rows.Next() {
        etf := new(ETF)
        err := rows.Scan(&etf.Name, &etf.Ticker, &etf.Identifier, &etf.SEDOL, &etf.Weight, &etf.Sector, &etf.Shares_Held, &etf.Local_Currency)
        if err != nil{panic(err.Error())}
        etfs = append(etfs, etf)
    }
    // for _, etf := range etfs {
    //     fmt.Println(etf.Name, etf.Ticker)
    // }

    jsonData, err := json.Marshal(etfs)
    if err != nil{panic(err.Error())}
    w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}

func signIn(w http.ResponseWriter, r *http.Request) {
	// w.Write([]byte("OK"))
	w.Header().Set("Content-Type", "application/json")
	w.Write(getJson([]byte(`{ "message":  "Hello from Weather App"  }`)))
}


func availableTickers(w http.ResponseWriter, r *http.Request){ 
    Db, err := sql.Open("mysql", "root:riddhi@tcp(127.0.0.1:3306)/stateStreet")
    if err != nil{panic(err.Error())}

	fmt.Println("before select query")
	rows, err := Db.Query("SELECT Ticker FROM ETFs")
    if err != nil{panic(err.Error())}
    defer rows.Close()

    tickers := make([]*Ticker, 0)
    fmt.Println("aftre make ticker")
    for rows.Next() {
        ticker := new(Ticker)
        err := rows.Scan(&ticker.Ticker)
        if err != nil{panic(err.Error())}
        tickers = append(tickers, ticker)
    }

    jsonData, err := json.Marshal(tickers)
    if err != nil{panic(err.Error())}
    w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}


func main(){ 
	fmt.Println("inside main")
   	models.DownloadData()
   	models.Csv_to_sql()
    router := mux.NewRouter()
    router.HandleFunc("/spdr", signIn).Methods("GET")
    router.HandleFunc("/spdr/etfs", availableETFs).Methods("GET")
    router.HandleFunc("/spdr/etfs/tickers", availableTickers).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
    if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
    fmt.Println("main complete")
}






