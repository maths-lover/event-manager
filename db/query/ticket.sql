-- name: CreateTicket :one
INSERT INTO ticket (event_id, attendee_id, type, price) 
VALUES ($1, $2, $3, $4) RETURNING *;

-- name: ListTicket :one
SELECT * FROM ticket WHERE id = $1 LIMIT 1;

-- name: ListTickets :many
SELECT * FROM ticket ORDER BY id LIMIT $1 OFFSET $2;

-- name: UpdateTicket :one
UPDATE ticket 
SET event_id = $2, attendee_id = $3, type = $4, price = $5 
WHERE id = $1 RETURNING *;

-- name: DeleteTicket :exec
DELETE FROM ticket WHERE id = $1;
