package handlers

import (
	"encoding/json"
	"net/http"
)

func EventHandler(w http.ResponseWriter, r *http.Request, accounts map[string]account) {
	var body body

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

			response := struct {
				Destination account `json:"destination"`
			}{
				Destination: accounts[destination],
			}

			json.NewEncoder(w).Encode(response)
			return
		}

		accounts[destination] = account{ID: destination, Balance: accounts[destination].Balance + body.Amount}

		response := struct {
			Destination account `json:"destination"`
		}{
			Destination: accounts[destination],
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	case "withdraw":
		origin := body.Origin
		id := accounts[origin].ID
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(0)
			return
		}

		accounts[origin] = account{ID: origin, Balance: accounts[origin].Balance - body.Amount}

		response := struct {
			Origin account `json:"origin"`
		}{
			Origin: accounts[origin],
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	case "transfer":
		origin := body.Origin
		destination := body.Destination
		id := accounts[origin].ID
		if id == "" {
			w.WriteHeader(http.StatusNotFound)
			json.NewEncoder(w).Encode(0)
			return
		}

		accounts[destination] = account{ID: destination, Balance: 0}

		accounts[origin] = account{ID: origin, Balance: accounts[origin].Balance - body.Amount}
		accounts[destination] = account{ID: destination, Balance: accounts[destination].Balance + body.Amount}

		response := struct {
			Origin      account `json:"origin"`
			Destination account `json:"destination"`
		}{
			Origin:      accounts[origin],
			Destination: accounts[destination],
		}

		w.WriteHeader(http.StatusCreated)
		json.NewEncoder(w).Encode(response)

	}
}
