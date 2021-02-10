-- name: CreateCurrencyRate :one
INSERT INTO currency_rate (currency_id, base_currency_id, rate)
VALUES ($1, $2, $3)
RETURNING *;
-- name: GetCurrencyRate :one
SELECT *
FROM currency_rate
WHERE id = $1
LIMIT 1;
-- name: UpdateCurrencyRate :exec
UPDATE currency_rate
SET rate = $2
WHERE id = $1
RETURNING *;