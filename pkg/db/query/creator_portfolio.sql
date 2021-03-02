-- name: CreateCreatorPortfolio :one
INSERT INTO creator_portfolio (creator_id, stock_id, quantity)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetPortfolioByCreatorID :one
SELECT *
FROM creator_portfolio
WHERE creator_id = $1;
-- name: UpdateCreatorStockQuantity :exec
UPDATE creator_portfolio
SET quantity = $3
WHERE creator_id = $1
  and stock_id = $2;
-- name: DeleteStockFromCreatorPortfolio :exec
DELETE FROM creator_portfolio
WHERE stock_id = $2
  and creator_id = $1;