package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"
	"testing"
)

func TestGetSheets(t *testing.T) {
	os.Setenv("DB_DATABASE", "torb")
	os.Setenv("DB_HOST", "127.0.0.1")
	os.Setenv("DB_PORT", "3306")
	os.Setenv("DB_USER", "isucon")
	os.Setenv("DB_PASS", "isucon")
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
	if len(*result) != 1000 {
		t.Fatalf("failed: result length is expected is 1000, but got %d", len(*result))
	}

	for _, sheet := range *result {
		fmt.Printf("%#v", sheet)
	}
}
