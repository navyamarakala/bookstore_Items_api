package controller

import (
	"net/http"

	"github.com/bookstore_Items_api/http_Utils"
	"github.com/bookstore_Items_api/model"
	"github.com/bookstore_Items_api/service"
)

func CreateItems(w http.ResponseWriter, r *http.Request) {
	var item model.Item
	data, err := service.CreateItem(item)
	if err != nil {
		http_Utils.RespondError(w, *err)
	}

	http_Utils.RespondJSON(w, http.StatusCreated, data)
}

func GetItems(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("pong"))
}
