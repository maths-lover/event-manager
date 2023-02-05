package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomOrganizer(t *testing.T) Organizer {
	arg := CreateOrganizerParams{
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Company: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		Logo: []byte(util.RandomString(10)),
		Address: sql.NullString{
			String: util.RandomString(100),
			Valid:  true,
		},
	}

	organizer, err := testQueries.CreateOrganizer(context.Background(), arg)
	require.NoError(t, err)
	require.NotEmpty(t, organizer)

	require.NotZero(t, organizer.ID)
	require.NotZero(t, organizer.Name)
	require.NotZero(t, organizer.Email)
	require.NotZero(t, organizer.Logo)

	require.Equal(t, arg.Name, organizer.Name)
	require.Equal(t, arg.Email, organizer.Email)
	require.Equal(t, arg.Phone, organizer.Phone)
	require.Equal(t, arg.Company, organizer.Company)
	require.Equal(t, arg.Logo, organizer.Logo)
	require.Equal(t, arg.Address, organizer.Address)

	return organizer
}

func TestCreateOrganizer(t *testing.T) {
	createRandomOrganizer(t)
}

func TestListOrganizer(t *testing.T) {
	organizer1 := createRandomOrganizer(t)
	organizer2, err := testQueries.ListOrganizer(context.Background(), organizer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, organizer2)

	require.Equal(t, reflect.DeepEqual(organizer2, organizer1), true)
}

func TestDeleteOrganizer(t *testing.T) {
	organizer1 := createRandomOrganizer(t)
	err := testQueries.DeleteOrganizer(context.Background(), organizer1.ID)
	require.NoError(t, err)

	organizer2, err := testQueries.ListOrganizer(context.Background(), organizer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, organizer2)
}

func TestUpdateOrganizer(t *testing.T) {
	organizer1 := createRandomOrganizer(t)
	args := UpdateOrganizerParams{
		ID:    organizer1.ID,
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Company: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		Logo: []byte(util.RandomString(10)),
		Address: sql.NullString{
			String: util.RandomString(100),
			Valid:  true,
		},
	}

	organizer2, err := testQueries.UpdateOrganizer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, organizer2)

	require.Equal(t, args.Name, organizer2.Name)
	require.Equal(t, args.Email, organizer2.Email)
	require.Equal(t, args.Phone, organizer2.Phone)
	require.Equal(t, args.Company, organizer2.Company)
	require.Equal(t, args.Logo, organizer2.Logo)
	require.Equal(t, args.Address, organizer2.Address)
}

func TestListOrganizers(t *testing.T) {
	// create random organizers first
	for i := 0; i < 10; i++ {
		createRandomOrganizer(t)
	}

	args := ListOrganizersParams{
		Limit:  5,
		Offset: 5,
	}

	organizers, err := testQueries.ListOrganizers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, organizers, 5)

	for _, organizer := range organizers {
		require.NotEmpty(t, organizer)
	}
}
