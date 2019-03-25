package api

import (
	"log"
	"net/http"
	"time"

	"github.com/justinas/alice"

	"github.com/gorilla/mux"
)

// todo: Comment
type Api struct {
	Router *mux.Router
}

// todo: Comment
func (a *Api) Initialise() {
	a.Router = mux.NewRouter()
}

// todo: Comment
func (a *Api) Run(addr string) {
	log.Fatal(http.ListenAndServe(addr, alice.New(loggerHandler, recoverHandler).Then(a.Router)))
}

func loggerHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		h.ServeHTTP(w, r)
		log.Printf("<< %s %s %v", r.Method, r.URL.Path, time.Since(start))
	})
}

func recoverHandler(next http.Handler) http.Handler {
	fn := func(w http.ResponseWriter, r *http.Request) {
		defer func() {
			if err := recover(); err != nil {
				log.Printf("panic: %+v", err)
				http.Error(w, http.StatusText(500), 500)
			}
		}()
		next.ServeHTTP(w, r)
	}
	return http.HandlerFunc(fn)
}
