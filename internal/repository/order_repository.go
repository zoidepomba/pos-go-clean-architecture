package repository

import (
    "database/sql"
    "project/internal/models"
)

type OrderRepository struct {
    DB *sql.DB
}

func (r *OrderRepository) ListOrders() ([]models.Order, error) {
    rows, err := r.DB.Query("SELECT id, customer_name, total_amount, created_at FROM orders")
    if err != nil {
        return nil, err
    }
    defer rows.Close()

    var orders []models.Order
    for rows.Next() {
        var order models.Order
        if err := rows.Scan(&order.ID, &order.CustomerName, &order.TotalAmount, &order.CreatedAt); err != nil {
            return nil, err
        }
        orders = append(orders, order)
    }
    return orders, nil
}