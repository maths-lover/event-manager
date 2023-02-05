-- name: CreateOrganizer :one
INSERT INTO "organizer" (
  "name",
  "email",
  "phone",
  "company",
  "logo",
  "address"
)
VALUES (
  $1,
  $2,
  $3,
  $4,
  $5,
  $6
) RETURNING *;

-- name: ListOrganizer :one
SELECT *
FROM "organizer"
WHERE "id" = $1;

-- name: ListOrganizers :many
SELECT *
FROM "organizer"
ORDER BY "id"
LIMIT $1
OFFSET $2;

-- name: UpdateOrganizer :one
UPDATE "organizer"
SET
  "name" =  $2,
  "email" = $3,
  "phone" = $4,
  "company" = $5,
  "logo" = $6,
  "address" = $7
WHERE "id" = $1
RETURNING *;

-- name: DeleteOrganizer :exec
DELETE FROM "organizer"
WHERE "id" = $1;
