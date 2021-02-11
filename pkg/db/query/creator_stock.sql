-- name: CreateCreatorStock :one
INSERT INTO creator_stock (creator_id, stock_id)
VALUES ($1, $2)
RETURNING *;
-- name: GetCreatorStock :one
SELECT *
FROM creator_stock
WHERE id = $1
LIMIT 1;