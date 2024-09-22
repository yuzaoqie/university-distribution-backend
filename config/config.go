package config

import (
    "database/sql"
    "fmt"
    "log"
    _ "github.com/lib/pq"
)

var DB *sql.DB

func InitDB() {
    var err error
    connStr := "host=localhost user=postgres password=123456 dbname=university_data port=5433 sslmode=disable TimeZone=Asia/Shanghai"
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    err = DB.Ping()
    if err != nil {
        log.Fatal("Failed to ping database:", err)
    }
    fmt.Println("Database connected")
}
