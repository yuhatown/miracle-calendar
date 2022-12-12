package model

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func GetRoutine() {
	err := godotenv.Load(".env")
	if err != nil {
  		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()


	var cal [6]string
	err = db.QueryRow("SELECT date FROM test").Scan(&cal)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cal)
}

func PostRoutine(routine, memo string) {
	err := godotenv.Load(".env")
	if err != nil {
  		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	

	err = db.QueryRow("INSERT INTO test(date, routine, memo, is_achieved)) VALUES (now(), ?, ?,0)", routine, memo)
	if err != nil {
		log.Fatal(err)
	}
}

func PutRoutine(routine, memo string, is_achieved, id int) {
	err := godotenv.Load(".env")
	if err != nil {
  		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.QueryRow("UPDATE test SET routine=?, memo=?, is_achieved=? WHERE id=?", routine, memo, is_achieved, id)
	if err != nil {
		log.Fatal(err)
	}
}

func DeleteRoutine(id int) {
	err := godotenv.Load(".env")
	if err != nil {
  		log.Fatal("Error loading .env file")
	}

	db, err := sql.Open("mysql", os.Getenv("DATA_SOURCE_NAME"))
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	err = db.QueryRow("DELETE FROM test WHERE id=?", id)
	if err != nil {
		log.Fatal(err)
	}
}