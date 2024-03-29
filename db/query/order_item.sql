-- name: CreateOrderItem :one
INSERT INTO order_items (
    order_id,
    product_id,
    quantity
) VALUES (
    $1, $2, $3
) RETURNING *;

-- name: GetOrderItem :one
SELECT * FROM order_items 
WHERE order_item_id = $1 LIMIT 1;

-- name: ListOrderItem :many
SELECT * FROM order_items
ORDER BY order_id
LIMIT $1
OFFSET $2;

-- name: GetOrderItemCount :one
SELECT COUNT(*)::int AS COUNT
FROM order_items
WHERE order_id = $1;
