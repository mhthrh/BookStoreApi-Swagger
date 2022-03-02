package Controller

import (
	"github.com/mhthrh/ApiStore/Model/Result"
	"github.com/mhthrh/ApiStore/Model/Wine"
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

// AddWine handles POST requests to add new Controller
func (b *Controller) AddWine(rw http.ResponseWriter, r *http.Request) {
	db := b.db.Pull()
	w := Wine.New(db.Db)
	wine := r.Context().Value(Key{}).(*Wine.Wine)

	b.l.Printf(b.e.SelectException(1004), wine)
	w.Add(*wine)
	b.db.Push(db)
	Result.New((*wine).ID, 0, http.StatusOK, "Success", "").SendResponse(&rw)
}

// swagger:route GET /Controller Controller listBooks
// Return a list of Controller from disk
// responses:
//	200: booksResponse

// WineSList handles GET requests and returns all current Controller
func (b *Controller) WineSList(w http.ResponseWriter, r *http.Request) {
	b.l.Println(b.e.SelectException(1005), nil)
	db := b.db.Pull()
	wine := Wine.New(db.Db)
	wines, err := wine.List(0, 0)
	if err != nil {

	}
	b.db.Push(db)
	Result.New(0, 0, http.StatusOK, "Success", JsonUtil.New(w, nil).Struct2Json(wines)).SendResponse(&w)
}

// swagger:route GET /Controller/{id} Controller listSingle
// Return a list of Controller from disk
// responses:
//	200: bookResponse
//	404: errorResponse

// GetWine handles GET requests
func (b *Controller) GetWine(rw http.ResponseWriter, r *http.Request) {
	id := HttpUtil.GetID(r)
	b.l.Println(b.e.SelectException(1007), id)
	db := b.db.Pull()
	wine := Wine.New(db.Db)
	Win, err := wine.Find(id)
	if err != nil {

	}

	switch err {
	case nil:

	case Wine.WineNotFound:
		b.l.Println(b.e.SelectException(1008), err)
		Result.New(0, 0, http.StatusNotFound, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&rw)

		return
	default:
		b.l.Println(b.e.SelectException(1008), err)
		Result.New(0, 0, http.StatusInternalServerError, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&rw)
		return
	}

	b.db.Push(db)
	Result.New(0, 0, http.StatusOK, "Success", JsonUtil.New(nil, nil).Struct2Json(Win)).SendResponse(&rw)

}

// swagger:route DELETE /Controller/{id} Controller deleteBook
// Update a Controller details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  501: errorResponse

// DeleteWine handles DELETE requests and removes items from the database
func (b *Controller) DeleteWine(w http.ResponseWriter, r *http.Request) {
	id := HttpUtil.GetID(r)

	b.l.Println(b.e.SelectException(1014), id)
	db := b.db.Pull()
	wine := Wine.New(db.Db)
	_, err := wine.Delete(id)
	if err != nil {

	}
	if err == Wine.WineNotFound {
		b.l.Println(b.e.SelectException(1015))
		Result.New(0, -1, http.StatusNotFound, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&w)

		return
	}

	if err != nil {
		b.l.Println("[ERROR] deleting record", err)
		Result.New(0, -1, http.StatusInternalServerError, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: err.Error()})).SendResponse(&w)
		return
	}

	b.db.Push(db)
	Result.New(0, 0, http.StatusOK, "Success", "").SendResponse(&w)

}

// swagger:route PUT /Controller Controller updateBook
// Update a Controller details
//
// responses:
//	201: noContentResponse
//  404: errorResponse
//  422: errorValidation

// UpdateWine handles PUT requests to update Controller
func (b *Controller) UpdateWine(w http.ResponseWriter, r *http.Request) {

	wine := r.Context().Value(Key{}).(Wine.Wine)
	b.l.Println(b.e.SelectException(1010), wine.ID)
	db := b.db.Pull()
	win := Wine.New(db.Db)
	_, err := win.Update(0, 0)

	if err == Wine.WineNotFound {
		b.l.Println(b.e.SelectException(1011), err)
		Result.New(0, -1, http.StatusNotFound, "UnSuccess", JsonUtil.New(nil, nil).Struct2Json(&GenericError{Message: "Wine not found"})).SendResponse(&w)
		return
	}

	Result.New(0, 0, http.StatusOK, "Success", "").SendResponse(&w)
}
