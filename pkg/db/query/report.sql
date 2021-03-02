-- name: CreateReport :one
INSERT INTO report (
  trading_date,
  stock_id,
  currency_id,
  first_price,
  last_price,
  min_price,
  max_price,
  avg_price,
  total_amount,
  volume
  ) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8, $9, $10
) RETURNING *;

-- name: GetReport :one
SELECT * FROM report
WHERE id = $1
LIMIT 1;
