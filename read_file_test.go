package main

import (
	"encoding/json"
	"testing"
)

func TestReadFile(t *testing.T) {
	fileName := "handout.pdf"
	readFileInput := ReadFileInput{
		Path: fileName,
	}

	b, err := json.Marshal(readFileInput)
	if err != nil {
		t.Fatal(err)
	}

	if _, err := ReadFile(b); err != nil {
		t.Fatal(err)
	}
}
