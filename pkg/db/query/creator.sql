-- name: CreateCreator :one
INSERT INTO creator (
  first_name,
  last_name,
  user_name,
  email,
  password,
  preferred_currency_id,
  creator_stock,
  virgin_tokens_left
  ) VALUES (
  $1, $2, $3, $4, $5, $6, $7, $8
) RETURNING *;

-- name: GetCreator :one
SELECT * FROM creator
WHERE id = $1
LIMIT 1;

-- name: GetVirginTokensLeft :one
SELECT virgin_tokens_left FROM creator
WHERE id = $1
LIMIT 1;

-- name: ListCreators :many
SELECT * FROM creator
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateCreatorEmail :exec
UPDATE creator
SET email = $2
WHERE id = $1;

-- name: UpdateCreatorPassword :exec
UPDATE creator
SET password = $2
WHERE id = $1;

-- name: UpdateCreatorPreferredCurrency :exec
UPDATE creator
SET preferred_currency_id = $2
WHERE id = $1;

-- name: UpdateVirginTokensLeft :exec
UPDATE creator
SET virgin_tokens_left = $2
WHERE id = $1;
