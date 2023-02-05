-- name: CreateAttendee :one
INSERT INTO attendee (name, email, phone, event_id, ticket_number, payment_status, payment_date) 
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: ListAttendee :one
SELECT * FROM attendee WHERE id = $1 LIMIT 1;

-- name: ListAttendees :many
SELECT * FROM attendee ORDER BY "id" LIMIT $1 OFFSET $2;

-- name: UpdateAttendee :one
UPDATE attendee 
SET name = $1, email = $2, phone = $3, event_id = $4, ticket_number = $5, payment_status = $6, payment_date = $7
WHERE id = $8 RETURNING *;

-- name: DeleteAttendee :exec
DELETE FROM attendee WHERE id = $1;
