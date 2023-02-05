package db

import (
	"context"
	"database/sql"
	"reflect"
	"testing"

	"github.com/maths-lover/event-manager/util"
	"github.com/stretchr/testify/require"
)

func createRandomExhibition(t *testing.T) Exhibition {
	event := createRandomEvent(t)
	args := CreateExhibitionParams{
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
		CompanyName:   util.RandomString(10),
		ContactPerson: util.RandomString(10),
		Email:         util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		BoothNumber: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		FloorPlan: []byte(util.RandomString(20)),
	}

	exhibition, err := testQueries.CreateExhibition(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, exhibition)

	require.NotZero(t, exhibition.ID)
	require.NotZero(t, exhibition.CompanyName)

	require.Equal(t, args.EventID, exhibition.EventID)
	require.Equal(t, args.CompanyName, exhibition.CompanyName)
	require.Equal(t, args.ContactPerson, exhibition.ContactPerson)
	require.Equal(t, args.Email, exhibition.Email)
	require.Equal(t, args.Phone, exhibition.Phone)
	require.Equal(t, args.BoothNumber, exhibition.BoothNumber)
	require.Equal(t, args.FloorPlan, exhibition.FloorPlan)

	return exhibition
}

func TestCreateExhibition(t *testing.T) {
	createRandomExhibition(t)
}

func TestListExhibition(t *testing.T) {
	exhibition1 := createRandomExhibition(t)
	exhibition2, err := testQueries.ListExhibition(context.Background(), exhibition1.ID)
	require.NoError(t, err)
	require.NotEmpty(t, exhibition2)

	require.Equal(t, reflect.DeepEqual(exhibition2, exhibition1), true)
}

func TestDeleteExhibition(t *testing.T) {
	exhibition1 := createRandomExhibition(t)
	err := testQueries.DeleteExhibition(context.Background(), exhibition1.ID)
	require.NoError(t, err)

	exhibition2, err := testQueries.ListExhibition(context.Background(), exhibition1.ID)
	require.Error(t, err)
	require.EqualError(t, err, sql.ErrNoRows.Error())
	require.Empty(t, exhibition2)
}

func TestUpdateExhibition(t *testing.T) {
	exhibition1 := createRandomExhibition(t)
	event := createRandomEvent(t)
	args := UpdateExhibitionParams{
		ID: exhibition1.ID,
		EventID: sql.NullInt32{
			Int32: event.ID,
			Valid: true,
		},
		CompanyName:   util.RandomString(10),
		ContactPerson: util.RandomString(10),
		Email:         util.RandomString(7) + "@mail.com",
		Phone: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		BoothNumber: sql.NullString{
			String: util.RandomPhoneNum(),
			Valid:  true,
		},
		FloorPlan: []byte(util.RandomString(20)),
	}

	exhibition2, err := testQueries.UpdateExhibition(context.Background(), args)
	require.NoError(t, err)
	require.NotEmpty(t, exhibition2)

	require.Equal(t, args.EventID, exhibition2.EventID)
	require.Equal(t, args.CompanyName, exhibition2.CompanyName)
	require.Equal(t, args.ContactPerson, exhibition2.ContactPerson)
	require.Equal(t, args.Email, exhibition2.Email)
	require.Equal(t, args.Phone, exhibition2.Phone)
	require.Equal(t, args.BoothNumber, exhibition2.BoothNumber)
	require.Equal(t, args.FloorPlan, exhibition2.FloorPlan)
}

func TestListExhibitions(t *testing.T) {
	// create random events first
	for i := 0; i < 10; i++ {
		createRandomExhibition(t)
	}

	args := ListExhibitionsParams{
		Limit:  5,
		Offset: 5,
	}

	exhibitions, err := testQueries.ListExhibitions(context.Background(), args)
	require.NoError(t, err)
	require.Len(t, exhibitions, 5)

	for _, exhibition := range exhibitions {
		require.NotEmpty(t, exhibition)
	}
}
