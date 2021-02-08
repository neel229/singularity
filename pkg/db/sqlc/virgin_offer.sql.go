// Code generated by sqlc. DO NOT EDIT.
// source: virgin_offer.sql

package db

import (
	"context"
	"database/sql"
)

const createVirginOffer = `-- name: CreateVirginOffer :one
INSERT INTO virgin_offer (
  creator_id,
  stock_id,
  quantity,
  price
  ) VALUES (
  $1, $2, $3, $4
) RETURNING id, creator_id, stock_id, quantity, price, ts
`

type CreateVirginOfferParams struct {
	CreatorID sql.NullInt64  `json:"creator_id"`
	StockID   sql.NullInt64  `json:"stock_id"`
	Quantity  sql.NullString `json:"quantity"`
	Price     sql.NullString `json:"price"`
}

func (q *Queries) CreateVirginOffer(ctx context.Context, arg CreateVirginOfferParams) (VirginOffer, error) {
	row := q.db.QueryRowContext(ctx, createVirginOffer,
		arg.CreatorID,
		arg.StockID,
		arg.Quantity,
		arg.Price,
	)
	var i VirginOffer
	err := row.Scan(
		&i.ID,
		&i.CreatorID,
		&i.StockID,
		&i.Quantity,
		&i.Price,
		&i.Ts,
	)
	return i, err
}

const getVirginOffer = `-- name: GetVirginOffer :one
SELECT id, creator_id, stock_id, quantity, price, ts FROM virgin_offer
WHERE id = $1
LIMIT 1
`

func (q *Queries) GetVirginOffer(ctx context.Context, id int64) (VirginOffer, error) {
	row := q.db.QueryRowContext(ctx, getVirginOffer, id)
	var i VirginOffer
	err := row.Scan(
		&i.ID,
		&i.CreatorID,
		&i.StockID,
		&i.Quantity,
		&i.Price,
		&i.Ts,
	)
	return i, err
}

const listVirginOffers = `-- name: ListVirginOffers :many
SELECT id, creator_id, stock_id, quantity, price, ts FROM virgin_offer
WHERE creator_id = $1
ORDER BY id
LIMIT $2
OFFSET $3
`

type ListVirginOffersParams struct {
	CreatorID sql.NullInt64 `json:"creator_id"`
	Limit     int32         `json:"limit"`
	Offset    int32         `json:"offset"`
}

func (q *Queries) ListVirginOffers(ctx context.Context, arg ListVirginOffersParams) ([]VirginOffer, error) {
	rows, err := q.db.QueryContext(ctx, listVirginOffers, arg.CreatorID, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []VirginOffer
	for rows.Next() {
		var i VirginOffer
		if err := rows.Scan(
			&i.ID,
			&i.CreatorID,
			&i.StockID,
			&i.Quantity,
			&i.Price,
			&i.Ts,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}