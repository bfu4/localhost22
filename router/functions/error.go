package functions

import (
	"encoding/json"
	"net/http"
)

type printableError struct {
	Message string `json:"error"`
}

func SendError(err string, code int, w http.ResponseWriter) {
	headers := w.Header()
	headers.Add("Content-Type", "application/json")
	headers.Add("Status", string(rune(code)))

	_ = headers.Write(w)

	pe := printableError{Message: err}
	data, _ := json.Marshal(pe)
	_, _ = w.Write(data)
}
