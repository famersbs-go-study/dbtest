package main

import (
    "database/sql"
    "fmt"
    _"github.com/go-sql-driver/mysql"
)

func main(){

	/*
	production:
	  adapter: mysql
	  database: redmine
	  host: localhost
	  username: redmine
	  password: Netand1410
	  encoding: utf8

	 */

	db, err := sql.Open("mysql", "redmine:pass@tcp(192.168.0.203:3306)/redmine")

	if err != nil {
        panic("a" +err.Error())  // Just for example purpose. You should use proper error handling instead of panic
    }
    defer db.Close()


    // Prepare statement for reading data
    stmtOut, err := db.Prepare("select count(*) from users")
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    defer stmtOut.Close()


    var squareNum int // we "scan" the result in here

    // Query the square-number of 13
    err = stmtOut.QueryRow().Scan(&squareNum)
    if err != nil {
        panic(err.Error()) // proper error handling instead of panic in your app
    }
    fmt.Printf("user cnt is : %d\n", squareNum)
}