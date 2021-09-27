package http_Utils

import (
	"encoding/json"
	"net/http"

	"github.com/navyamarakala/bookstore_users_API/utility/errors"
)

func RespondJSON(w http.ResponseWriter, statusCode int64, body interface{}) {
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(int(statusCode))
	json.NewEncoder(w).Encode(body)
}

func RespondError(w http.ResponseWriter, err errors.RestErrors) {
	RespondJSON(w, err.Status, err)
}
