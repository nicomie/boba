// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: products.sql

package db

import (
	"context"
)

const createProduct = `-- name: CreateProduct :one
INSERT INTO products (
    name, 
    price
) VALUES (
    $1, $2
) RETURNING product_id, name, price
`

type CreateProductParams struct {
	Name  string `json:"name"`
	Price int32  `json:"price"`
}

func (q *Queries) CreateProduct(ctx context.Context, arg CreateProductParams) (Product, error) {
	row := q.db.QueryRow(ctx, createProduct, arg.Name, arg.Price)
	var i Product
	err := row.Scan(&i.ProductID, &i.Name, &i.Price)
	return i, err
}

const deleteProduct = `-- name: DeleteProduct :exec
DELETE FROM products 
WHERE product_id = $1
`

func (q *Queries) DeleteProduct(ctx context.Context, productID int64) error {
	_, err := q.db.Exec(ctx, deleteProduct, productID)
	return err
}

const getProduct = `-- name: GetProduct :one
SELECT product_id, name, price FROM products 
WHERE product_id = $1 LIMIT 1
`

func (q *Queries) GetProduct(ctx context.Context, productID int64) (Product, error) {
	row := q.db.QueryRow(ctx, getProduct, productID)
	var i Product
	err := row.Scan(&i.ProductID, &i.Name, &i.Price)
	return i, err
}

const listProducts = `-- name: ListProducts :many
SELECT product_id, name, price FROM products
ORDER BY product_id
LIMIT $1
OFFSET $2
`

type ListProductsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListProducts(ctx context.Context, arg ListProductsParams) ([]Product, error) {
	rows, err := q.db.Query(ctx, listProducts, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	items := []Product{}
	for rows.Next() {
		var i Product
		if err := rows.Scan(&i.ProductID, &i.Name, &i.Price); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateProduct = `-- name: UpdateProduct :exec
UPDATE products
SET name = $2, price = $3
WHERE product_id = $1
`

type UpdateProductParams struct {
	ProductID int64  `json:"product_id"`
	Name      string `json:"name"`
	Price     int32  `json:"price"`
}

func (q *Queries) UpdateProduct(ctx context.Context, arg UpdateProductParams) error {
	_, err := q.db.Exec(ctx, updateProduct, arg.ProductID, arg.Name, arg.Price)
	return err
}
