-- name: CreateCurrency :one
INSERT INTO currency (
  code, 
  name,
  is_base
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetCurrency :one
SELECT * FROM currency
WHERE id = $1 LIMIT 1;

-- name: ListCurrencies :many
SELECT * FROM currency
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCurrency :exec
UPDATE currency
SET is_base = $2
WHERE id = $1;

-- name: DeleteCurrency :exec
DELETE FROM currency
WHERE id = $1;
