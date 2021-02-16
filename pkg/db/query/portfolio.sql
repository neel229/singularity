-- name: CreatePortfolio :one
INSERT INTO portfolio (
  trader_id,
  stock_id,
  quantity
  ) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetPortfolio :one
SELECT * FROM portfolio
WHERE id = $1
LIMIT 1;

-- name: GetPortfolioByTraderID :one
SELECT * FROM portfolio
WHERE trader_id = $1
LIMIT 1;

-- name: UpdateStockQuantity :exec
UPDATE portfolio
SET quantity = $2
WHERE id = $1;

-- name: DeleteStockFromPortfolio :exec
DELETE FROM portfolio
WHERE id = $1;
