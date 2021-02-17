-- name: CreateFanPortfolio :one
INSERT INTO fan_portfolio (
  fan_id,
  stock_id,
  quantity
  ) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetFanPortfolio :one
SELECT * FROM fan_portfolio
WHERE id = $1
LIMIT 1;

-- name: GetPortfolioByFanID :one
SELECT * FROM fan_portfolio
WHERE fan_id = $1
LIMIT 1;

-- name: UpdateFanStockQuantity :exec
UPDATE fan_portfolio
SET quantity = $2
WHERE id = $1;

-- name: DeleteStockFromFanPortfolio :exec
DELETE FROM fan_portfolio
WHERE id = $1;
