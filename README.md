
### Go Application that parses SPDR ETFs (​www.spdrs.com​) and stores into the database. The data can be accessed using REST APIs. 

## Technologies: Go, mariadb

## Setting up Environment

1. Install Go using the link https://golang.org/dl/. Install packages using using following commangs:
```bash
go get -u github.com/gorilla/mux
```
```bash
go get -u github.com/go-sql-driver/mysql
```
```bash
go get -u github.com/dgrijalva/jwt-go
```

Refer notes.md to create database

## running the application
```bash
go run main.go
```

## APIs

### Any HTTP client with support for cookies (like Postman, or your web browser) can be used to access the APIs:

## 1. List of available ETF symbol
```bash
http://localhost:8080/spdr/etfs

GET {"username":"user1","password":"password1"}
```

## 2. Get data for ETF by ticker
```bash
http://localhost:8080/spdr/etfs/tickers

GET {"username":"user1","password":"password1"}
```

## stoping the server
```bash
go run exit.go
```