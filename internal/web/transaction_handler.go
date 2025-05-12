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
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	ctx := r.Context()
	output, err := h.CreateTransactionUseCase.Execute(ctx, dto)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err = json.NewEncoder(w).Encode(output)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}