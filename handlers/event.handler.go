package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func EventHandler(w http.ResponseWriter, r *http.Request, accounts map[string]account) {
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
		destination := body.Destination
		id := accounts[destination].ID
		if id == "" {
			accounts[destination] = account{ID: destination, Balance: body.Amount}
			w.WriteHeader(http.StatusCreated)
			response = fmt.Sprintf("%d {'destination': id:'%s', balance:%d}", http.StatusCreated, accounts[destination].ID, accounts[destination].Balance)
			json.NewEncoder(w).Encode(response)
			return
		}

		accounts[destination] = account{ID: destination, Balance: accounts[destination].Balance + body.Amount}

		response = fmt.Sprintf("%d {'destination': id:%s, balance:%d}", http.StatusCreated, accounts[destination].ID, accounts[destination].Balance)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	case "withdraw":
		origin := body.Origin
		id := accounts[origin].ID
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			response = fmt.Sprintf("%d 0", http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		accounts[origin] = account{ID: origin, Balance: accounts[origin].Balance - body.Amount}

		response = fmt.Sprintf("%d {'origin': id:%s, balance:%d}", http.StatusCreated, accounts[origin].ID, accounts[origin].Balance)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	case "transfer":
		origin := body.Origin
		destination := body.Destination
		id := accounts[origin].ID
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			response = fmt.Sprintf("%d 0", http.StatusNotFound)
			json.NewEncoder(w).Encode(response)
			return
		}

		accounts[destination] = account{ID: destination, Balance: 0}

		accounts[origin] = account{ID: origin, Balance: accounts[origin].Balance - body.Amount}
		accounts[destination] = account{ID: destination, Balance: accounts[destination].Balance + body.Amount}

		response = fmt.Sprintf("%d {'origin': id:'%s', balance:'%d'}, 'destination': {id:'%s', balance:'%d'}", http.StatusCreated, accounts[origin].ID, accounts[origin].Balance, accounts[destination].ID, accounts[destination].Balance)

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}
