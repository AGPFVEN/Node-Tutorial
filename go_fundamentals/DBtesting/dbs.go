package main

//user: ellicn8kupq98dtk3jk7
//password: pscale_pw_9s98hoDsemF1hsQsrrM0vkAmGIHmoAHreSY4go3qTCR
//.env DSN=ellicn8kupq98dtk3jk7:pscale_pw_9s98hoDsemF1hsQsrrM0vkAmGIHmoAHreSY4go3qTCR@tcp(aws.connect.psdb.cloud)/wss_test?tls=true

import (
    "database/sql"
    "log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
    //Connect
    db, err := sql.Open("mysql", dbcred)
    if err != nil {
        log.Fatalf("failed to connect: %v", err)
    }
    defer db.Close()

    //Verying connection
    if err := db.Ping(); err != nil {
        log.Fatalf("failed to ping: %v", err)
    }
    log.Println("Successfully connected to PlanetScale!")

    //Insert row

}