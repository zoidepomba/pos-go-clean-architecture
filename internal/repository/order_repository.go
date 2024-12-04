package repository

import (
	"database/sql"
	"log"
	"project/internal/models"
)

type OrderRepository struct {
	DB *sql.DB
}

func (r *OrderRepository) CreateOrder(order models.Order) error {
	query := "INSERT INTO orders (customer_name, total_amount, created_at) VALUES ($1, $2, $3)"
	_, err := r.DB.Exec(query, order.CustomerName, order.TotalAmount, order.CreatedAt)
	if err != nil {
		log.Printf("Erro ao inserir order: %v", err)
		return err
	}
	return nil
}

func (r *OrderRepository) ListOrders() ([]models.Order, error) {
	rows, err := r.DB.Query("SELECT id, customer_name, total_amount, created_at FROM orders")
	if err != nil {
		log.Printf("Erro ao executar query: %v", err)
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var order models.Order
		if err := rows.Scan(&order.ID, &order.CustomerName, &order.TotalAmount, &order.CreatedAt); err != nil {
			log.Printf("Erro ao escanear resultado: %v", err)
			return nil, err
		}
		orders = append(orders, order)
	}
	return orders, nil
}
