package service

import (
	"project/internal/models"
	"project/internal/repository"
	"time"
)

type OrderService struct {
	Repo *repository.OrderRepository
}

func (s *OrderService) ListOrders() ([]models.Order, error) {
	return s.Repo.ListOrders()
}

func (s *OrderService) CreateOrder(order models.Order) error {
	// Adiciona a data de criação
	order.CreatedAt = time.Now()
	return s.Repo.CreateOrder(order)
}
