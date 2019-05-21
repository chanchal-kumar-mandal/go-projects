package main

import (
    "fmt"
    "database/sql"
    _ "github.com/go-sql-driver/mysql"
)

func main() {

    // Database Connection
    db, err := sql.Open("mysql", "root:@tcp(127.0.0.1:3306)/")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Database created successfully")
    }

    // Create Database
    _,err = db.Exec("CREATE DATABASE go_mysql_connection_database")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Successfully created database..")
    }

    // Choose the Database
    _,err = db.Exec("USE go_mysql_connection_database")
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("DB selected successfully..")
    }

    // Create Table
    stmt, err := db.Prepare("CREATE Table employee(id int NOT NULL AUTO_INCREMENT, first_name varchar(50), last_name varchar(30), PRIMARY KEY (id));")
    if err != nil {
        fmt.Println(err.Error())
    }

    _, err = stmt.Exec()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Table employee created successfully..")
    }

    // Insert Data
    stmt, err := db.Prepare("INSERT INTO `employee` (`first_name`, `last_name`) VALUES ('Chanchal', 'Mandal');")
    if err != nil {
        fmt.Println(err.Error())
    }

    _, err = stmt.Exec()
    if err != nil {
        fmt.Println(err.Error())
    } else {
        fmt.Println("Inserted data into employee table successfully..")
    }

    // Close Database
    defer db.Close()
}