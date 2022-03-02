package Controller

import (
	"github.com/mhthrh/ApiStore/Model/Book"
	"github.com/mhthrh/ApiStore/Model/Result"
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

// AddBook handles POST requests to add new Controller
func (b *Controller) AddBook(rw http.ResponseWriter, r *http.Request) {
	book := r.Context().Value(Key{}).(*Book.Book)
	b.l.Printf(b.e.SelectException(1004), book)
	Book.AddBook(*book)
	Result.New((*book).Id, 0, http.StatusOK, "Success", "").SendResponse(&rw)

}

// swagger:route GET /Controller Controller listBooks
// Return a list of Controller from disk
// responses:
//	200: booksResponse

// BookSList handles GET requests and returns all current Controller
func (b *Controller) BookSList(w http.ResponseWriter, r *http.Request) {
	b.l.Println(b.e.SelectException(1005), nil)

	Result.New(0, 0, http.StatusOK, "Success", JsonUtil.New(nil, nil).Struct2Json(Book.GetBooks())).SendResponse(&w)

}

// swagger:route GET /Controller/{id} Controller listSingle
// Return a list of Controller from disk
// responses:
//	200: bookResponse
//	404: errorResponse

// GetBook handles GET requests
func (b *Controller) GetBook(rw http.ResponseWriter, r *http.Request) {
	id := HttpUtil.GetID(r)
	b.l.Println(b.e.SelectException(1007), id)
	bo, err := Book.GetBookByID(id)

	switch err {
	case nil:

	case Book.BookNotFound:
		b.l.Println(b.e.SelectException(1008), err)
		Result.New(0, 0, http.StatusNotFound, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&rw)

		return
	default:
		b.l.Println(b.e.SelectException(1008), err)
		Result.New(0, 0, http.StatusInternalServerError, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&rw)

		return
	}

	Result.New(0, 0, http.StatusOK, "Success", JsonUtil.New(nil, nil).Struct2Json(bo)).SendResponse(&rw)

}

// swagger:route DELETE /Controller/{id} Controller deleteBook
// Update a Controller details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// DeleteBook handles DELETE requests and removes items from the database
func (b *Controller) DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := HttpUtil.GetID(r)

	b.l.Println(b.e.SelectException(1014), id)

	err := Book.DeleteBook(id)
	if err == Book.BookNotFound {
		b.l.Println(b.e.SelectException(1015))
		Result.New(0, -1, http.StatusNotFound, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&w)
		return
	}

	if err != nil {
		b.l.Println("[ERROR] deleting record", err)
		Result.New(0, -1, http.StatusInternalServerError, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&w)
		return
	}
	Result.New(0, 0, http.StatusOK, "Success", "").SendResponse(&w)

}

// swagger:route PUT /Controller Controller updateBook
// Update a Controller details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// UpdateBook handles PUT requests to update Controller
func (b *Controller) UpdateBook(w http.ResponseWriter, r *http.Request) {

	book := r.Context().Value(Key{}).(Book.Book)
	b.l.Println(b.e.SelectException(1010), book.Id)

	err := Book.UpdateBook(book)
	if err == Book.BookNotFound {
		b.l.Println(b.e.SelectException(1011), err)
		Result.New(0, -1, http.StatusNotFound, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: "Wine not found"})).SendResponse(&w)

		return
	}
	Result.New(0, 0, http.StatusOK, "Success", "").SendResponse(&w)
}
