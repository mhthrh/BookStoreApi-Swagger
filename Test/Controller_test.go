package Test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/Controller"
	"github.com/mhthrh/ApiStore/Model/Book"
	"net/http"
	"net/http/httptest"
	"testing"
	"time"
)

var (
	b = Book.Book{
		Id:          232,
		ISBN:        "1234567878",
		Title:       "222",
		Authors:     []string{"65654", "3243"},
		Publisher:   "tokhmi",
		PublishDate: time.Now(),
		Pages:       245,
	}
)

func TestAddBook(t *testing.T) {
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&b)
	req, _ := http.NewRequest("POST", "/add", &data)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(Controller.AddBook)

	handler.ServeHTTP(r, req)
	fmt.Println(r.Body)
	checkStatusCode(r.Code, http.StatusCreated, t)

}

func TestFindBook(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/book/findbook", nil)
	req = mux.SetURLVars(req, map[string]string{"isbn": "123"})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(Controller.Book)

	handler.ServeHTTP(r, req)
	fmt.Println(r.Body)
	checkStatusCode(r.Code, http.StatusCreated, t)

}

func TestDeleteBook(t *testing.T) {
	req, _ := http.NewRequest("GET", "/api/book/delete", nil)
	req = mux.SetURLVars(req, map[string]string{"isbn": "123"})
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(Controller.Book)

	handler.ServeHTTP(r, req)
	fmt.Println(r.Body)
	checkStatusCode(r.Code, http.StatusCreated, t)

}

func TestUpdateBook(t *testing.T) {
	var data bytes.Buffer
	json.NewEncoder(&data).Encode(&b)
	req, _ := http.NewRequest("GET", "/api/book/update", &data)
	r := httptest.NewRecorder()
	handler := http.HandlerFunc(Controller.Book)
	handler.ServeHTTP(r, req)
	fmt.Println(r.Body)
	checkStatusCode(r.Code, http.StatusCreated, t)

}

func checkStatusCode(code int, want int, t *testing.T) {
	if code != want {
		t.Errorf("Wrong status code: got %v want %v", code, want)
	}
}
