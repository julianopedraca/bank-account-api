package handlers

import (
	"encoding/json"
	"net/http"
)

func BalanceHandler(w http.ResponseWriter, r *http.Request, accounts map[string]account) {
	query, ok := r.URL.Query()["account_id"]
	if !ok || len(query) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountId := query[0]

	if accounts[accountId].ID == "" {
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(0)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(accounts[accountId].Balance)
}
