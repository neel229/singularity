-- name: CreateFan :one
INSERT INTO fan (
    first_name,
    last_name,
    user_name,
    password,
    email,
    preferred_currency_id
  )
VALUES (
    $1,
    $2,
    $3,
    $4,
    $5,
    $6
  )
RETURNING *;
-- name: GetFan :one
SELECT *
FROM fan
WHERE id = $1
LIMIT 1;
-- name: UpdateEmail :exec
UPDATE fan
SET email = $2
WHERE id = $1;
-- name: UpdatePassword :exec
UPDATE fan
SET password = $2
WHERE id = $1;
-- name: UpdatePreferredCurrency :exec
UPDATE fan
SET preferred_currency_id = $2
WHERE id = $1;