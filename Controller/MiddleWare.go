package Controller

import (
	"context"
	"errors"
	"fmt"
	Book2 "github.com/mhthrh/ApiStore/Model/Book"
	"github.com/mhthrh/ApiStore/Model/Wine"
	"github.com/mhthrh/ApiStore/Utility/JsonUtil"
	"net/http"
	"time"
)

// HttpMiddleware validates the book in the request and calls next if ok
func (b *Controller) HttpMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		fErr := func(err error, i int, in interface{}) {
			start := time.Now()
			b.l.WithFields(map[string]interface{}{
				"method":     r.Method,
				"path":       r.URL,
				"status":     nil,
				"latency_ns": time.Since(start).Nanoseconds(),
			}).Info("request details")
			w.WriteHeader(i)
			JsonUtil.New(w, r.Body).ToJSON(&in)
		}
		fNext := func(in interface{}) {
			ctx := context.WithValue(r.Context(), KeyBook{}, in)
			r = r.WithContext(ctx)
			next.ServeHTTP(w, r)
		}

		if r.Host != fmt.Sprintf("%s:%d", b.c.Server.IP, b.c.Server.Port) {
			err := errors.New("access denied")
			fErr(err, http.StatusForbidden, GenericError{Message: err.Error()})
			return
		}
		if r.Method == http.MethodGet {
			fNext(nil)
		}

		var intFace interface{}
		switch r.RequestURI {
		case "/Controller":
			intFace = &Book2.Book{}
		case "/wines":
			intFace = &Wine.Wine{}
		}
		err := JsonUtil.New(w, r.Body).FromJSON(&intFace)
		if err != nil {
			fErr(err, http.StatusBadRequest, GenericError{Message: err.Error()})
			return
		}

		errs := b.v.Validate(intFace)
		if len(errs) != 0 {
			fErr(errors.New("validation issue"), http.StatusUnprocessableEntity, ValidationError{Messages: errs.Errors()})
			return
		}
		fNext(intFace)
	})
}
