package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EventHandler(w http.ResponseWriter, r *http.Request, accounts map[int]account) {
	var body body
	var response string

	err := json.NewDecoder(r.Body).Decode(&body)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	bodyType := body.Type

	switch bodyType {
	case "deposit":
		id := accounts[body.Destination].ID
		if id == 0 {
			accounts[body.Destination] = account{ID: body.Destination, Balance: 10}
			w.WriteHeader(http.StatusCreated)

			response = fmt.Sprintf("%d {'destination': id:%d, balance:%d}", http.StatusCreated, accounts[body.Destination].ID, accounts[body.Destination].Balance)
			json.NewEncoder(w).Encode(response)
			return
		}

		accounts[body.Destination] = account{ID: body.Destination, Balance: accounts[body.Destination].Balance + body.Amount}

		fmt.Printf("accountDestination é %+v\n", accounts[body.Destination])

		response = fmt.Sprintf("%d {'destination': id:%d, balance:%d}", http.StatusCreated, accounts[body.Destination].ID, accounts[body.Destination].Balance)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)
	}
}
