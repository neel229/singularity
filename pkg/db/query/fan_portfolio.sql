-- name: CreateFanPortfolio :one
INSERT INTO fan_portfolio (fan_id, stock_id, quantity)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPortfolioByFanID :one
SELECT *
FROM fan_portfolio
WHERE fan_id = $1;
-- name: UpdateFanStockQuantity :exec
UPDATE fan_portfolio
SET quantity = $3
WHERE fan_id = $1
  and stock_id = $2;
-- name: DeleteStockFromFanPortfolio :exec
DELETE FROM fan_portfolio
WHERE stock_id = $2
  and fan_id = $1;