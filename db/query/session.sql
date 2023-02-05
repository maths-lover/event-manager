-- name: CreateSession :one
INSERT INTO "session" (title, description, date_time, venue, event_id, speaker_id)
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: ListSessions :many
SELECT * FROM "session" ORDER BY "id" LIMIT $1 OFFSET $2;

-- name: ListSession :one
SELECT * FROM "session" WHERE id = $1 LIMIT 1;

-- name: UpdateSession :one
UPDATE "session" SET title = $1, description = $2, date_time = $3, venue = $4, event_id = $5, speaker_id = $6
WHERE id = $7 RETURNING *;

-- name: DeleteSession :exec
DELETE FROM "session" WHERE id = $1;
