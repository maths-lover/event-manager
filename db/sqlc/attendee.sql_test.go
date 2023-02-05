package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomAttendee(t *testing.T) Attendee {
	event := createRandomEvent(t)
	args := CreateAttendeeParams{
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		TicketNumber: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
		PaymentStatus: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	attendee, err := testQueries.CreateAttendee(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, attendee)

	require.NotZero(t, attendee.ID)
	require.NotZero(t, attendee.Name)

	require.Equal(t, args.Name, attendee.Name)
	require.Equal(t, args.Email, attendee.Email)
	require.Equal(t, args.Phone, attendee.Phone)
	require.Equal(t, args.EventID, attendee.EventID)
	require.Equal(t, args.TicketNumber, attendee.TicketNumber)
	require.Equal(t, args.PaymentStatus, attendee.PaymentStatus)

	return attendee
}

func TestCreateAttendee(t *testing.T) {
	createRandomAttendee(t)
}

func TestListAttendee(t *testing.T) {
	attendee1 := createRandomAttendee(t)
	attendee2, err := testQueries.ListAttendee(context.Background(), attendee1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, attendee2)

	require.Equal(t, reflect.DeepEqual(attendee2, attendee1), true)
}

func TestDeleteAttendee(t *testing.T) {
	attendee1 := createRandomAttendee(t)
	err := testQueries.DeleteAttendee(context.Background(), attendee1.ID)
	require.NoError(t, err)

	attendee2, err := testQueries.ListAttendee(context.Background(), attendee1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, attendee2)
}

func TestUpdateAttendee(t *testing.T) {
	attendee1 := createRandomAttendee(t)
	event := createRandomEvent(t)
	args := UpdateAttendeeParams{
		ID:    attendee1.ID,
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		TicketNumber: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
		PaymentStatus: sql.NullBool{
			Bool:  true,
			Valid: true,
		},
	}

	attendee2, err := testQueries.UpdateAttendee(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, attendee2)

	require.Equal(t, args.Name, attendee2.Name)
	require.Equal(t, args.Email, attendee2.Email)
	require.Equal(t, args.Phone, attendee2.Phone)
	require.Equal(t, args.EventID, attendee2.EventID)
	require.Equal(t, args.TicketNumber, attendee2.TicketNumber)
	require.Equal(t, args.PaymentStatus, attendee2.PaymentStatus)
}

func TestListAttendees(t *testing.T) {
	// create random events first
	for i := 0; i < 10; i++ {
		createRandomAttendee(t)
	}

	args := ListAttendeesParams{
		Limit:  5,
		Offset: 5,
	}

	attendees, err := testQueries.ListAttendees(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, attendees, 5)

	for _, attendee := range attendees {
		require.NotEmpty(t, attendee)
	}
}
