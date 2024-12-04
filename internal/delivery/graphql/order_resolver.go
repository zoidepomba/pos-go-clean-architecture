package graphql

import (
	"context"
	"project/internal/service"
	"time"
)

type Order struct {
	ID           int
	CustomerName string
	TotalAmount  float64
	CreatedAt    time.Time
}

type OrderResolver struct {
	Service *service.OrderService
}

func (r *OrderResolver) ListOrders(ctx context.Context) ([]*Order, error) {
	orders, err := r.Service.ListOrders()
	if err != nil {
		return nil, err
	}

	var gqlOrders []*Order
	for _, order := range orders {
		gqlOrders = append(gqlOrders, &Order{
			ID:           order.ID,
			CustomerName: order.CustomerName,
			TotalAmount:  order.TotalAmount,
			CreatedAt:    order.CreatedAt,
		})
	}

	return gqlOrders, nil
}
