package View

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/Controller"
	"github.com/mhthrh/ApiStore/Utility/ExceptionUtil"
	"github.com/mhthrh/ApiStore/Utility/ValidationUtil"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunApiOnRouter(sm *mux.Router, l *logrus.Entry) {

	v := ValidationUtil.NewValidation()
	e := ExceptionUtil.New()
	ph := Controller.NewBooks(l, v, e)

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/books", ph.ListAll)
	getR.HandleFunc("/books/{id:[0-9]+}", ph.ListSingle)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/books", ph.Update)
	putR.Use(ph.HttpMiddleware)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/books", ph.Create)
	postR.Use(ph.HttpMiddleware)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/books/{id:[0-9]+}", ph.Delete)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

}
