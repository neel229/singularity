-- name: CreateOffer :one
INSERT INTO offer (
  trader_id,
  stock_id,
  quantity,
  buy,
  sell,
  price
  ) VALUES (
  $1, $2, $3, $4, $5, $6
) RETURNING *;

-- name: GetOffer :one
SELECT * FROM offer
WHERE id = $1
LIMIT 1;

-- name: ListOffers :many
SELECT * FROM offer
WHERE trader_id = $1
ORDER BY id
LIMIT $2
OFFSET $3;
