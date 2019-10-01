package main

import (
	"net/http"
	"strings"
	"testing"
)

func TestDecodeBody(t *testing.T) {

	body := `{"Test":"Hello"}`
	req, err := http.NewRequest("GET", "/test", strings.NewReader(body))

	if err != nil {
		t.Fatal(err)
	}

	type Dummy struct {
		Test string
	}

	var got Dummy

	err = decodeBody(req, &got)

	if err != nil {
		t.Fatal(err)
	}

	if got.Test != "Hello" {
		t.Errorf("got %v want Hello", got.Test)
	}
}
