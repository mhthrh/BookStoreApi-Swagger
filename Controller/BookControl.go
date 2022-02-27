package Controller

import (
	Book2 "github.com/mhthrh/ApiStore/Model/Book"
	"github.com/mhthrh/ApiStore/Utility/HttpUtil"
	"github.com/mhthrh/ApiStore/Utility/JsonUtil"
	"net/http"
)

// swagger:route POST /Controller Controller createBook
// Create a new book
//
// responses:
//	200: bookResponse
//  422: errorValidation
//  501: errorResponse

// Create handles POST requests to add new Controller
func (b *Controller) Create(rw http.ResponseWriter, r *http.Request) {
	bo := r.Context().Value(KeyBook{}).(Book2.Book)

	b.l.Printf(b.e.SelectException(1004), bo)
	Book2.AddBook(bo)
}

// swagger:route GET /Controller Controller listBooks
// Return a list of Controller from disk
// responses:
//	200: booksResponse

// ListAll handles GET requests and returns all current Controller
func (b *Controller) ListAll(w http.ResponseWriter, r *http.Request) {
	b.l.Println(b.e.SelectException(1005), nil)

	if err := JsonUtil.New(w, nil).ToJSON(Book2.GetBooks()); err != nil {
		b.l.Println(b.e.SelectException(1006), err)
	}
}

// swagger:route GET /Controller/{id} Controller listSingle
// Return a list of Controller from disk
// responses:
//	200: bookResponse
//	404: errorResponse

// ListSingle handles GET requests
func (b *Controller) ListSingle(rw http.ResponseWriter, r *http.Request) {
	id := HttpUtil.GetID(r)
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

// swagger:route DELETE /Controller/{id} Controller deleteBook
// Update a Controller details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// Delete handles DELETE requests and removes items from the database
func (b *Controller) Delete(w http.ResponseWriter, r *http.Request) {
	id := HttpUtil.GetID(r)

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

// swagger:route PUT /Controller Controller updateBook
// Update a Controller details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// Update handles PUT requests to update Controller
func (b *Controller) Update(w http.ResponseWriter, r *http.Request) {

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
