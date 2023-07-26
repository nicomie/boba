-- name: CreateOrder :one
INSERT INTO orders (
    user_id, 
    status
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders 
WHERE order_id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY order_id
LIMIT $1
OFFSET $2;

