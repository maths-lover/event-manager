package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomSpeaker(t *testing.T) Speaker {
	event := createRandomEvent(t)
	args := CreateSpeakerParams{
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Bio: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		ProfileImage: []byte(util.RandomString(10)),
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
	}

	speaker, err := testQueries.CreateSpeaker(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, speaker)

	require.NotZero(t, speaker.ID)
	require.NotZero(t, speaker.Name)

	require.Equal(t, args.Name, speaker.Name)
	require.Equal(t, args.Email, speaker.Email)
	require.Equal(t, args.Phone, speaker.Phone)
	require.Equal(t, args.EventID, speaker.EventID)
	require.Equal(t, args.ProfileImage, speaker.ProfileImage)

	return speaker
}

func TestCreateSpeaker(t *testing.T) {
	createRandomSpeaker(t)
}

func TestListSpeaker(t *testing.T) {
	speaker1 := createRandomSpeaker(t)
	speaker2, err := testQueries.ListSpeaker(context.Background(), speaker1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, speaker2)

	require.Equal(t, reflect.DeepEqual(speaker2, speaker1), true)
}

func TestDeleteSpeaker(t *testing.T) {
	speaker1 := createRandomSpeaker(t)
	err := testQueries.DeleteSpeaker(context.Background(), speaker1.ID)
	require.NoError(t, err)

	speaker2, err := testQueries.ListSpeaker(context.Background(), speaker1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, speaker2)
}

func TestUpdateSpeaker(t *testing.T) {
	speaker1 := createRandomSpeaker(t)
	event := createRandomEvent(t)
	args := UpdateSpeakerParams{
		ID:    speaker1.ID,
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Bio: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		ProfileImage: []byte(util.RandomString(10)),
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
	}

	speaker2, err := testQueries.UpdateSpeaker(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, speaker2)

	require.Equal(t, args.Name, speaker2.Name)
	require.Equal(t, args.Email, speaker2.Email)
	require.Equal(t, args.Phone, speaker2.Phone)
	require.Equal(t, args.EventID, speaker2.EventID)
	require.Equal(t, args.ProfileImage, speaker2.ProfileImage)
}

func TestListSpeakers(t *testing.T) {
	// create random events first
	for i := 0; i < 10; i++ {
		createRandomSpeaker(t)
	}

	args := ListSpeakersParams{
		Limit:  5,
		Offset: 5,
	}

	speakers, err := testQueries.ListSpeakers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, speakers, 5)

	for _, speaker := range speakers {
		require.NotEmpty(t, speaker)
	}
}
