package app

import (
	"net/http"

	"github.com/bookstore_Items_api/controller"
)

func MapUrls() {
	//create routers here
	//log.Println("inside mapurls")
	r.HandleFunc("/create", controller.CreateItems).Methods(http.MethodPost)
	r.HandleFunc("/get", controller.GetItems).Methods(http.MethodGet)
}
