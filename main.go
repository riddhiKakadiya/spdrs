package main

import(
    "fmt"
    // "Users/riddhikakadiya/go/src/github.com/riddhikakadiya/spdrs/models/db"
    "encoding/json"
    "github.com/riddhikakadiya/spdrs/models"
    "time"
    "net/http"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
    "github.com/gorilla/mux"
    "log"
    "github.com/dgrijalva/jwt-go"
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

/***********************************************************************************************************************************/
//user authentication using JWT
var jwtKey = []byte("my_secret_key")
var users = map[string]string{
	"user1": "password1",
	"user2": "password2",
}
// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}
// Create a struct that will be encoded to a JWT.
// We add jwt.StandardClaims as an embedded type, to provide fields like expiry time
type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

func Signin(w http.ResponseWriter, r *http.Request) {
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(r.Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	// Get the expected password from our in memory map
	expectedPassword, ok := users[creds.Username]

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		w.WriteHeader(http.StatusUnauthorized)
		return
	}
	// Declare the expiration time of the token, here, we have kept it as 5 minutes
	expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		Username: creds.Username,
		StandardClaims: jwt.StandardClaims{
			// In JWT, the expiry time is expressed as unix milliseconds
			ExpiresAt: expirationTime.Unix(),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	// Finally, we set the client cookie for "token" as the JWT we just generated,we also set an expiry time which is the same as the token itself
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   tokenString,
		Expires: expirationTime,
	})
}
// func authenticateUser(w http.ResponseWriter, r *http.Request){ 

// }
/***********************************************************************************************************************************/

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

func AvailableETFs(w http.ResponseWriter, r *http.Request){  
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
    jsonData, err := json.Marshal(etfs)
    if err != nil{panic(err.Error())}
    w.Header().Set("Content-Type", "application/json")
	w.Write(jsonData)
}


func AvailableTickers(w http.ResponseWriter, r *http.Request){ 
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

// func signIn(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	w.Write(getJson([]byte(`{ "message":  "Hello from Weather App"  }`)))
// }

/********************************************************************************************************************************************/
// URLs 
func main(){ 
	fmt.Println("inside main")
   	models.DownloadData()
   	models.Csv_to_sql()
    router := mux.NewRouter()
    router.HandleFunc("/spdr", Signin)
    router.HandleFunc("/spdr/etfs", AvailableETFs).Methods("GET")
    router.HandleFunc("/spdr/etfs/tickers", AvailableTickers).Methods("GET")

    log.Fatal(http.ListenAndServe(":8080", router))
    if err := http.ListenAndServe(":8080", nil); err != nil {
		panic(err)
	}
    fmt.Println("main complete")
}






