-- name: CreateOrder :one
INSERT INTO orders (
    user_id 
) VALUES (
    $1
) RETURNING *;

-- name: GetOrder :one
SELECT * FROM orders 
WHERE order_id = $1 LIMIT 1;

-- name: ListOrders :many
SELECT * FROM orders
ORDER BY order_id
LIMIT $1
OFFSET $2;

-- name: UpdateOrder :exec
UPDATE orders
SET status = $1
WHERE order_id = $2
AND sqlc.arg(current_items)::int > 0;