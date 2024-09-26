package config

import (
    "database/sql"
    "fmt"
    "log"
    "os"
    _ "github.com/lib/pq"
    "github.com/joho/godotenv"
)

var DB *sql.DB

func InitDB() {
    // 加载 .env 文件
    err := godotenv.Load()
    if err != nil {
        log.Fatal("Error loading .env file")
    }

    // 从环境变量中获取数据库连接信息
    dbHost := os.Getenv("DB_HOST")
    dbUser := os.Getenv("DB_USER")
    dbPassword := os.Getenv("DB_PASSWORD")
    dbName := os.Getenv("DB_NAME")
    dbPort := os.Getenv("DB_PORT")
    dbSSLMode := os.Getenv("DB_SSLMODE")
    dbTimeZone := os.Getenv("DB_TIMEZONE")

    // 构建连接字符串
    connStr := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
        dbHost, dbUser, dbPassword, dbName, dbPort, dbSSLMode, dbTimeZone)

    // 打开数据库连接
    DB, err = sql.Open("postgres", connStr)
    if err != nil {
        log.Fatal("Failed to connect to database:", err)
    }

    // 测试数据库连接
    err = DB.Ping()
    if err != nil {
        log.Fatal("Failed to ping database:", err)
    }

    fmt.Println("Database connected")
}
