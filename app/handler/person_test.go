package handler

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rakmaster/goapi/app/db"
	"github.com/rakmaster/goapi/app/model"
	"github.com/rakmaster/goapi/config"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/x/bsonx"
)

const succeed = "\u2713"
const failed = "\u2717"

func handleRequest(db *mongo.Database, handler func(db *mongo.Database, w http.ResponseWriter, r *http.Request)) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		handler(db, w, r)
	}
}

func createNewRequestNewRecorder(method, endpoint string, body io.Reader) (*http.Request, *httptest.ResponseRecorder) {
	req, _ := http.NewRequest("POST", "/person", body)
	rr := httptest.NewRecorder()
	return req, rr
}

func TestCreatePerson(t *testing.T) {
	configuration := config.NewConfig()
	dbConnection := db.InitialConnection("test", configuration.MongoURI())

	person, _ := json.Marshal(model.NewPerson("john", "doe", "john_doe", "john@gmail.com", nil))

	httpHandler := http.HandlerFunc(handleRequest(dbConnection, CreatePerson))

	// check http created status
	req, rr := createNewRequestNewRecorder("POST", "/person", bytes.NewBuffer(person))
	httpHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusCreated {
		t.Errorf("%s check StatusCreated is failed: got %d want %d", failed, status, http.StatusCreated)
	} else {
		t.Logf("%s check StatusCreated is successfull.", succeed)
	}

	// check http bad request status
	req, rr = createNewRequestNewRecorder("POST", "/person", bytes.NewBuffer([]byte("{'username': 'download'email:'john@gmail.com'}"))) // wrong json body
	httpHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusBadRequest {
		t.Errorf("%s check StatusBadRequest return wrong status code: got %d want %d", failed, status, http.StatusBadRequest)
	} else {
		t.Logf("%s check StatusBadRequest is successfull.", succeed)
	}

	// check http not acceptable status
	// username or email that you sent already exists collection
	// username and email will be unique.
	keys := bsonx.Doc{
		{Key: "username", Value: bsonx.Int32(1)},
		{Key: "email", Value: bsonx.Int32(1)},
	}
	people := dbConnection.Collection("people")
	db.SetIndexes(people, keys)

	req, rr = createNewRequestNewRecorder("POST", "/person", bytes.NewBuffer(person))
	httpHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusNotAcceptable {
		t.Errorf("%s check StatusNotAcceptable return wrong status code: got %d want %d", failed, status, http.StatusNotAcceptable)
	} else {
		t.Logf("%s check StatusNotAcceptable is successfull.", succeed)
	}

	// check http internal server error status
	// drop database and close connection for raise internal server error
	dbConnection.Drop(nil)
	dbConnection.Client().Disconnect(nil)
	req, _ = http.NewRequest("POST", "/person", bytes.NewBuffer(person))
	rr = httptest.NewRecorder()
	httpHandler.ServeHTTP(rr, req)
	if status := rr.Code; status != http.StatusInternalServerError {
		t.Errorf("%s check StatusInternalServerError is failed: got %d want %d", failed, status, http.StatusInternalServerError)
	} else {
		t.Logf("%s check StatusInternalServerError is successfull.", succeed)
	}
}
