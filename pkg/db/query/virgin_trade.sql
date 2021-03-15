-- name: CreateVirginTrade :one
INSERT INTO virgin_trade (
    stock_id,
    creator_id,
    fan_id,
    quantity,
    unit_price,
    details
  )
VALUES ($1, $2, $3, $4, $5, $6)
RETURNING *;
-- name: GetVirginTrade :one
SELECT *
FROM virgin_trade
WHERE id = $1
LIMIT 1;
-- name: ListVirginTradesByCreator :many
SELECT *
FROM virgin_trade
WHERE creator_id = $1
LIMIT $2 OFFSET $3;
-- name: ListVirginTradesByFan :many
SELECT *
FROM virgin_trade
WHERE fan_id = $1
LIMIT $2 OFFSET $3;