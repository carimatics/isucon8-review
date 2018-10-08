package main

import (
	"testing"
)

func TestGetSheets(t *testing.T) {
	result, err := getSheets()
	if err != nil {
		t.Fatalf("failed test %#v", err)
	}
	if len(*result) != 1 {
		t.Fatal("failed")
	}
}
