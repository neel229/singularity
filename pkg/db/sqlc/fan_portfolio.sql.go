// Code generated by sqlc. DO NOT EDIT.
// source: fan_portfolio.sql

package db

import (
	"context"
)

const createFanPortfolio = `-- name: CreateFanPortfolio :one
INSERT INTO fan_portfolio (fan_id, stock_id, quantity)
VALUES ($1, $2, $3)
RETURNING id, fan_id, stock_id, quantity
`

type CreateFanPortfolioParams struct {
	FanID    int64  `json:"fan_id"`
	StockID  int64  `json:"stock_id"`
	Quantity string `json:"quantity"`
}

func (q *Queries) CreateFanPortfolio(ctx context.Context, arg CreateFanPortfolioParams) (FanPortfolio, error) {
	row := q.db.QueryRowContext(ctx, createFanPortfolio, arg.FanID, arg.StockID, arg.Quantity)
	var i FanPortfolio
	err := row.Scan(
		&i.ID,
		&i.FanID,
		&i.StockID,
		&i.Quantity,
	)
	return i, err
}

const deleteStockFromFanPortfolio = `-- name: DeleteStockFromFanPortfolio :exec
DELETE FROM fan_portfolio
WHERE stock_id = $2
  and fan_id = $1
`

type DeleteStockFromFanPortfolioParams struct {
	FanID   int64 `json:"fan_id"`
	StockID int64 `json:"stock_id"`
}

func (q *Queries) DeleteStockFromFanPortfolio(ctx context.Context, arg DeleteStockFromFanPortfolioParams) error {
	_, err := q.db.ExecContext(ctx, deleteStockFromFanPortfolio, arg.FanID, arg.StockID)
	return err
}

const getPortfolioByFanID = `-- name: GetPortfolioByFanID :one
SELECT id, fan_id, stock_id, quantity
FROM fan_portfolio
WHERE fan_id = $1
`

func (q *Queries) GetPortfolioByFanID(ctx context.Context, fanID int64) (FanPortfolio, error) {
	row := q.db.QueryRowContext(ctx, getPortfolioByFanID, fanID)
	var i FanPortfolio
	err := row.Scan(
		&i.ID,
		&i.FanID,
		&i.StockID,
		&i.Quantity,
	)
	return i, err
}

const updateFanStockQuantity = `-- name: UpdateFanStockQuantity :exec
UPDATE fan_portfolio
SET quantity = $3
WHERE fan_id = $1
  and stock_id = $2
`

type UpdateFanStockQuantityParams struct {
	FanID    int64  `json:"fan_id"`
	StockID  int64  `json:"stock_id"`
	Quantity string `json:"quantity"`
}

func (q *Queries) UpdateFanStockQuantity(ctx context.Context, arg UpdateFanStockQuantityParams) error {
	_, err := q.db.ExecContext(ctx, updateFanStockQuantity, arg.FanID, arg.StockID, arg.Quantity)
	return err
}
