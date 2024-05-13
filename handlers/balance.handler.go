package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func BalanceHandler(w http.ResponseWriter, r *http.Request, accounts map[string]account) {
	var response string
	query, ok := r.URL.Query()["account_id"]
	if !ok || len(query) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountId := query[0]

	if accounts[accountId].ID == "" {
		w.WriteHeader(http.StatusNotFound)
		response = fmt.Sprintf("%d 0", http.StatusNotFound)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response = fmt.Sprintf("%d %d", http.StatusOK, accounts[accountId].Balance)
	json.NewEncoder(w).Encode(response)
}
