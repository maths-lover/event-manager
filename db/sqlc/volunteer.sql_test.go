package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomVolunteer(t *testing.T) Volunteer {
	event := createRandomEvent(t)
	args := CreateVolunteerParams{
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Role: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
		Availability: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
	}

	volunteer, err := testQueries.CreateVolunteer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, volunteer)

	require.NotZero(t, volunteer.ID)
	require.NotZero(t, volunteer.Name)

	require.Equal(t, args.Name, volunteer.Name)
	require.Equal(t, args.Email, volunteer.Email)
	require.Equal(t, args.Phone, volunteer.Phone)
	require.Equal(t, args.EventID, volunteer.EventID)
	require.Equal(t, args.Role, volunteer.Role)
	require.Equal(t, args.Availability, volunteer.Availability)

	return volunteer
}

func TestCreateVolunteer(t *testing.T) {
	createRandomVolunteer(t)
}

func TestListVolunteer(t *testing.T) {
	volunteer1 := createRandomVolunteer(t)
	volunteer2, err := testQueries.ListVolunteer(context.Background(), volunteer1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, volunteer2)

	require.Equal(t, reflect.DeepEqual(volunteer2, volunteer1), true)
}

func TestDeleteVolunteer(t *testing.T) {
	volunteer1 := createRandomVolunteer(t)
	err := testQueries.DeleteVolunteer(context.Background(), volunteer1.ID)
	require.NoError(t, err)

	volunteer2, err := testQueries.ListVolunteer(context.Background(), volunteer1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, volunteer2)
}

func TestUpdateVolunteer(t *testing.T) {
	volunteer1 := createRandomVolunteer(t)
	event := createRandomEvent(t)
	args := UpdateVolunteerParams{
		ID:    volunteer1.ID,
		Name:  util.RandomString(10),
		Email: util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		Role: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
		Availability: sql.NullString{
			String: util.RandomString(10),
			Valid:  true,
		},
	}

	volunteer2, err := testQueries.UpdateVolunteer(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, volunteer2)

	require.Equal(t, args.Name, volunteer2.Name)
	require.Equal(t, args.Email, volunteer2.Email)
	require.Equal(t, args.Phone, volunteer2.Phone)
	require.Equal(t, args.EventID, volunteer2.EventID)
	require.Equal(t, args.Role, volunteer2.Role)
	require.Equal(t, args.Availability, volunteer2.Availability)
}

func TestListVolunteers(t *testing.T) {
	// create random events first
	for i := 0; i < 10; i++ {
		createRandomVolunteer(t)
	}

	args := ListVolunteersParams{
		Limit:  5,
		Offset: 5,
	}

	volunteers, err := testQueries.ListVolunteers(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, volunteers, 5)

	for _, volunteer := range volunteers {
		require.NotEmpty(t, volunteer)
	}
}
