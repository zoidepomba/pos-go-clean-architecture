package service

import (
    "project/internal/models"
    "project/internal/repository"
)

type OrderService struct {
    Repo *repository.OrderRepository
}

func (s *OrderService) ListOrders() ([]models.Order, error) {
    return s.Repo.ListOrders()
}