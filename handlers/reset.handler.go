package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"golang.org/x/exp/maps"
)

func ResetHandler(w http.ResponseWriter, r *http.Request, accounts map[string]account) {
	var response string
	maps.Clear(accounts)
	w.WriteHeader(http.StatusOK)
	response = fmt.Sprintf("%d OK", http.StatusOK)
	json.NewEncoder(w).Encode(response)
}
