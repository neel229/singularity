-- name: CreateVirginTrade :one
INSERT INTO virgin_trade (
  stock_id,
  creator_id,
  buyer_id,
  quantity,
  unit_price,
  details,
  virgin_offer_id
  ) VALUES (
  $1, $2, $3, $4, $5, $6, $7
) RETURNING *;

-- name: GetVirginTrade :one
SELECT * FROM virgin_trade
WHERE id = $1
LIMIT 1;

-- name: ListVirginTradesByCreator :many
SELECT * FROM virgin_trade
WHERE creator_id = $1
LIMIT $2
OFFSET $3;

-- name: ListVirginTradesByFan :many
SELECT * FROM virgin_trade
WHERE buyer_id = $1
LIMIT $2
OFFSET $3;
