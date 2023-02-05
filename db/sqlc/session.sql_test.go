package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"
	"time"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomSession(t *testing.T) Session {
	speaker := createRandomSpeaker(t)
	args := CreateSessionParams{
		Title: util.RandomString(10),
		Description: sql.NullString{
			String: util.RandomString(100),
			Valid:  true,
		},
		Venue: sql.NullString{
			String: util.RandomString(50),
			Valid:  true,
		},
		EventID: speaker.EventID,
		SpeakerID: sql.NullInt32{
			Int32: speaker.ID,
			Valid: true,
		},
		DateTime: time.Now().UTC(),
	}

	session, err := testQueries.CreateSession(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, session)

	require.NotZero(t, session.ID)
	require.NotZero(t, session.Title)
	require.NotZero(t, session.DateTime)

	require.Equal(t, args.Title, session.Title)
	require.Equal(t, args.Description, session.Description)
	require.Equal(t, args.Venue, session.Venue)
	require.Equal(t, args.EventID, session.EventID)
	require.Equal(t, args.SpeakerID, session.SpeakerID)
	require.WithinDuration(t, args.DateTime, session.DateTime, time.Second)

	return session
}

func TestCreateSession(t *testing.T) {
	createRandomSession(t)
}

func TestListSession(t *testing.T) {
	session1 := createRandomSession(t)
	session2, err := testQueries.ListSession(context.Background(), session1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, reflect.DeepEqual(session2, session1), true)
}

func TestDeleteSession(t *testing.T) {
	session1 := createRandomSession(t)
	err := testQueries.DeleteSession(context.Background(), session1.ID)
	require.NoError(t, err)

	session2, err := testQueries.ListSession(context.Background(), session1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, session2)
}

func TestUpdateSession(t *testing.T) {
	session1 := createRandomSession(t)
	speaker := createRandomSpeaker(t)
	args := UpdateSessionParams{
		ID:    session1.ID,
		Title: util.RandomString(10),
		Description: sql.NullString{
			String: util.RandomString(100),
			Valid:  true,
		},
		Venue: sql.NullString{
			String: util.RandomString(50),
			Valid:  true,
		},
		EventID: speaker.EventID,
		SpeakerID: sql.NullInt32{
			Int32: speaker.ID,
			Valid: true,
		},
		DateTime: time.Now().UTC(),
	}

	session2, err := testQueries.UpdateSession(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, session2)

	require.Equal(t, args.Title, session2.Title)
	require.Equal(t, args.Description, session2.Description)
	require.Equal(t, args.Venue, session2.Venue)
	require.Equal(t, args.EventID, session2.EventID)
	require.Equal(t, args.SpeakerID, session2.SpeakerID)
	require.WithinDuration(t, args.DateTime, session2.DateTime, time.Second)
}

func TestListSessions(t *testing.T) {
	// create random organizers first
	for i := 0; i < 10; i++ {
		createRandomSession(t)
	}

	args := ListSessionsParams{
		Limit:  5,
		Offset: 5,
	}

	sessions, err := testQueries.ListSessions(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, sessions, 5)

	for _, session := range sessions {
		require.NotEmpty(t, session)
	}
}
