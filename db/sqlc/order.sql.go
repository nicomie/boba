// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: order.sql

package db

import (
	"context"
)

const createOrder = `-- name: CreateOrder :one
INSERT INTO orders (
    user_id, 
    status
) VALUES (
    $1, $2
) RETURNING order_id, user_id, status, created_at
`

type CreateOrderParams struct {
	UserID int64  `json:"user_id"`
	Status string `json:"status"`
}

func (q *Queries) CreateOrder(ctx context.Context, arg CreateOrderParams) (Order, error) {
	row := q.db.QueryRow(ctx, createOrder, arg.UserID, arg.Status)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const getOrder = `-- name: GetOrder :one
SELECT order_id, user_id, status, created_at FROM orders 
WHERE order_id = $1 LIMIT 1
`

func (q *Queries) GetOrder(ctx context.Context, orderID int64) (Order, error) {
	row := q.db.QueryRow(ctx, getOrder, orderID)
	var i Order
	err := row.Scan(
		&i.OrderID,
		&i.UserID,
		&i.Status,
		&i.CreatedAt,
	)
	return i, err
}

const listOrders = `-- name: ListOrders :many
SELECT order_id, user_id, status, created_at FROM orders
ORDER BY order_id
LIMIT $1
OFFSET $2
`

type ListOrdersParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListOrders(ctx context.Context, arg ListOrdersParams) ([]Order, error) {
	rows, err := q.db.Query(ctx, listOrders, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Order{}
	for rows.Next() {
		var i Order
		if err := rows.Scan(
			&i.OrderID,
			&i.UserID,
			&i.Status,
			&i.CreatedAt,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateOrder = `-- name: UpdateOrder :exec
UPDATE orders
SET status = $1
WHERE order_id = $2
AND $3::int > 0
`

type UpdateOrderParams struct {
	Status       string `json:"status"`
	OrderID      int64  `json:"order_id"`
	CurrentItems int32  `json:"current_items"`
}

func (q *Queries) UpdateOrder(ctx context.Context, arg UpdateOrderParams) error {
	_, err := q.db.Exec(ctx, updateOrder, arg.Status, arg.OrderID, arg.CurrentItems)
	return err
}
