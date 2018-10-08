package main

import (
	"src/github.com/joho/godotenv"
	"testing"
)

func TestGetSheets(t *testing.T) {
	godotenv.Load("../.env.sh")

	result, err := getSheets()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*result) != 1 {
		t.Fatal("failed")
	}
}
