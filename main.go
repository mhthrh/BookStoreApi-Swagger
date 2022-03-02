package main

import (
	"context"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/mhthrh/ApiStore/Utility/ConfigUtil"
	"github.com/mhthrh/ApiStore/Utility/DbUtil/DbPool"
	"github.com/mhthrh/ApiStore/Utility/LogUtil"
	"github.com/mhthrh/ApiStore/View"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

func main() {

	cfg := ConfigUtil.ReadConfig("Files/ConfigCoded.json")
	l := LogUtil.New()
	sm := mux.NewRouter()
	db := DbPool.New(&DbPool.DbInfo{
		Host:            "localhost",
		Port:            5432,
		User:            "postgres",
		Pass:            "123456",
		Dbname:          "Curency",
		Driver:          "postgres",
		ConnectionCount: 10,
		RefreshPeriod:   20,
	})

	View.RunApiOnRouter(sm, l, cfg, db)

	server := http.Server{
		Addr:              fmt.Sprintf("%s:%d", cfg.Server.IP, cfg.Server.Port),
		Handler:           sm,
		ErrorLog:          log.New(LogUtil.LogrusErrorWriter{}, "", 0),
		ReadTimeout:       10 * time.Second,
		WriteTimeout:      20 * time.Second,
		IdleTimeout:       180 * time.Second,
		ReadHeaderTimeout: 0,
		MaxHeaderBytes:    0,
		TLSConfig:         nil,
		TLSNextProto:      nil,
		ConnState:         nil,
		BaseContext:       nil,
		ConnContext:       nil,
	}

	go func() {
		l.Println("Starting server on  %s:%d", cfg.Server.IP, cfg.Server.Port)

		err := server.ListenAndServe()
		if err != nil {
			l.Printf("Error starting server: %s\n", err)
			os.Exit(1)
		}
	}()

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	signal.Notify(c, os.Kill)

	log.Println("Got signal:", <-c)

	ctx, _ := context.WithTimeout(context.Background(), 30*time.Second)
	server.Shutdown(ctx)

}
