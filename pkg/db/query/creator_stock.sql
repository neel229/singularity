-- name: CreateCreatorStock :one
INSERT INTO creator_stock (creator_id, stock_id, mint_price, current_price)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetCreatorStock :one
SELECT *
FROM creator_stock
WHERE id = $1
LIMIT 1;
-- name: ListCreatorStocks :many
SELECT *
FROM creator_stock
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateStockPrice :exec
UPDATE creator_stock
SET current_price = $2
WHERE creator_id = $1;