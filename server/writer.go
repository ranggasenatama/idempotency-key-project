package main

import (
	"encoding/json"
	"net/http"
)

// Write writes the Response sturct in json format, if it can't marshal the Response sturct then it will write
// http.StatusInternalServerError to the response.
func (resp *Response) Write(w http.ResponseWriter, r *http.Request, status int) (int, error) {
	w.Header().Set("Content-Type", "application/json")
	b, err := json.Marshal(resp)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		writeLen, writeErr := w.Write([]byte(`{"errors":["Internal Server Error"]}`))
		if writeErr != nil {
			return writeLen, writeErr
		}
		return writeLen, err
	}

	w.WriteHeader(status)
	return w.Write(b)
}
