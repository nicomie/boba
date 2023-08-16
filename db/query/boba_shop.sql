-- name: CreateBobaShop :one
INSERT INTO boba_shop (
    name
) VALUES (
    $1
) RETURNING *;

