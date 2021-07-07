package functions

import (
	"encoding/json"
	"net/http"
)

type printableError struct {
	Message string `json:"error"`
}

func SendError(err string, code int, w http.ResponseWriter) {
	w.Header().Add("Content-Type", "application/json")

	w.WriteHeader(code)

	pe := printableError{err}
	data, _ := json.Marshal(pe)
	_, _ = w.Write(data)
}
