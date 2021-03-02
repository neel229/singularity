-- name: CreateVirginOffer :one
INSERT INTO virgin_offer (
    creator_id,
    stock_id,
    quantity,
    price
  )
VALUES ($1, $2, $3, $4)
RETURNING *;
-- name: GetVirginOffer :one
SELECT *
FROM virgin_offer
WHERE id = $1
LIMIT 1;
-- name: GetVirginOfferByCreator :one
SELECT *
FROM virgin_offer
WHERE creator_id = $1
LIMIT 1;
-- name: ListVirginOffers :many
SELECT *
FROM virgin_offer
ORDER BY id
LIMIT $1 OFFSET $2;