package View

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/Controller"
	"github.com/mhthrh/ApiStore/Helper/Validation"
	"log"
	"net/http"
)

func RunApiOnRouter(sm *mux.Router, l *log.Logger) {

	v := Validation.NewValidation()
	// create the handlers
	ph := Controller.NewBooks(l, v)

	// handlers for API
	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/books", ph.ListAll)
	getR.HandleFunc("/books/{id:[0-9]+}", ph.ListSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/books", ph.Update)
	putR.Use(ph.MiddlewareValidateBook)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/books", ph.Create)
	postR.Use(ph.MiddlewareValidateBook)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/books/{id:[0-9]+}", ph.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

}
