package main

import (
	"context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type fakeStorage struct {
}

func (fs fakeStorage) AddRacer(ctx context.Context, r *Racer) error {

	return nil
}

func (fs fakeStorage) AddSessionToRacer(ctx context.Context, ses session, racerID string) error {
	return nil
}

func (fs fakeStorage) GetRacer(ctx context.Context, name string) (Racer, error) {
	return Racer{}, nil
}

func TestCreateNewRacerHandler(t *testing.T) {

	srv := server{
		context:        context.Background(),
		storageService: fakeStorage{},
	}

	testCases := []struct {
		Name               string
		Body               string
		ExpectedStatusCode int
		ExpectedBody       string
	}{
		{
			Name:               "Send valid body, Created response",
			Body:               `{"name": "will"}`,
			ExpectedStatusCode: 201,
			ExpectedBody:       `{"ID":"000000000000000000000000","Name":"will","FastestLap":0,"Sessions":null}`,
		},
		{
			Name:               "Send invalid rbody, BadRequest response",
			Body:               `{`,
			ExpectedStatusCode: 400,
			ExpectedBody:       "Bad Request",
		},
	}
	for _, test := range testCases {

		req, err := http.NewRequest("POST", "/create-racer", strings.NewReader(test.Body))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(srv.CreateNewRacerHandler())

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.ExpectedStatusCode {
			t.Errorf("Wanted code %v got code %v", http.StatusCreated, status)
		}

		got := rr.Body.String()

		if strings.TrimSpace(got) != strings.TrimSpace(test.ExpectedBody) {
			t.Errorf("got %s but want %s", got, test.ExpectedBody)
		}
	}
}

func TestAddSessionDataHandler(t *testing.T) {

	srv := server{
		context:        context.Background(),
		storageService: fakeStorage{},
	}

	testCases := []struct {
		Name               string
		Body               string
		ExpectedStatusCode int
		ExpectedBody       string
	}{
		{
			Name:               "Send valid body, Created response",
			Body:               `{"id": "5d9dbcaa39c754e09e65aee7","session": {"lapTimes": [1,2,3]}}`,
			ExpectedStatusCode: 200,
			ExpectedBody:       ``,
		},
		{
			Name:               "Send invalid rbody, BadRequest response",
			Body:               `{`,
			ExpectedStatusCode: 400,
			ExpectedBody:       "Bad Request",
		},
	}
	for _, test := range testCases {

		req, err := http.NewRequest("POST", "/save-sessions", strings.NewReader(test.Body))

		if err != nil {
			t.Fatal(err)
		}

		rr := httptest.NewRecorder()
		handler := http.HandlerFunc(srv.AddSessionDataHandler())

		handler.ServeHTTP(rr, req)

		if status := rr.Code; status != test.ExpectedStatusCode {
			t.Errorf("Wanted code %v got code %v", http.StatusCreated, status)
		}

		got := rr.Body.String()

		if strings.TrimSpace(got) != strings.TrimSpace(test.ExpectedBody) {
			t.Errorf("got %s but want %s", got, test.ExpectedBody)
		}
	}
}
