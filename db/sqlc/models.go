// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

package db

import (
	"database/sql"
	"time"
)

type Attendee struct {
	ID            int32          `json:"id"`
	Name          string         `json:"name"`
	Email         string         `json:"email"`
	Phone         sql.NullString `json:"phone"`
	EventID       sql.NullInt32  `json:"event_id"`
	TicketNumber  sql.NullString `json:"ticket_number"`
	PaymentStatus sql.NullBool   `json:"payment_status"`
	PaymentDate   sql.NullTime   `json:"payment_date"`
}

type Event struct {
	ID          int32          `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	DateTime    time.Time      `json:"date_time"`
	Venue       sql.NullString `json:"venue"`
	Type        sql.NullString `json:"type"`
	OrganizerID sql.NullInt32  `json:"organizer_id"`
}

type Exhibition struct {
	ID            int32          `json:"id"`
	EventID       sql.NullInt32  `json:"event_id"`
	CompanyName   string         `json:"company_name"`
	ContactPerson string         `json:"contact_person"`
	Email         string         `json:"email"`
	Phone         sql.NullString `json:"phone"`
	BoothNumber   sql.NullString `json:"booth_number"`
	FloorPlan     []byte         `json:"floor_plan"`
}

type Organizer struct {
	ID      int32          `json:"id"`
	Name    string         `json:"name"`
	Email   string         `json:"email"`
	Phone   sql.NullString `json:"phone"`
	Company sql.NullString `json:"company"`
	Logo    []byte         `json:"logo"`
	Address sql.NullString `json:"address"`
}

type Session struct {
	ID          int32          `json:"id"`
	Title       string         `json:"title"`
	Description sql.NullString `json:"description"`
	DateTime    time.Time      `json:"date_time"`
	Venue       sql.NullString `json:"venue"`
	EventID     sql.NullInt32  `json:"event_id"`
	SpeakerID   sql.NullInt32  `json:"speaker_id"`
}

type Speaker struct {
	ID           int32          `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Phone        sql.NullString `json:"phone"`
	EventID      sql.NullInt32  `json:"event_id"`
	ProfileImage []byte         `json:"profile_image"`
	Bio          sql.NullString `json:"bio"`
}

type Sponsorship struct {
	ID            int32          `json:"id"`
	EventID       sql.NullInt32  `json:"event_id"`
	CompanyName   string         `json:"company_name"`
	ContactPerson string         `json:"contact_person"`
	Email         string         `json:"email"`
	Phone         sql.NullString `json:"phone"`
	Package       sql.NullString `json:"package"`
}

type Ticket struct {
	ID         int32          `json:"id"`
	EventID    sql.NullInt32  `json:"event_id"`
	AttendeeID sql.NullInt32  `json:"attendee_id"`
	Type       sql.NullString `json:"type"`
	Price      string         `json:"price"`
}

type Volunteer struct {
	ID           int32          `json:"id"`
	Name         string         `json:"name"`
	Email        string         `json:"email"`
	Phone        sql.NullString `json:"phone"`
	EventID      sql.NullInt32  `json:"event_id"`
	Role         sql.NullString `json:"role"`
	Availability sql.NullString `json:"availability"`
}
