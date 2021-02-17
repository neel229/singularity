-- name: CreateCreatorPortfolio :one
INSERT INTO creator_portfolio (
  creator_id,
  stock_id,
  quantity
  ) VALUES (
  $1, $2, $3
) RETURNING *;

-- name: GetCreatorPortfolio :one
SELECT * FROM creator_portfolio
WHERE id = $1
LIMIT 1;

-- name: GetPortfolioByCreatorID :one
SELECT * FROM creator_portfolio
WHERE creator_id = $1
LIMIT 1;

-- name: UpdateCreatorStockQuantity :exec
UPDATE creator_portfolio
SET quantity = $2
WHERE id = $1;

-- name: DeleteStockFromCreatorPortfolio :exec
DELETE FROM creator_portfolio
WHERE id = $1;
