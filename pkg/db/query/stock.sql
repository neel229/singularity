-- name: CreateStock :one
INSERT INTO stock (ticker, details, mint_price, current_price)
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetStock :one
SELECT *
FROM stock
WHERE id = $1
LIMIT 1;
-- name: ListStocks :many
SELECT *
FROM stock
ORDER BY id
LIMIT $1 OFFSET $2;
-- name: UpdateStock :exec
UPDATE stock
SET details = $2
WHERE id = $1;
-- name: UpdateStockPrice :exec
UPDATE stock
SET current_price = $2
WHERE ticker = $1;
-- name: DeleteStock :exec
DELETE FROM stock
WHERE id = $1;