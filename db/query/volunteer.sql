-- name: CreateVolunteer :one
INSERT INTO "volunteer" (name, email, phone, event_id, role, availability) 
VALUES ($1, $2, $3, $4, $5, $6) RETURNING *;

-- name: ListVolunteer :one
SELECT * FROM "volunteer"
WHERE "id" = $1 LIMIT 1;

-- name: ListVolunteers :many
SELECT * FROM "volunteer" ORDER BY "id" LIMIT $1 OFFSET $2;

-- name: UpdateVolunteer :one
UPDATE "volunteer" 
SET name=$1, email=$2, phone=$3, event_id=$4, role=$5, availability=$6 
WHERE id=$7 RETURNING *;

-- name: DeleteVolunteer :exec
DELETE FROM "volunteer" WHERE id=$1;
