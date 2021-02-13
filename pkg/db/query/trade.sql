-- name: CreateTrade :one
INSERT INTO trade (
    stock_id,
    buyer_id,
    seller_id,
    quantity,
    unit_price,
    details,
    offer_id
  )
VALUES ($1, $2, $3, $4, $5, $6, $7)
RETURNING *;
-- name: GetTrade :one
SELECT *
FROM trade
WHERE id = $1
LIMIT 1;
-- name: GetTradesByBuyer :many
SELECT *
FROM trade
WHERE buyer_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;
-- name: GetTradesBySeller :many
SELECT *
FROM trade
WHERE seller_id = $1
ORDER BY id
LIMIT $2 OFFSET $3;