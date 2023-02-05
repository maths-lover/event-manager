package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomEvent(t *testing.T) Event {
	organizer := createRandomOrganizer(t)
	args := CreateEventParams{
		Title: util.RandomString(10),
		Description: sql.NullString{
			String: util.RandomString(100),
			Valid:  true,
		},
		Venue: sql.NullString{
			String: util.RandomString(50),
			Valid:  true,
		},
		Type: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		OrganizerID: sql.NullInt32{
			Int32: organizer.ID,
			Valid: true,
		},
	}

	event, err := testQueries.CreateEvent(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, event)

	require.NotZero(t, event.ID)
	require.NotZero(t, event.Title)

	require.Equal(t, args.Title, event.Title)
	require.Equal(t, args.Description, event.Description)
	require.Equal(t, args.Venue, event.Venue)
	require.Equal(t, args.Type, event.Type)
	require.Equal(t, args.OrganizerID, event.OrganizerID)

	return event
}

func TestCreateEvent(t *testing.T) {
	createRandomEvent(t)
}

func TestListEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	event2, err := testQueries.ListEvent(context.Background(), event1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, event2)

	require.Equal(t, reflect.DeepEqual(event2, event1), true)
}

func TestDeleteEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	err := testQueries.DeleteEvent(context.Background(), event1.ID)
	require.NoError(t, err)

	event2, err := testQueries.ListEvent(context.Background(), event1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, event2)
}

func TestUpdateEvent(t *testing.T) {
	event1 := createRandomEvent(t)
	organizer := createRandomOrganizer(t)
	args := UpdateEventParams{
		ID:    event1.ID,
		Title: util.RandomString(10),
		Description: sql.NullString{
			String: util.RandomString(100),
			Valid:  true,
		},
		Venue: sql.NullString{
			String: util.RandomString(50),
			Valid:  true,
		},
		Type: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		OrganizerID: sql.NullInt32{
			Int32: organizer.ID,
			Valid: true,
		},
	}

	event2, err := testQueries.UpdateEvent(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, event2)

	require.Equal(t, args.Title, event2.Title)
	require.Equal(t, args.Description, event2.Description)
	require.Equal(t, args.Venue, event2.Venue)
	require.Equal(t, args.Type, event2.Type)
	require.Equal(t, args.OrganizerID, event2.OrganizerID)
}

func TestListEvents(t *testing.T) {
	// create random organizers first
	for i := 0; i < 10; i++ {
		createRandomEvent(t)
	}

	args := ListEventsParams{
		Limit:  5,
		Offset: 5,
	}

	events, err := testQueries.ListEvents(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, events, 5)

	for _, event := range events {
		require.NotEmpty(t, event)
	}
}
