package View

import (
	"github.com/go-openapi/runtime/middleware"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/Controller"
	"github.com/mhthrh/ApiStore/Utility/ConfigUtil"
	"github.com/mhthrh/ApiStore/Utility/DbUtil/DbPool"
	"github.com/mhthrh/ApiStore/Utility/ExceptionUtil"
	"github.com/mhthrh/ApiStore/Utility/ValidationUtil"
	"github.com/sirupsen/logrus"
	"net/http"
)

func RunApiOnRouter(sm *mux.Router, l *logrus.Entry, config *ConfigUtil.Config, db *DbPool.DBs) {

	v := ValidationUtil.NewValidation()
	e := ExceptionUtil.New()
	ph := Controller.NewBook(l, v, e, config, db)
	sm.Use(ph.HttpMiddleware)

	getR := sm.Methods(http.MethodGet).Subrouter()
	getR.HandleFunc("/books", ph.BookSList)
	getR.HandleFunc("/wines", ph.WineSList)
	getR.HandleFunc("/books/{id:[0-9]+}", ph.GetBook)
	getR.HandleFunc("/wines/{id:[0-9]+}", ph.GetWine)

	putR := sm.Methods(http.MethodPut).Subrouter()
	putR.HandleFunc("/books", ph.UpdateBook)
	putR.HandleFunc("/wines", ph.UpdateWine)

	postR := sm.Methods(http.MethodPost).Subrouter()
	postR.HandleFunc("/books", ph.AddBook)
	postR.HandleFunc("/wines", ph.AddWine)

	deleteR := sm.Methods(http.MethodDelete).Subrouter()
	deleteR.HandleFunc("/books/{id:[0-9]+}", ph.DeleteBook)
	deleteR.HandleFunc("/wines/{id:[0-9]+}", ph.DeleteWine)

	// handler for documentation
	opts := middleware.RedocOpts{SpecURL: "/swagger.yaml"}
	sh := middleware.Redoc(opts, nil)

	getR.Handle("/docs", sh)
	getR.Handle("/swagger.yaml", http.FileServer(http.Dir("./")))

}
