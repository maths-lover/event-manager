package db

import (
	"context"
	"database/sql"
	"fmt"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomTicket(t *testing.T) Ticket {
	attendee := createRandomAttendee(t)
	args := CreateTicketParams{
		Price: fmt.Sprintf("%.2f", util.RandomDecNum()),
		Type: sql.NullString{
			String: util.RandomString(50),
			Valid:  true,
		},
		EventID: attendee.EventID,
		AttendeeID: sql.NullInt32{
			Int32: attendee.ID,
			Valid: true,
		},
	}

	session, err := testQueries.CreateTicket(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.NotZero(t, session.ID)
	require.NotZero(t, session.Price)

	require.Equal(t, args.Price, session.Price)
	require.Equal(t, args.Type, session.Type)
	require.Equal(t, args.EventID, session.EventID)
	require.Equal(t, args.AttendeeID, session.AttendeeID)

	return session
}

func TestCreateTicket(t *testing.T) {
	createRandomTicket(t)
}

func TestListTicket(t *testing.T) {
	session1 := createRandomTicket(t)
	session2, err := testQueries.ListTicket(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, reflect.DeepEqual(session2, session1), true)
}

func TestDeleteTicket(t *testing.T) {
	session1 := createRandomTicket(t)
	err := testQueries.DeleteTicket(context.Background(), session1.ID)
	require.NoError(t, err)

	session2, err := testQueries.ListTicket(context.Background(), session1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, session2)
}

func TestUpdateTicket(t *testing.T) {
	session1 := createRandomTicket(t)
	attendee := createRandomAttendee(t)
	args := UpdateTicketParams{
		ID:    session1.ID,
		Price: fmt.Sprintf("%.2f", util.RandomDecNum()),
		Type: sql.NullString{
			String: util.RandomString(50),
			Valid:  true,
		},
		EventID: attendee.EventID,
		AttendeeID: sql.NullInt32{
			Int32: attendee.ID,
			Valid: true,
		},
	}

	session2, err := testQueries.UpdateTicket(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, args.Price, session2.Price)
	require.Equal(t, args.Type, session2.Type)
	require.Equal(t, args.EventID, session2.EventID)
	require.Equal(t, args.AttendeeID, session2.AttendeeID)
}

func TestListTickets(t *testing.T) {
	// create random organizers first
	for i := 0; i < 10; i++ {
		createRandomTicket(t)
	}

	args := ListTicketsParams{
		Limit:  5,
		Offset: 5,
	}

	sessions, err := testQueries.ListTickets(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, sessions, 5)

	for _, session := range sessions {
		require.NotEmpty(t, session)
	}
}
