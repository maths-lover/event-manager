-- name: CreateSpeaker :one
INSERT INTO "speaker" (
"name",
"email",
"phone",
"event_id",
"profile_image",
"bio"
) VALUES (
$1,
$2,
$3,
$4,
$5,
$6
) RETURNING *;

-- name: ListSpeakers :many
SELECT * FROM "speaker" ORDER BY "id" LIMIT $1 OFFSET $2;

-- name: ListSpeaker :one
SELECT * FROM "speaker" WHERE "id" = $1 LIMIT 1;

-- name: UpdateSpeaker :one
UPDATE "speaker"
SET
"name" = $1,
"email" = $2,
"phone" = $3,
"event_id" = $4,
"profile_image" = $5,
"bio" = $6
WHERE "id" = $7 RETURNING *;

-- name: DeleteSpeaker :exec
DELETE FROM "speaker" WHERE "id" = $1;
