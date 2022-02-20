package Controller

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/Helper"
	"github.com/mhthrh/ApiStore/Helper/Validation"
	Book2 "github.com/mhthrh/ApiStore/Model/Book"
	"log"
	"net/http"
	"strconv"
)

// KeyBook is a key used for the book obj on context
type KeyBook struct{}

// books handler for getting and updating books
type books struct {
	l *log.Logger
	v *Validation.Validation
}

// NewBooks returns a new book handler with  logger
func NewBooks(l *log.Logger, v *Validation.Validation) *books {
	return &books{l, v}
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

// getBookID returns the book ID from the URL
// Panics if it cannot convert the id into an integer
// this must never happen as the router ensures that
// this is a valid number
func getBookID(r *http.Request) int {
	// parse the book id from the url
	vars := mux.Vars(r)
	// convert the id into an integer and return
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		// never must happen
		panic(err)
	}
	return id
}

// MiddlewareValidateBook validates the book in the request and calls next if ok
func (b *books) MiddlewareValidateBook(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		bo := &Book2.Books{}

		err := Helper.FromJSON(bo, r.Body)
		if err != nil {
			b.l.Println("[ERROR] deserializing book", err)

			w.WriteHeader(http.StatusBadRequest)
			if Helper.ToJSON(&GenericError{Message: err.Error()}, w) != nil {
				b.l.Println("[ERROR] deserializing book", err)
			}
			return
		}

		// validate the book
		errs := b.v.Validate(bo)
		if len(errs) != 0 {
			b.l.Println("[ERROR] validating book", errs)

			// return the validation messages as an array
			w.WriteHeader(http.StatusUnprocessableEntity)
			Helper.ToJSON(&ValidationError{Messages: errs.Errors()}, w)
			return
		}

		// add the book to the context
		ctx := context.WithValue(r.Context(), KeyBook{}, bo)
		r = r.WithContext(ctx)

		// Call the next handler, which can be another middleware in the chain, or the final handler.
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
	// fetch the book from the context
	bo := r.Context().Value(KeyBook{}).(Book2.Book)

	b.l.Printf("[DEBUG] Inserting book: %#v\n", bo)
	Book2.AddBook(bo)
}

// swagger:route GET /books books listBooks
// Return a list of books from disk
// responses:
//	200: booksResponse

// ListAll handles GET requests and returns all current books
func (b *books) ListAll(w http.ResponseWriter, r *http.Request) {
	b.l.Println("[DEBUG] get all records")

	bo := Book2.GetBooks()

	err := Helper.ToJSON(bo, w)
	if err != nil {
		// we should never be here but log the error just in case
		b.l.Println("[ERROR] serializing book", err)
	}
}

// swagger:route GET /books/{id} books listSingle
// Return a list of books from disk
// responses:
//	200: bookResponse
//	404: errorResponse

// ListSingle handles GET requests
func (b *books) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := getBookID(r)

	b.l.Println("[DEBUG] get record id", id)

	bo, err := Book2.GetBookByID(id)

	switch err {
	case nil:

	case Book2.BookNotFound:
		b.l.Println("[ERROR] fetching book", err)

		rw.WriteHeader(http.StatusNotFound)
		Helper.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	default:
		b.l.Println("[ERROR] fetching book", err)

		rw.WriteHeader(http.StatusInternalServerError)
		Helper.ToJSON(&GenericError{Message: err.Error()}, rw)
		return
	}

	err = Helper.ToJSON(bo, rw)
	if err != nil {
		// we should never be here but log the error just in case
		b.l.Println("[ERROR] serializing book", err)
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
	id := getBookID(r)

	b.l.Println("deleting record id", id)

	err := Book2.DeleteBook(id)
	if err == Book2.BookNotFound {
		b.l.Println("[ERROR] deleting record id does not exist")

		w.WriteHeader(http.StatusNotFound)
		Helper.ToJSON(&GenericError{Message: err.Error()}, w)
		return
	}

	if err != nil {
		b.l.Println("[ERROR] deleting record", err)

		w.WriteHeader(http.StatusInternalServerError)
		Helper.ToJSON(&GenericError{Message: err.Error()}, w)
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

	// fetch the book from the context
	pr := r.Context().Value(KeyBook{}).(Book2.Book)
	b.l.Println("[DEBUG] updating record id", pr.Id)

	err := Book2.UpdateBook(pr)
	if err == Book2.BookNotFound {
		b.l.Println("[ERROR] book not found", err)

		w.WriteHeader(http.StatusNotFound)
		Helper.ToJSON(&GenericError{Message: "book not found"}, w)
		return
	}

	// write the no content success header
	w.WriteHeader(http.StatusNoContent)
}

func minmax(array []Book2.Book) (int, int) {
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
