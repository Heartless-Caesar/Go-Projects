-- name: CreateTransfers :one
INSERT INTO transfers (
  account_from_id, account_to_id, amount
) VALUES (
  $1, $2, $3
)
RETURNING *;

-- name: GetTransfers :one
SELECT * FROM transfers
WHERE id = $1 LIMIT 1;

-- name: ListTransfers :many
SELECT * FROM transfers
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateTransfers :one
UPDATE transfers
set amount = $2
WHERE id = $1
RETURNING *;

-- name: DeleteTransfers :exec
DELETE FROM transfers
WHERE id = $1;