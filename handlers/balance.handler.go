package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func BalanceHandler(w http.ResponseWriter, r *http.Request, accounts map[int]account) {
	var response string
	query, ok := r.URL.Query()["account_id"]
	if !ok || len(query) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	accountId, err := strconv.Atoi(query[0])
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	if accounts[accountId].ID == 0 {
		w.WriteHeader(http.StatusNotFound)
		response = fmt.Sprintf("%d %d", http.StatusNotFound, accounts[accountId].ID)
		json.NewEncoder(w).Encode(response)
		return
	}

	w.WriteHeader(http.StatusOK)
	response = fmt.Sprintf("%d %d", http.StatusOK, accounts[accountId].Balance)
	json.NewEncoder(w).Encode(response)
}
