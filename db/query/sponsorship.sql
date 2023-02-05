-- name: CreateSponsorship :one
INSERT INTO "sponsorship" ("event_id", "company_name", "contact_person", "email", "phone", "package")
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: ListSponsorship :one
SELECT * FROM "sponsorship" WHERE "id" = $1 LIMIT 1;

-- name: ListSponsorships :many
SELECT * FROM "sponsorship" ORDER BY "id" LIMIT $1 OFFSET $2;

-- name: UpdateSponsorship :one
UPDATE "sponsorship"
SET "event_id" = $2, "company_name" = $3, "contact_person" = $4, "email" = $5, "phone" = $6, "package" = $7
WHERE "id" = $1 RETURNING *;

-- name: DeleteSponsorship :exec
DELETE FROM "sponsorship" WHERE "id" = $1;
