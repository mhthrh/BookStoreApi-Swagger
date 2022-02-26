package Controller

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	Book2 "github.com/mhthrh/ApiStore/Model/Book"
	"github.com/mhthrh/ApiStore/Utility/ExceptionUtil"
	"github.com/mhthrh/ApiStore/Utility/JsonUtil"
	"github.com/mhthrh/ApiStore/Utility/ValidationUtil"
	"github.com/sirupsen/logrus"
	"net/http"
	"strconv"
	"time"
)

// KeyBook is a key used for the book obj on context
type KeyBook struct{}

// books handler for getting and updating books
type books struct {
	l *logrus.Entry
	v *ValidationUtil.Validation
	e *ExceptionUtil.Exception
}

// NewBooks returns a new book handler with  logger
func NewBooks(l *logrus.Entry, v *ValidationUtil.Validation, e *ExceptionUtil.Exception) *books {
	return &books{l, v, e}
}

// InvalidPath is an error message when the book path is not valid
var InvalidPath = fmt.Errorf("invalid Path, path must be /books/[id]")

// GenericError is a generic error message returned by a server
type GenericError struct {
	Message string `json:"message"`
}

// ValidationError is a collection of validation error messages
type ValidationError struct {
	Messages []string `json:"messages"`
}

// getID returns the book ID from the URL
// Panics if it cannot convert the id into an integer
// this must never happen as the router ensures that
// this is a valid number
func getID(r *http.Request) int {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		panic(err)
	}
	return id
}

// HttpMiddleware validates the book in the request and calls next if ok
func (b *books) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var bo Book2.Book
		j := JsonUtil.New(w, r.Body)
		err := j.FromJSON(&bo)
		if err != nil {
			start := time.Now()
			b.l.WithFields(map[string]interface{}{
				"method":     r.Method,
				"path":       r.URL,
				"status":     nil,
				"latency_ns": time.Since(start).Nanoseconds(),
			}).Info("request details")
			w.WriteHeader(http.StatusBadRequest)
			if j.ToJSON(&GenericError{Message: err.Error()}) != nil {
				b.l.Errorf(b.e.SelectException(1), err)
			}
			return
		}

		errs := b.v.Validate(bo)
		if len(errs) != 0 {
			b.l.Println(b.e.SelectException(3), err)

			w.WriteHeader(http.StatusUnprocessableEntity)
			j.ToJSON(&ValidationError{Messages: errs.Errors()})
			return
		}

		ctx := context.WithValue(r.Context(), KeyBook{}, bo)
		r = r.WithContext(ctx)

		next.ServeHTTP(w, r)
	})
}

// swagger:route POST /books books createBook
// Create a new book
//
// responses:
//	200: bookResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new books
func (b *books) Create(rw http.ResponseWriter, r *http.Request) {
	bo := r.Context().Value(KeyBook{}).(Book2.Book)

	b.l.Printf(b.e.SelectException(1004), bo)
	Book2.AddBook(bo)
}

// swagger:route GET /books books listBooks
// Return a list of books from disk
// responses:
//	200: booksResponse

// ListAll handles GET requests and returns all current books
func (b *books) ListAll(w http.ResponseWriter, r *http.Request) {
	b.l.Println(b.e.SelectException(1005), nil)

	if err := JsonUtil.New(w, nil).ToJSON(Book2.GetBooks()); err != nil {
		b.l.Println(b.e.SelectException(1006), err)
	}
}

// swagger:route GET /books/{id} books listSingle
// Return a list of books from disk
// responses:
//	200: bookResponse
//	404: errorResponse

// ListSingle handles GET requests
func (b *books) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getID(r)
	b.l.Println(b.e.SelectException(1007), id)
	bo, err := Book2.GetBookByID(id)

	switch err {
	case nil:

	case Book2.BookNotFound:
		b.l.Println(b.e.SelectException(1008), err)

		rw.WriteHeader(http.StatusNotFound)
		JsonUtil.New(rw, nil).ToJSON(&GenericError{Message: err.Error()})
		return
	default:
		b.l.Println(b.e.SelectException(1008), err)

		rw.WriteHeader(http.StatusInternalServerError)
		JsonUtil.New(rw, nil).ToJSON(&GenericError{Message: err.Error()})
		return
	}

	err = JsonUtil.New(rw, nil).ToJSON(bo)
	if err != nil {
		b.l.Println(b.e.SelectException(1009), err)
	}
}

// swagger:route DELETE /books/{id} books deleteBook
// Update a books details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (b *books) Delete(w http.ResponseWriter, r *http.Request) {
	id := getID(r)

	b.l.Println(b.e.SelectException(1014), id)

	err := Book2.DeleteBook(id)
	if err == Book2.BookNotFound {
		b.l.Println(b.e.SelectException(1015))

		w.WriteHeader(http.StatusNotFound)
		JsonUtil.New(w, nil).ToJSON(&GenericError{Message: err.Error()})
		return
	}

	if err != nil {
		b.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		JsonUtil.New(w, nil).ToJSON(&GenericError{Message: err.Error()})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}

// swagger:route PUT /books books updateBook
// Update a books details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update books
func (b *books) Update(w http.ResponseWriter, r *http.Request) {

	pr := r.Context().Value(KeyBook{}).(Book2.Book)
	b.l.Println(b.e.SelectException(1010), pr.Id)

	err := Book2.UpdateBook(pr)
	if err == Book2.BookNotFound {
		b.l.Println(b.e.SelectException(1011), err)
		w.WriteHeader(http.StatusNotFound)
		JsonUtil.New(w, nil).ToJSON(&GenericError{Message: "book not found"})
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
