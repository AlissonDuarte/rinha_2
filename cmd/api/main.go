package main

import (
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github/alisson/chi/internal/infrastructure/database"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	database.CreateDb()
	r.Get("/clientes/{pk}/extrato", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// realizar uma consulta no banco e retornar o extrato do cliente

	})

	r.Post("/clientes/{pk}/transacoes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(`{"message": "Transação criada para o cliente"}`))
	})

	http.ListenAndServe(":3000", r)
}
