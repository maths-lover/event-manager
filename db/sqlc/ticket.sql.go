// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0
// source: ticket.sql

package db

import (
	"context"
	"database/sql"
)

const createTicket = `-- name: CreateTicket :one
INSERT INTO ticket (event_id, attendee_id, type, price) 
VALUES ($1, $2, $3, $4) RETURNING id, event_id, attendee_id, type, price
`

type CreateTicketParams struct {
	EventID    sql.NullInt32  `json:"event_id"`
	AttendeeID sql.NullInt32  `json:"attendee_id"`
	Type       sql.NullString `json:"type"`
	Price      string         `json:"price"`
}

func (q *Queries) CreateTicket(ctx context.Context, arg CreateTicketParams) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, createTicket,
		arg.EventID,
		arg.AttendeeID,
		arg.Type,
		arg.Price,
	)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.AttendeeID,
		&i.Type,
		&i.Price,
	)
	return i, err
}

const deleteTicket = `-- name: DeleteTicket :exec
DELETE FROM ticket WHERE id = $1
`

func (q *Queries) DeleteTicket(ctx context.Context, id int32) error {
	_, err := q.db.ExecContext(ctx, deleteTicket, id)
	return err
}

const listTicket = `-- name: ListTicket :one
SELECT id, event_id, attendee_id, type, price FROM ticket WHERE id = $1 LIMIT 1
`

func (q *Queries) ListTicket(ctx context.Context, id int32) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, listTicket, id)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.AttendeeID,
		&i.Type,
		&i.Price,
	)
	return i, err
}

const listTickets = `-- name: ListTickets :many
SELECT id, event_id, attendee_id, type, price FROM ticket ORDER BY id LIMIT $1 OFFSET $2
`

type ListTicketsParams struct {
	Limit  int32 `json:"limit"`
	Offset int32 `json:"offset"`
}

func (q *Queries) ListTickets(ctx context.Context, arg ListTicketsParams) ([]Ticket, error) {
	rows, err := q.db.QueryContext(ctx, listTickets, arg.Limit, arg.Offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var items []Ticket
	for rows.Next() {
		var i Ticket
		if err := rows.Scan(
			&i.ID,
			&i.EventID,
			&i.AttendeeID,
			&i.Type,
			&i.Price,
		); err != nil {
			return nil, err
		}
		items = append(items, i)
	}
	if err := rows.Close(); err != nil {
		return nil, err
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return items, nil
}

const updateTicket = `-- name: UpdateTicket :one
UPDATE ticket 
SET event_id = $2, attendee_id = $3, type = $4, price = $5 
WHERE id = $1 RETURNING id, event_id, attendee_id, type, price
`

type UpdateTicketParams struct {
	ID         int32          `json:"id"`
	EventID    sql.NullInt32  `json:"event_id"`
	AttendeeID sql.NullInt32  `json:"attendee_id"`
	Type       sql.NullString `json:"type"`
	Price      string         `json:"price"`
}

func (q *Queries) UpdateTicket(ctx context.Context, arg UpdateTicketParams) (Ticket, error) {
	row := q.db.QueryRowContext(ctx, updateTicket,
		arg.ID,
		arg.EventID,
		arg.AttendeeID,
		arg.Type,
		arg.Price,
	)
	var i Ticket
	err := row.Scan(
		&i.ID,
		&i.EventID,
		&i.AttendeeID,
		&i.Type,
		&i.Price,
	)
	return i, err
}