package View

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/Controller"
	"github.com/mhthrh/ApiStore/Model/Result"
	"net/http"
)

func RunApi(add string) {
	fmt.Println("initial server on: ", add)
	router := mux.NewRouter()
	RunApiOnRouter(router)
	http.ListenAndServe(add, router)
}

func RunApiOnRouter(r *mux.Router) {

	sub := r.PathPrefix("/api/v1/book").Subrouter()
	sub.Methods("GET").Path("/books/{range}").HandlerFunc(Controller.AllBooks)
	sub.Methods("GET").Path("/findbook/{isbn}").HandlerFunc(Controller.Book)
	sub.Methods("POST").Path("/add").HandlerFunc(Controller.AddBook)
	sub.Methods("GET").Path("/delete/{isbn}").HandlerFunc(Controller.DeleteBook)
	sub.Methods("POST").Path("/update").HandlerFunc(Controller.UpdateBook)

	r.NotFoundHandler = http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			Result.New(1, 1, 99, "Not found!", nil).Create(&w, http.StatusNotFound)
			return
		})
}
