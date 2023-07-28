-- name: CreateProduct :one
INSERT INTO products (
    name, 
    price
) VALUES (
    $1, $2
) RETURNING *;

-- name: GetProduct :one
SELECT * FROM products 
WHERE product_id = $1 LIMIT 1;

-- name: ListProducts :many
SELECT * FROM products
ORDER BY product_id
LIMIT $1
OFFSET $2;

-- name: UpdateProduct :exec
UPDATE products
SET name = $2, price = $3
WHERE product_id = $1;

-- name: DeleteProduct :exec
DELETE FROM products 
WHERE product_id = $1;
