package config

import (
    "database/sql"
    "log"

    _ "github.com/go-sql-driver/mysql" // Import MySQL driver
)

var DB *sql.DB

func InitDB() {
    var err error
    // Initialize your database connection
    connectionString := "root:root@tcp(localhost:3306)/employee_db"
    // Adjust connection string as needed
    DB, err = sql.Open("mysql", connectionString)
    if err != nil {
        log.Fatalf("Error connecting to database: %v", err)
    }

    // Test the connection
    err = DB.Ping()
    if err != nil {
        log.Fatalf("Error pinging database: %v", err)
    }

    log.Println("Connected to database")
}
