package main

import (
    "os"
    "fmt"
    "github.com/ziutek/mymysql/mysql"
    _ "github.com/ziutek/mymysql/native" // Native engine
    // _ "github.com/ziutek/mymysql/thrsafe" // Thread safe engine
)

func main() {
    db := mysql.New("tcp", "", "192.168.0.203:3306", "redmine", "pass", "redmine")

    err := db.Connect()
    if err != nil {
        panic(err)
    }

    res, err := db.Start("select count(*) from users")
    if err != nil {
        panic(err)
    }

    for _, field := range res.Fields() {
        fmt.Print(field.Name, " ")
    }
    fmt.Println()

        // Print all rows
    for {
        row, err := res.GetRow()
            //checkError(err)
            if err != nil {
                panic(err)
            }

            if row == nil {
                // No more rows
                break
            }

        // Print all cols
        for _, col := range row {
            if col == nil {
                fmt.Print("<NULL>")
            } else {
                os.Stdout.Write(col.([]byte))
            }
            fmt.Print(" ")
        }
        fmt.Println()
    }
}