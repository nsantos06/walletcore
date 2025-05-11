package web

import (
	"encoding/json"
	"net/http"

	createclient "github.com.br/devfullcycle/fc-ms-wallet/internal/usecase/create_client"
)

type WebClientHandler struct {
	createClientUseCase createclient.CreateClientUseCase
}

func NewWebClientHandler(createClientUseCase createclient.CreateClientUseCase) *WebClientHandler {
	return &WebClientHandler{
		createClientUseCase: createClientUseCase,
	}
}
func (h *WebClientHandler) CreateClient(w http.ResponseWriter, r *http.Request)  {
	var dto createclient.CreateClientInputDTO
	err := json.NewDecoder(r.Body).Decode(&dto)
	if err != nil {
		
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	output, err := h.createClientUseCase.Execute(dto)
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