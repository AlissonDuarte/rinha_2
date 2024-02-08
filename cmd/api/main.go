package main

import (
	"fmt"
	"io"
	"encoding/json"
	"net/http"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github/alisson/chi/internal/infrastructure/database"
)

type Transaction struct {
	Valor int `json:"valor"`
	Tipo string `json:"tipo"`
	Descricao string `json:"descricao"`
}

func transactionStruct(jsonData []byte) (Transaction, error) {

	var transaction Transaction

	err := json.Unmarshal(jsonData, &transaction)

	if err != nil{
		return transaction, nil
	}

	return transaction, nil
}

func main() {
	con, err := database.StartConnection()
	
	if err != nil{
		panic("Error to connect to database")
	}

	database.Migrate(con)

	
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	
	r.Get("/clientes/{pk}/extrato", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		// realizar uma consulta no banco e retornar o extrato do cliente

	})

	r.Post("/clientes/{pk}/transacoes", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		defer r.Body.Close()
		defer w.Write([]byte(`{"message": "Transação criada para o cliente"}`))


		bodyBrute, err := io.ReadAll(r.Body)

		if err != nil{
			fmt.Println("ERROR", err)
		}

		data, err := transactionStruct(bodyBrute)
		if err != nil{
			fmt.Println("ERROR", err)
		}

		fmt.Printf("Transaction deserializada: %+v\n", data)
	})

	http.ListenAndServe(":3000", r)
}
