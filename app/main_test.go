package main

import (
	"bytes"
	"context"
	"encoding/json"
	"log"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"

	"bitbucket.org/oaroz/hai/app/model"
)

const url string = "/api/messages"

const tableCreationQuery string = `
CREATE TABLE IF NOT EXISTS messages (
    id bigserial PRIMARY KEY,
    code varchar(100) NOT NULL,
    title varchar(255) NOT NULL,
    content text NOT NULL,
    email varchar(100) NOT NULL
);`

var app App

func ensureTableExists() {
	if _, err := app.Db.Exec(context.Background(), tableCreationQuery); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	app.Db.Exec(context.Background(), "DELETE FROM messages")
	app.Db.Exec(context.Background(), "ALTER SEQUENCE messages_id_seq RESTART WITH 1")
}

func executeRequest(req *http.Request) *httptest.ResponseRecorder {
	rr := httptest.NewRecorder()
	app.Router.ServeHTTP(rr, req)
	return rr
}

func checkHttpCode(t *testing.T, expected, actual int) {
	if expected != actual {
		t.Errorf("Expected response code %d. But is %d\n", expected, actual)
	}
}

func TestMain(m *testing.M) {
	app = App{}
	app.Init("../config/app_test.yml")
	ensureTableExists()
	code := m.Run()
	clearTable()
	os.Exit(code)
}

func TestEmptyTable(t *testing.T) {
	clearTable()

	req, _ := http.NewRequest("GET", url, nil)
	res := executeRequest(req)

	checkHttpCode(t, http.StatusOK, res.Code)

	var body string = res.Body.String()
	if strings.TrimSpace(body) != "[]" {
		t.Errorf("Expected an empty array. But is %s", body)
	}
}

func TestCreateMessage(t *testing.T) {

	clearTable()

	var jsonStr = []byte(`{"email":"jan.kowalski@example.com","title":"Interview 2","content":"simple text 2"}`)
	req, _ := http.NewRequest("POST", url, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	res := executeRequest(req)
	checkHttpCode(t, http.StatusCreated, res.Code)

	var m model.CreateMessageResponse
	json.Unmarshal(res.Body.Bytes(), &m)

	if m.Code == "" {
		t.Errorf("Expected message code to be not empty but is %v", m.Code)
	}

	if len(m.Code) != 20 {
		t.Errorf("Expected message code to be 20 chars long but is %v", len(m.Code))
	}

	if m.Title != "Interview 2" {
		t.Errorf("Expected message title to be Interview 2 but is %v", m.Title)
	}

	if m.Email != "jan.kowalski@example.com" {
		t.Errorf("Expected message email to be jan.kowalski@example.com but is %v", m.Email)
	}

	if res.Code != http.StatusCreated {
		t.Errorf("Expected status http code 201 but is %v", res.Code)
	}
}
