-- name: CreateExhibition :one
INSERT INTO exhibition (event_id, company_name, contact_person, email, phone, booth_number, floor_plan) 
VALUES ($1, $2, $3, $4, $5, $6, $7) RETURNING *;

-- name: ListExhibition :one
SELECT * FROM exhibition WHERE id = $1 LIMIT 1;

-- name: ListExhibitions :many
SELECT * FROM exhibition ORDER BY "id" LIMIT $1 OFFSET $2;

-- name: UpdateExhibition :one
UPDATE exhibition 
SET event_id = $1, company_name = $2, contact_person = $3, email = $4, phone = $5, booth_number = $6, floor_plan = $7 
WHERE id = $8 RETURNING *;

-- name: DeleteExhibition :exec
DELETE FROM exhibition WHERE id = $1;
