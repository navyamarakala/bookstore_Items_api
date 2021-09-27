package app

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var (
	r = mux.NewRouter()
)

func StartApp() {
	MapUrls()

	srv := &http.Server{
		Handler: r,
		Addr:    "127.0.0.1:8082",
		// Good practice: enforce timeouts for servers you create!

	}
	log.Fatal(srv.ListenAndServe())

}
