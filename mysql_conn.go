package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// connecting to local mysql (xampp) with root username and empty password
var db, err = sql.Open("mysql", "root:@tcp(localhost:3306)/booking_app")

func MysqlConnectAdd(firstName string, lastName string, email string, tickets uint) {
	//checking if parameters were typed correctly
	if err != nil {
		fmt.Println("Error connecting to the MYSQL")
		log.Fatal(err)
	}
	defer db.Close()

	//checking if the connection to the database was successsful
	err = db.Ping()
	if err != nil {
		fmt.Println("Error verifying connection with Database.")
		log.Fatal(err)
	}

	//this functions adds entry of the client to the database
	//preparing a query to insert values from variables
	insert, err := db.Prepare("INSERT INTO users (name, surname, email, bought_tickets) VALUES (?, ?, ?, ?)")
	//executing SQL command with variables provided from the user
	insert.Exec(firstName, lastName, email, tickets)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()

}

func MysqlConnectDel(id int) {
	//checking if parameters were typed correctly
	if err != nil {
		fmt.Println("Error connecting to the MYSQL")
		log.Fatal(err)
	}
	defer db.Close()

	//checking if the connection to the database was successsful
	err = db.Ping()
	if err != nil {
		fmt.Println("Error verifying connection with Database.")
		log.Fatal(err)
	}

	delete, err := db.Prepare("DELETE FROM users WHERE id=(?)")
	delete.Exec(id)

	if err != nil {
		log.Fatal(err)
	}
	defer delete.Close()

}

func MysqlConnectShow() {
	var (
		id      int
		name    string
		surname string
		email   string
		tickets int
	)
	//checking if parameters were typed correctly
	if err != nil {
		fmt.Println("Error connecting to the MYSQL")
		log.Fatal(err)
	}
	defer db.Close()

	//checking if the connection to the database was successsful
	err = db.Ping()
	if err != nil {
		fmt.Println("Error verifying connection with Database.")
		log.Fatal(err)
	}
	//creating a query to show all of the users
	show, err := db.Query("SELECT * FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer show.Close()
	//looping through database
	for show.Next() {
		//with this line we scan trough all of the database and we put each value into a variable
		err := show.Scan(&id, &name, &surname, &email, &tickets)
		if err != nil {
			log.Fatal(err)
		}
		//outputing the variables
		log.Println(id, name, surname, email, tickets)
	}
	err = show.Err()
	if err != nil {
		log.Fatal(err)
	}
}
