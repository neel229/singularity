-- name: CreateCreatorStock :one
INSERT INTO creator_stock (creator_id, stock_id)
VALUES ($1, $2)
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