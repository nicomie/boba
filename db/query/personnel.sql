-- name: CreatePersonnel :one
INSERT INTO personnel (
    name, 
    shop,
    hashed_pin
) VALUES (
    $1, $2, $3
) RETURNING *;

