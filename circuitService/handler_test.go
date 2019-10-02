package main

import (
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakePublisher struct {
}

func (fp fakePublisher) publishMessage(body []byte, keyName string) error {

	return nil
}

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

func TestStartSessionHandler(t *testing.T) {

	srv := server{
		publisher: fakePublisher{},
	}

	testCases := []struct {
		Name               string
		Body               string
		ExpectedStatusCode int
		ExpectedBody       string
	}{
		{
			Name:               "Send valid race details, OK response",
			Body:               validRequestBody,
			ExpectedStatusCode: 200,
			ExpectedBody:       "",
		},
		{
			Name:               "Send invalid race session time, BadRequest response",
			Body:               invalidSessionTimeRequestBody,
			ExpectedStatusCode: 400,
			ExpectedBody:       "Session time must be greater than 0",
		},
		{
			Name:               "Send no racers, BadRequest response",
			Body:               noRacersRequestBody,
			ExpectedStatusCode: 400,
			ExpectedBody:       "There must be racers in the session",
		},
		{
			Name:               "Send invalid request body, BadRequest response",
			Body:               "{",
			ExpectedStatusCode: 400,
			ExpectedBody:       "Bad Request",
		},
	}

	for _, test := range testCases {

		req, err := http.NewRequest("POST", "/start", strings.NewReader(test.Body))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(srv.StartSessionHandler())

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.ExpectedStatusCode {
			t.Errorf("Wanted code %v got code %v", test.ExpectedStatusCode, status)
		}

		got := rr.Body.String()

		if strings.TrimSpace(got) != strings.TrimSpace(test.ExpectedBody) {
			t.Errorf("got %s but want %s", got, test.ExpectedBody)
		}
	}
}

var validRequestBody = `{   
    "SessionTime" : 10,
    "Racers": [
    	{
    		"ID" : "952279ba-bb97-4e2d-9950-2da8481bebef",
    		"KartNumber": 1
    	},
    	{
    		"ID" : "e2c2e4cb-e8fd-4bca-b03c-0ea92a1a6246",
    		"KartNumber": 2
    	}
    	]
}`

var invalidSessionTimeRequestBody = `{   
    "SessionTime" : 0,
    "Racers": [
    	{
    		"ID" : "952279ba-bb97-4e2d-9950-2da8481bebef",
    		"KartNumber": 1
    	},
    	{
    		"ID" : "e2c2e4cb-e8fd-4bca-b03c-0ea92a1a6246",
    		"KartNumber": 2
    	}
    	]
}`

var noRacersRequestBody = `{   
    "SessionTime" : 20,
    "Racers": []
}`
