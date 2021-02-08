-- name: CreatePortfolio :one
INSERT INTO portfolio (
  fan_id,
  creator_id,
  stock_id,
  quantity
  ) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPortfolio :one
SELECT * FROM portfolio
WHERE id = $1
LIMIT 1;

-- name: UpdateStockQuantity :exec
UPDATE portfolio
SET quantity = $2
WHERE id = $1;

-- name: DeleteStockFromPortfolio :exec
DELETE FROM portfolio
WHERE id = $1;
