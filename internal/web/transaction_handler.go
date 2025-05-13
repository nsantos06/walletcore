package web

import (
	"encoding/json"
	"net/http"

	createtransaction "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_transaction"
)

type WeTransactionHandler struct {
	CreateTransactionUseCase createtransaction.CreateTransactionUseCase
}
	
func NewWebTransactionHandler(createTransactionUseCase createtransaction.CreateTransactionUseCase) *WeTransactionHandler {
	return &WeTransactionHandler{
		CreateTransactionUseCase: createTransactionUseCase,
	}
}
func (h *WeTransactionHandler) CreateTransaction(w http.ResponseWriter, r *http.Request) {	
	var dto createtransaction.CreateTransactionInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON: " + err.Error(), http.StatusBadRequest)
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}
	//Use Case
	ctx := r.Context()
	output, err := h.CreateTransactionUseCase.Execute(ctx, dto)
	if err != nil {
		http.Error(w, "Erro ao criar transação: " + err.Error(), http.StatusInternalServerError)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	// Retornar Json

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)

		if err := json.NewEncoder(w).Encode(output); err != nil {
		http.Error(w, "Erro ao codificar resposta: " + err.Error(), http.StatusInternalServerError)
	}
	
}