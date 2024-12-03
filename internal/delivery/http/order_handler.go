package http

import (
    "encoding/json"
    "net/http"
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