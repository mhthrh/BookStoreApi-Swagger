package Controller

import (
	"encoding/json"
	"github.com/gorilla/mux"
	Book2 "github.com/mhthrh/ApiStore/Model/Book"
	"github.com/mhthrh/ApiStore/Model/Result"
	"net/http"
	"strconv"
	"strings"
)

func AllBooks(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	search, ok := vars["range"]
	if !ok {
		Result.New(0, 0, 100101, "bad request", nil).Create(&w, http.StatusBadRequest)
		return
	}

	i, err := strconv.Atoi(strings.Split(search, "-")[0])
	if err != nil {
		Result.New(0, 0, 100102, "min val check", nil).Create(&w, http.StatusBadRequest)
		return
	}
	j, err := strconv.Atoi(strings.Split(search, "-")[1])
	if err != nil {
		Result.New(0, 0, 100103, "max val check", nil).Create(&w, http.StatusBadRequest)
		return
	}
	min, max := minmax(*Book2.Books)
	if j < i || j-i > 100 {
		Result.New(min, max, 100104, "Overflow", nil).Create(&w, http.StatusBadRequest)
		return
	}
	Result.New(min, max, 100100, "Success", Book2.Load(i, j)).Create(&w, http.StatusOK)
}

func Book(w http.ResponseWriter, r *http.Request) {
	var array []Book2.Book
	vars := mux.Vars(r)
	search, ok := vars["isbn"]
	if !ok {
		Result.New(0, 0, 100105, "bad request", nil).Create(&w, http.StatusBadRequest)
		return
	}
	b, _ := Book2.Find(search)
	if b == nil {
		Result.New(0, 0, 100106, "Not found", nil).Create(&w, http.StatusNotFound)
		return
	}
	array = append(array, *b)
	Result.New(1, 1, 100100, "success", &array).Create(&w, http.StatusOK)
}

func AddBook(w http.ResponseWriter, r *http.Request) {
	var b Book2.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		Result.New(0, 0, 100107, "Check input(s)", nil).Create(&w, http.StatusBadRequest)
		return
	}
	err = Book2.Add(b)
	if err != nil {
		Result.New(0, 0, 100108, err.Error(), nil).Create(&w, http.StatusNotAcceptable)
		return
	}
	Result.New(1, 1, 100100, "book added", nil).Create(&w, http.StatusOK)
}

func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var b Book2.Book
	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		Result.New(0, 0, 1001009, "Check input(s)", nil).Create(&w, http.StatusBadRequest)
		return
	}
	err = Book2.Update(b)
	if err != nil {
		Result.New(0, 0, 100110, err.Error(), nil).Create(&w, http.StatusBadRequest)
		return
	}
	Result.New(1, 1, 100100, "success", nil).Create(&w, http.StatusOK)
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	search, ok := vars["isbn"]
	if !ok {
		Result.New(0, 0, 100111, "Check input(s)", nil).Create(&w, http.StatusBadRequest)
		return
	}

	err := Book2.Delete(search)
	if err != nil {
		Result.New(0, 0, 100112, err.Error(), nil).Create(&w, http.StatusNotAcceptable)
		return
	}
	Result.New(1, 1, 100100, "success", nil).Create(&w, http.StatusOK)
}
func minmax(array []Book2.Book) (uint32, uint32) {
	var max = array[0].Id
	var min = array[0].Id
	for _, value := range array {
		if max < value.Id {
			max = value.Id
		}
		if min > value.Id {
			min = value.Id
		}
	}
	return min, max
}
