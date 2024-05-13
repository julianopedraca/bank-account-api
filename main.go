package main

import (
	"fmt"
	"net/http"

	handlers "github.com/julianopedraca/account-operations/handlers"
	structs "github.com/julianopedraca/account-operations/structs"
)

type account = structs.Account

var accounts map[int]account = make(map[int]structs.Account)

func main() {

	router := http.NewServeMux()
	// router.HandleFunc("GET /", func(w http.ResponseWriter, r *http.Request) { handlers.AllHandler(w, r, accounts) })
	// router.HandleFunc("POST /reset", func(w http.ResponseWriter, r *http.Request) { handlers.ResetHandler(w, r, accounts) })
	router.HandleFunc("POST /event", func(w http.ResponseWriter, r *http.Request) { handlers.EventHandler(w, r, accounts) })
	router.HandleFunc("GET /balance", func(w http.ResponseWriter, r *http.Request) { handlers.BalanceHandler(w, r, accounts) })

	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server listening on port :8080")
	server.ListenAndServe()
}
