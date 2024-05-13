package handlers

import (
	"net/http"

	"golang.org/x/exp/maps"
)

func ResetHandler(w http.ResponseWriter, r *http.Request, accounts map[string]account) {
	maps.Clear(accounts)
	w.WriteHeader(http.StatusOK)
	_, err := w.Write([]byte("OK"))
	if err != nil {
		// Handle error if unable to write response
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
