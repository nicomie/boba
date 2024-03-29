// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.19.1
// source: personnel.sql

package db

import (
	"context"
)

const createPersonnel = `-- name: CreatePersonnel :one
INSERT INTO personnel (
    name, 
    shop,
    hashed_pin
) VALUES (
    $1, $2, $3
) RETURNING badge_number, name, shop, hashed_pin, email, created_at
`

type CreatePersonnelParams struct {
	Name      string `json:"name"`
	Shop      int64  `json:"shop"`
	HashedPin string `json:"hashed_pin"`
}

func (q *Queries) CreatePersonnel(ctx context.Context, arg CreatePersonnelParams) (Personnel, error) {
	row := q.db.QueryRow(ctx, createPersonnel, arg.Name, arg.Shop, arg.HashedPin)
	var i Personnel
	err := row.Scan(
		&i.BadgeNumber,
		&i.Name,
		&i.Shop,
		&i.HashedPin,
		&i.Email,
		&i.CreatedAt,
	)
	return i, err
}
