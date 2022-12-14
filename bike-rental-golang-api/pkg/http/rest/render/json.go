package render

import (
	"encoding/json"
	"net/http"
)

// ToJson renders a provided value v to a json response and send it.
func ToJson(w http.ResponseWriter, v interface{}) {
	js, err := json.Marshal(v)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}

	w.Header().Set("Content-Type", "application/json")

	_, err = w.Write(js)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)

		return
	}
}
