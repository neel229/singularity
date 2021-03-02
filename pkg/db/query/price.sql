-- name: CreatePrice :one
INSERT INTO price (
  stock_id,
  currency_id,
  buy,
  sell
  ) VALUES (
  $1, $2, $3, $4
) RETURNING *;

-- name: GetPrice :one
SELECT * FROM price
WHERE id = $1
LIMIT 1;

-- name: UpdateBuyingPrice :exec
UPDATE price
SET buy = $2
WHERE id = $1;

-- name: UpdateSellingPrice :exec
UPDATE price
SET sell = $2
WHERE id = $1;

-- name: DeletePrice :exec
DELETE FROM price
WHERE id = $1;
