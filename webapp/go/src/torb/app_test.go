package main

import (
	"database/sql"
	"fmt"
	"os"
	"src/github.com/joho/godotenv"
	"src/github.com/labstack/gommon/log"
	"testing"
)

func TestGetSheets(t *testing.T) {
	godotenv.Load("../.env.sh")
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?parseTime=true&charset=utf8mb4",
		os.Getenv("DB_USER"), os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"),
		os.Getenv("DB_DATABASE"),
	)

	var err error
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		log.Fatal(err)
	}

	result, err := getSheets()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*result) != 1 {
		t.Fatal("failed")
	}
}
