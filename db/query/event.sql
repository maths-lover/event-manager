-- name: CreateEvent :one
INSERT INTO event (
  title,
  description,
  venue,
  type,
  organizer_id
) VALUES (
  $1, $2, $3, $4, $5
) RETURNING *;

-- name: ListEvent :one
SELECT * FROM event
WHERE id = $1 LIMIT 1;

-- name: ListEvents :many
SELECT * FROM event
ORDER BY id
LIMIT $1
OFFSET $2;

-- name: UpdateEvent :one
UPDATE "event" SET
"title" = $2,
"description" = $3,
"date_time" = $4,
"venue" = $5,
"type" = $6,
"organizer_id" = $7
WHERE "id" = $1
RETURNING *;

-- name: DeleteEvent :exec
DELETE FROM "event" WHERE "id" = $1;
