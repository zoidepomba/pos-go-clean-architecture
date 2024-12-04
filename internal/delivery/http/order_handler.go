package http

import (
	"encoding/json"
	"net/http"
	"project/internal/models"
	"project/internal/service"
)

type OrderHandler struct {
    Service *service.OrderService
}

func (h *OrderHandler) ListOrders(w http.ResponseWriter, r *http.Request) {
    orders, err := h.Service.ListOrders()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return
    }
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(orders)
}

func (h *OrderHandler) CreateOrder(w http.ResponseWriter, r *http.Request) {
    var order models.Order
    
    // Decodifica o corpo da requisição
    if err := json.NewDecoder(r.Body).Decode(&order); err != nil {
        http.Error(w, "Erro ao decodificar o corpo da requisição", http.StatusBadRequest)
        return
    }
    
    // Chama o serviço para criar a order
    if err := h.Service.CreateOrder(order); err != nil {
        http.Error(w, "Erro ao criar order", http.StatusInternalServerError)
        return
    }
    
    // Retorna sucesso
    w.WriteHeader(http.StatusCreated)
    json.NewEncoder(w).Encode(map[string]string{"message": "Order criada com sucesso"})
    }